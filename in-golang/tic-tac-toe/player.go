package tictactoe

type Player struct {
	id int
}

func (p *Player) GetId() int {
	return p.id
}
