package entities

type PlayersQueue []*Player

func (q *PlayersQueue) enqueue(player *Player) {
	*q = append(*q, player)
}

func (q *PlayersQueue) dequeue() (*Player, bool) {
	if len(*q) == 0 {
		return nil, false
	}

	player := (*q)[0]
	*q = (*q)[1:]
	return player, true
}
