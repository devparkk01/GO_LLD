package main

import "fmt"

type Game struct {
	board        *Board
	players      []*Player
	dice         *Dice
	winner       *Player
	playersCount int
}

func NewGame(boardSize int, playersCount int) *Game {
	return &Game{
		board:        NewBoard(boardSize), // setup the cells
		players:      make([]*Player, 0, playersCount),
		playersCount: playersCount,
	}
}

func (g *Game) AddPlayer(id string) {
	g.players = append(g.players, NewPlayer(id))
}

func (g *Game) AddObstacle(ob *Obstacle) {
	// Add that obstacle into the board
	g.board.AddObstacle(ob)
}

func (g *Game) AddDice(dice *Dice) {
	g.dice = dice
}

func (g *Game) StartGame() {
	// While there is no winnder
	for g.winner == nil {
		// loop through all players
		for _, player := range g.players {
			currentPos := player.pos // current position of the player
			diceNo := g.dice.Roll()  // roll the dice

			// if the next position gets outside the board, then skip this player's chance
			// fmt.Println("Printing nextPos for player ", nextPos, player.id )
			if currentPos + diceNo > g.board.size {
				continue // move to the next player
				// Keep the player at it's original position 
			}
			// if that's not the case 
			nextPos := currentPos + diceNo 
			fmt.Println("Printing nextPos for player ", nextPos, player.id )

			nextPos = g.board.GetNextPosition(nextPos)
			// set this new position of this player
			player.pos = nextPos
			if nextPos == g.board.size {
				fmt.Printf("Player %s won the game \n", player.id)
				g.winner = player
				break
			}
		}
	}
}
