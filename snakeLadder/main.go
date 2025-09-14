package main 

import (
	"fmt"
)

func main() {
	var boardSize, playersCount int 
	fmt.Print("Enter boardsize and players count: ")
	fmt.Scanln(&boardSize, &playersCount)	
	game := NewGame(boardSize, playersCount)
	obstacleFactory := NewObstacleFactory()

	fmt.Print("Enter the names of the players: ")
	var name string 
	for i := 0 ; i < playersCount ; i++ {
		fmt.Scanln(&name)
		game.AddPlayer(name)
	}	
	var faces int 
	fmt.Print("Enter the no of faces in the dice : ")
	fmt.Scanln(&faces)
	game.AddDice(NewDice(faces))


	var laddersCount , snakesCount int 
	fmt.Print("Enter the ladders count ")
	fmt.Scanln(&laddersCount)
	var src, dest int 
	for i := 0 ; i < laddersCount ; i++ {
		fmt.Scanln(&src, &dest)
		ob := obstacleFactory.CreateObstacle(src, dest, LADDER)
		game.AddObstacle(ob)
	}


	fmt.Print("Enter the snakes count ")
	fmt.Scanln(&snakesCount)
	for i := 0 ; i < snakesCount ; i++ {
		fmt.Scanln(&src, &dest)
		ob := obstacleFactory.CreateObstacle(src, dest, SNAKE)
		game.AddObstacle(ob)
	}

	fmt.Println("Starting the game now ")
	game.StartGame()
}

/*

Enter boardsize and players count: 50 2
Enter the names of the players: alice
bob
Enter the no of faces in the dice : 7
Enter the ladders count3
12 20
25 45
8 40
Enter the snakes count3
43 7
34 19
25 22

*/