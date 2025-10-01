package main

import "fmt"

type Cell byte

const (
	CELL_EMPTY Cell = '-'
	CELL_X     Cell = 'X'
	CELL_O     Cell = 'O'
)

type Board struct {
	cells      [3][3]Cell
	movesCount int
}

func NewBoard() *Board {
	b := &Board{}
	// initialise all cells with empty cell
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.cells[i][j] = CELL_EMPTY
		}
	}
	return b
}

// returns if the board is full 
func (b *Board) IsFull() bool {
	return b.movesCount == 9
}

// returns if the given cell is empty 
func (b *Board) IsCellEmpty(row, col int) bool {
	return b.cells[row][col] == CELL_EMPTY
}


// make a move with the given row, col and symbol 
func (b *Board) MakeMove(row, col int, symbol Cell) error {
	if row < 0 || row > 2 || col < 0 || col > 2 || !b.IsCellEmpty(row, col) {
		return fmt.Errorf("invalid move")
	}
	b.cells[row][col] = symbol 
	b.movesCount++ 
	return nil 
}

// checks if the board has any winner yet 
func (b *Board) HasWinner() bool {
	// check row 
	for row := 0 ; row <= 2; row++ {
		if b.cells[row][0] != CELL_EMPTY && b.cells[row][0] == b.cells[row][1] && b.cells[row][1] == b.cells[row][2] {
			return true 
		}
	}
	// check column 
	for col := 0 ; col <= 2; col++ {
		if b.cells[0][col] != CELL_EMPTY && b.cells[0][col] == b.cells[1][col] && b.cells[1][col] == b.cells[2][col] {
			return true
		}
	}
	// check first diagonal 
	if b.cells[0][0] != CELL_EMPTY && b.cells[0][0] == b.cells[1][1] && b.cells[1][1] == b.cells[2][2] {
		return true 
	}
	// check second diagonal 
	if b.cells[0][2] != CELL_EMPTY && b.cells[0][2] == b.cells[1][1] && b.cells[1][1] == b.cells[2][0] {
		return true 
	}
	return false 
}

// prints board 
func (b *Board) PrintBoard() {
	for row := 0 ; row <= 2 ; row++ {
		for col := 0 ; col <= 2; col++ {
			fmt.Printf(string(b.cells[row][col]) + " ")
		}
		fmt.Println()
	}
}