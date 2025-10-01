package main

import "fmt"

type Game struct {
	board         *Board
	player1       *Player
	player2       *Player
	currentPlayer *Player
	winner        *Player
}

func NewGame(player1 *Player, player2 *Player) *Game {
	return &Game{
		board:         NewBoard(),
		player1:       player1,
		player2:       player2,
		currentPlayer: player1,
		winner:        nil,
	}
}

// start playing the game 
func (g *Game) Play() {
	g.board.PrintBoard()
	for !g.board.IsFull() && g.winner == nil {
		fmt.Printf("\nPlayer %s turn \n", g.currentPlayer.name)
		validMove := false 
		row := 1
		col := 1 
		for !validMove {
			row, col = g.getInput() 
			if g.isValidMove(row, col) {
				validMove = true 
			}
		}

		err := g.board.MakeMove(row, col, g.currentPlayer.symbol)
		if err != nil {
			fmt.Println(err.Error())
			continue 
		}
		g.board.PrintBoard() 

		if g.board.HasWinner() {
			g.winner = g.currentPlayer 
			continue 
		}
		g.switchPlayer()

	}
	if g.winner != nil{
		fmt.Printf("\n\nPlayer %s is the winner.\n", g.winner.name)
	} else {
		fmt.Println("\n\nIt's a draw ")
	}
}

// take user input 
func(g *Game) getInput() (int, int) {
	var row, col int 
	fmt.Println("Enter row (0-2)")
	fmt.Scanln(&row)
	fmt.Println("Enter col (0-2)")
	fmt.Scanln(&col)
	return row, col 
}

func (g *Game) isValidMove(row, col int) bool {
	if (row < 0 || row > 2) {
		fmt.Println("Row values must be between 0 and 2")
		return false 
	}
	if (col < 0 || col > 2) {
		fmt.Println("col values must be between 0 and 2")
		return false 
	}

	if !g.board.IsCellEmpty(row, col) {
		fmt.Println("Given cell is not empty.")
		return false 
	}
	return true 

}

func (g *Game) switchPlayer() {
	if g.currentPlayer == g.player1 {
		g.currentPlayer = g.player2
	} else {
		g.currentPlayer = g.player1 
	}
}