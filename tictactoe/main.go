package main 


func main() {
	p1 := NewPlayer("Player 1" , 'X')
	p2 := NewPlayer("Player 2" , 'O')
	g := NewGame(p1, p2)

	g.Play()

}