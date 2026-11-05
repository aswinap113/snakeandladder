package main

import "snake/entities"

func main() {

	boardEntities := []entities.BoardEntities{
		entities.NewSnake(17, 7),
		entities.NewSnake(54, 34),
		entities.NewSnake(62, 19),
		entities.NewSnake(98, 79),
		entities.NewLadder(3, 38),
		entities.NewLadder(24, 33),
		entities.NewLadder(42, 93),
		entities.NewLadder(72, 84),
	}

	players := []string{"Alice", "Bob"}

	game := entities.NewGameBuilder().
		SetBoard(100, boardEntities).
		SetPlayers(players).
		SetDice(entities.NewDice(1, 6)).
		Build()

	game.Play()
}
