package main

import (
	"time"

	dice "github.com/demonsthere/dicelogger/pkg/dice"
	env "github.com/demonsthere/dicelogger/pkg/env"
	logline "github.com/demonsthere/dicelogger/pkg/logline"
	log "github.com/sirupsen/logrus"
)

func main() {
	env.InitConfig()
	log.SetFormatter(&log.TextFormatter{})
	d := dice.NewDice(env.Config.DiceSize)
	ll := logline.NewLogLines()
	for range time.NewTicker(time.Duration(env.Config.TimeInterval) * time.Millisecond).C {
		rolling(d, ll)
	}
}

func rolling(d *dice.Dice, ll *logline.LogLines) {
	log.Info(ll.GetLine(d.Roll()))
}
