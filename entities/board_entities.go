package entities

type BoardEntities interface {
	Start() int
	End() int
}

type boardEntities struct {
	start int
	end   int
}

func (b *boardEntities) Start() int {
	return b.start
}

func (b *boardEntities) End() int {
	return b.end
}
