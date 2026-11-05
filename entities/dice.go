package entities

import "math/rand/v2"

type Dice struct {
	minNumber int
	maxNumber int
}

func NewDice(minNumber, maxNumber int) *Dice {
	return &Dice{
		minNumber: minNumber,
		maxNumber: maxNumber,
	}
}

func (d *Dice) Roll() int {
	val := rand.IntN(d.maxNumber-d.minNumber+1) + d.minNumber
	return val
}
