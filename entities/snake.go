package entities

type Snake struct {
	boardEntities
}

func NewSnake(start int, end int) *Snake {
	if start <= end {
		panic("Start should be greater than the end for a snake!")
	}

	return &Snake{
		boardEntities{
			start: start,
			end:   end,
		},
	}
}
