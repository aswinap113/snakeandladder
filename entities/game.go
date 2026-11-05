package entities

import (
	"fmt"
	"snake/enums"
)

type Game struct {
	board        Board
	dice         Dice
	playersQueue PlayersQueue
	status       enums.GameStatus
	winner       Player
}

func newGame(builder *GameBuilder) *Game {
	playersQ := make(PlayersQueue, 0, len(builder.players))
	for _, player := range builder.players {
		playersQ.enqueue(player)
	}
	return &Game{
		board:        *builder.board,
		dice:         *builder.dice,
		playersQueue: playersQ,
		status:       enums.GameStatusPaused,
	}
}

func (game *Game) Play() {
	if len(game.playersQueue) < 2 {
		panic("Can not start without 2 players")
	}

	game.status = enums.GameStatusPlaying
	fmt.Println("Game started")

	for game.status == enums.GameStatusPlaying {
		currentPlayer, err := game.playersQueue.dequeue()
		if err != true {
			break
		}

		game.takeTurn(currentPlayer)
		if game.status == enums.GameStatusPlaying {
			game.playersQueue.enqueue(currentPlayer)
		}
	}

	fmt.Println("Game finished")

}

func (game *Game) takeTurn(player *Player) {
	fmt.Printf("Player %s turn : \n", player.name)
	turnStartPosition := player.Position()
	consecutiveMaxRolls := 0

	for {
		diceVal := game.dice.Roll()
		fmt.Printf("Player %s rolled : %d\n", player.name, diceVal)

		if diceVal == game.dice.maxNumber {
			consecutiveMaxRolls++
			if consecutiveMaxRolls == 3 {
				player.SetPosition(turnStartPosition)
				fmt.Printf("Player %s rolled the maximum three times in a row and returns to %d.\n", player.name, turnStartPosition)
				return
			}
		} else {
			consecutiveMaxRolls = 0
		}

		currentPos := player.Position()
		newPosition := currentPos + diceVal

		if newPosition > game.board.size {
			fmt.Printf("Oops! Player %s needs to land exactly on %d. Turn skipped.\n", player.name, game.board.size-currentPos)
			if diceVal != game.dice.maxNumber {
				return
			}
			continue
		}
		if newPosition == game.board.size {
			game.status = enums.GameStatusGameOver
			player.SetPosition(newPosition)
			game.winner = *player
			fmt.Println("---")
			fmt.Printf("Hooray! %s reached the final square %d and won!", player.name, game.board.size)
			return
		}

		updatedPositionAfterEntities := game.board.GetUpdatedPositionAfterEntities(newPosition)
		if updatedPositionAfterEntities < newPosition {
			fmt.Println("---")
			fmt.Printf("Oops! Player %s was bitten by a snake at %d and slid down to %d.\n", player.name, newPosition, updatedPositionAfterEntities)
		}
		if updatedPositionAfterEntities > newPosition {
			fmt.Printf("Wow! Player %s found a ladder at %d and climbed to %d.\n", player.name, newPosition, updatedPositionAfterEntities)
		}
		fmt.Printf("Player %s moved to %d.\n", player.name, updatedPositionAfterEntities)
		player.SetPosition(updatedPositionAfterEntities)

		if diceVal != game.dice.maxNumber {
			return
		}
	}
}

type GameBuilder struct {
	board   *Board
	players []*Player
	dice    *Dice
}

func NewGameBuilder() *GameBuilder {
	return &GameBuilder{}
}

func (b *GameBuilder) SetBoard(boardSize int, boardEntities []BoardEntities) *GameBuilder {
	b.board = NewBoard(boardSize, boardEntities)
	return b
}

func (b *GameBuilder) SetPlayers(playerNames []string) *GameBuilder {
	b.players = []*Player{}
	for _, playerName := range playerNames {
		b.players = append(b.players, NewPlayer(playerName))
	}
	return b
}

func (b *GameBuilder) SetDice(dice *Dice) *GameBuilder {
	b.dice = dice
	return b
}

func (b *GameBuilder) Build() *Game {
	if b.board == nil || len(b.players) == 0 || b.dice == nil {
		panic("Board, Players, and Dice must be set.")
	}
	return newGame(b)
}
