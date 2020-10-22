package main

import (
	"time"

	dice "github.com/demonsthere/dicelogger/pkg/dice"
	logline "github.com/demonsthere/dicelogger/pkg/logline"
	log "github.com/sirupsen/logrus"
)

const (
	diceSize     = 100
	timeInterval = 2
)

func main() {
	// log.SetFormatter(&log.JSONFormatter{})
	d := dice.NewDice(diceSize)
	ll := logline.NewLogLines()
	for range time.NewTicker(timeInterval * time.Second).C {
		rolling(d, ll)
	}
}

func rolling(d *dice.Dice, ll *logline.LogLines) {
	log.Info(ll.GetLine(d.Roll()))
}
