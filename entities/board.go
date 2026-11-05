package entities

type Board struct {
	size     int
	entities map[int]int
}

func NewBoard(size int, entities []BoardEntities) *Board {
	board := &Board{
		size:     size,
		entities: make(map[int]int),
	}
	for _, entity := range entities {
		board.entities[entity.Start()] = entity.End()
	}
	return board
}

func (b *Board) Size() int {
	return b.size
}

func (b *Board) GetUpdatedPositionAfterEntities(position int) int {
	if _, ok := b.entities[position]; ok {
		endPosition := b.entities[position]
		return endPosition
	} else {
		return position
	}
}
