package main 

type Player struct {
	id string 
	pos int 
}

func NewPlayer(id string) *Player {
	return &Player{
		id: id,
		pos: 1,
	}
}