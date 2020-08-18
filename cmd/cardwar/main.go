package main

import (
	"os"

	"github.com/akamensky/argparse"
	"github.com/arkangle/cardwar/pkg/game"
	"github.com/sirupsen/logrus"
)

func main() {
	parser := argparse.NewParser("cardwar", "Card Game of War!")

	playerNames := parser.StringList("p", "player", &argparse.Options{Required: true, Help: "Player name"})

	logger := logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Level = logrus.InfoLevel

	err := parser.Parse(os.Args)
	if err != nil {
		logger.Fatal(parser.Usage(err))
	}
	players := game.NewPlayerCollection(*playerNames)
	deck := game.NewFullDeck()
	dealer := game.NewDealer(&players, &deck)
	dealer.Shuffle()
	dealer.Deal()
}
