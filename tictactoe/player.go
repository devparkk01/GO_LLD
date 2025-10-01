package main 

type Player struct {
	name string 
	symbol Cell 
}

func NewPlayer(name string, symbol Cell) *Player {
	return &Player{
		name: name, 
		symbol: symbol,
	}
}