package main  

type ObstacleFactory struct {}


func NewObstacleFactory() *ObstacleFactory{
	return &ObstacleFactory{}
}

func(o *ObstacleFactory) CreateObstacle(src, dest int, obstacleType ObstacleType) *Obstacle {
	return &Obstacle{
		src: src, 
		dest: dest, 
		obstacleType: obstacleType,
	}
}