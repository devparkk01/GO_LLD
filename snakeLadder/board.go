package main 

type Board struct {
	size int 
	cells []*Cell
}


func NewBoard(boardSize int) *Board{
	cells := make([]*Cell, 0, boardSize + 1)
	for i := 0 ; i <= boardSize ; i++ {
		cells = append(cells, NewCell(i))
	}

	return &Board{
		size: boardSize,
		cells: cells,
	}
}

func (b *Board) AddObstacle(obstacle *Obstacle) {
	src := obstacle.src 
	b.cells[src].AddObstacle(obstacle)
}


func (b *Board) GetNextPosition(currentPosition int) int {
	cell := b.cells[currentPosition]
	finalPosition := currentPosition
	// if the cell has obstacle( maybe snake or ladder )
	if cell.obstacle != nil {
		finalPosition = cell.obstacle.dest 
	}

	return finalPosition 
}

