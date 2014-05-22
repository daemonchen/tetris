// the exported functions
package tetris

import (
	"encoding/json"
	"fmt"
)

const (
	MinHeight = 20
	MinWidth  = 10
)

var (
	ErrGameScreenTooNarrow = fmt.Errorf("the game screen should be larger than width*height %d*%d", MinWidth, MinHeight)
	defaultMarshal         = json.Marshal
)

type Tetris struct {
	*gameController
	marshal func(interface{}) ([]byte, error)
}

// new tetris game with default json serialization
func NewTetris(height, width int) (*Tetris, error) {
	if height < MinHeight || width < MinWidth {
		return nil, ErrGameScreenTooNarrow
	}
	gc := newGameControllerWithDefaultSetting(height, width)
	return &Tetris{
		gameController: gc,
		marshal:        defaultMarshal,
	}, nil
}

// new tetris game with user defined serialization
func NewTetrisWithSerialization(height, width int, marshal func(interface{}) ([]byte, error)) (*Tetris, error) {
	if height < MinHeight || width < MinWidth {
		return nil, ErrGameScreenTooNarrow
	}
	gc := newGameControllerWithDefaultSetting(height, width)
	return &Tetris{
		gameController: gc,
		marshal:        marshal,
	}, nil
}

func (t *Tetris) IsGameOver() bool {
	return <-t.gameOverChan
}

func (t *Tetris) Input(key string) {
	t.inputChan <- key
}

func (t *Tetris) Score() int {
	return <-t.scoreChan
}

func (t *Tetris) ComboScore() int {
	return <-t.comboScoreChan
}

func (t *Tetris) NextPiece() interface{} {
	return <-t.nextPieceChan
}

func (t *Tetris) GameScreen() interface{} {
	return <-t.zoneChan
}
