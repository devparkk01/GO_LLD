package main

import "math/rand"

type Dice struct {
	faces int // represents no of faces it has
}

func NewDice(faces int) *Dice {
	return &Dice{faces: faces }
}

func (d *Dice) Roll() int {
	next := rand.Intn(d.faces) + 1 // returns a random value from 1 to d.faces 
	return next 
}