package enums

type GameStatus string

const (
	GameStatusPlaying  GameStatus = "playing"
	GameStatusPaused   GameStatus = "not_started"
	GameStatusGameOver GameStatus = "game_over"
)
