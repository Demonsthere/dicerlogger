package dice

import (
	"fmt"
	"math/rand"
	"time"
)

//Dice dice for randomizing log entry
type Dice struct {
	side int
}

//NewDice Create a new Dice of size x
func NewDice(size int) *Dice {
	return &Dice{side: size}
}

//Roll return a random number 1 - size
func (d Dice) Roll() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(d.side) + 1
}

//RollString return a random number 1 - size
func (d Dice) RollString() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("Number: %d", rand.Intn(d.side)+1)
}
