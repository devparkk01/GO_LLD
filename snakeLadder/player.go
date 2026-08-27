package main 

type Player struct {
	id string 
	pos int 
}

func NewPlayer(id string) *Player {
	// Player starts from position 1 
	return &Player{
		id: id,
		pos: 1,
	}
}