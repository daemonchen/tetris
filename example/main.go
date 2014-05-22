package main

import (
	"log"
	"os"
	"time"

	"github.com/gogames/tetris"
)

var (
	game          *tetris.Tetris
	height, width = 30, 20
)

func init() {
	var err error
	game, err = tetris.NewTetris(height, width)
	if err != nil {
		panic(err)
	}
}

func fetchGameScreen() {
	var screen interface{}
	for {
		screen = game.GameScreen()
		// do something
		log.Println("screen: ", screen)
	}
}

func fetchGameScore() {
	var score int
	for {
		score = game.Score()
		// do something with score
		log.Println("score: ", score)
	}
}

func fetchComboScore() {
	var combo int
	for {
		combo = game.ComboScore()
		// do something with score
		log.Println("combo: ", combo)
	}
}

func handleGameOver() {
	// game over signal is a channel, so only call it once
	if game.IsGameOver() {
		log.Println("game is over")
		os.Exit(1)
	}
}

func inputs() {
	game.Input(game.KeyLeft)
	game.Input(game.KeyRight)
	game.Input(game.KeyDown)
}

func main() {
	go fetchGameScore()
	go fetchComboScore()
	go fetchGameScreen()
	go handleGameOver()
	time.Sleep(10 * time.Minute)
}