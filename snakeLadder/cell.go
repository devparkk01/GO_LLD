package main 


type Cell struct {
	position int 
	obstacle *Obstacle
}

func NewCell(position int) *Cell {
	return &Cell{
		position: position,
	}
}


func(c *Cell) AddObstacle(ob *Obstacle) {
	c.obstacle = ob
}