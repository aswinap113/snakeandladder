package entities

type Player struct {
	name     string
	position int
}

func NewPlayer(name string) *Player {
	return &Player{
		name:     name,
		position: 0,
	}
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Position() int {
	return p.position
}

func (p *Player) SetPosition(pos int) {
	p.position = pos
}
