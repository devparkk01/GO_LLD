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
	// if the cell has obstacle( maybe snake or ladder )
	cell := b.cells[currentPosition]
	// we can have chain of ladders or snakes
	for cell.obstacle != nil {
		nextPos := cell.obstacle.dest
		cell = b.cells[nextPos]
	}
	return cell.position
}

