package entities

type Ladder struct {
	boardEntities
}

func NewLadder(start int, end int) *Ladder {
	if end < start {
		panic("End should be greater that Start for a Laddder!")
	}
	return &Ladder{
		boardEntities: boardEntities{
			start: start,
			end:   end,
		},
	}
}
