package tetris

import "sync"

const (
	inputBuffer = 1 << 7
	zoneBuffer  = 1 << 8
	scoreBuffer = 1 << 5
	pieceBuffer = 1 << 4
)

type gameController struct {
	// game zone
	zone *gameZone
	sync.Mutex
	zoneChan      chan interface{}
	nextPieceChan chan interface{}

	// timer
	timer *timer

	// score
	score     int
	scoreChan chan int

	// TODO: combo level, will implement some interesting stuff in the future
	combo          *combos
	comboScoreChan chan int

	// game over signal
	gameOverChan chan bool

	// input channel
	inputChan chan string
}

func newGameControllerWithDefaultSetting(height, width int) *gameController {
	gc := &gameController{
		zone:           newGameZone(height, width),
		timer:          newTimer(defaultInitTickInterval),
		inputChan:      make(chan string, inputBuffer),
		gameOverChan:   make(chan bool, 1),
		combo:          newCombos(),
		zoneChan:       make(chan interface{}, zoneBuffer),
		scoreChan:      make(chan int, scoreBuffer),
		nextPieceChan:  make(chan interface{}, pieceBuffer),
		comboScoreChan: make(chan int, scoreBuffer),
	}
	return gc.init()
}

func newGameControllerWithIntervalDefined(height, width, interval int) *gameController {
	gc := &gameController{
		zone:           newGameZone(height, width),
		timer:          newTimer(interval),
		inputChan:      make(chan string, inputBuffer),
		gameOverChan:   make(chan bool, 1),
		combo:          newCombos(),
		zoneChan:       make(chan interface{}, zoneBuffer),
		scoreChan:      make(chan int, scoreBuffer),
		nextPieceChan:  make(chan interface{}, pieceBuffer),
		comboScoreChan: make(chan int, scoreBuffer),
	}
	return gc.init()
}

func (gc *gameController) init() *gameController {
	// a goroutine handle game zone update
	go gc.updateGameZone()
	// a goroutine handle user input
	go gc.processingInputs()
	return gc
}

// handle game over signal
func (gc *gameController) gameOver() {
	gc.gameOverChan <- true
}

// combo score
func (gc *gameController) comboScore(score int) int {
	gc.combo.newCombo(score)
	return gc.combo.comboScore()
}

// update score
func (gc *gameController) addScore(val int) {
	gc.score += val
}

// write score chan
func (gc *gameController) writeScoreChan() {
	gc.scoreChan <- gc.score
}

// write combo score chan
func (gc *gameController) writeComboChan() {
	gc.comboScoreChan <- gc.combo.comboScore()
}

// write next piece chan
func (gc *gameController) writePieceChan() {
	gc.nextPieceChan <- gc.zone.nextPiece
}

// write game zone chan
func (gc *gameController) writeZoneChan() {
	gc.zoneChan <- gc.zone.dotsOnZone()
}

// update the gameZone
func (gc *gameController) updateGameZone() {
	for {
		gc.timer.wait()
		func() {
			gc.Lock()
			defer gc.Unlock()
			if gc.zone.isGameOver() {
				gc.gameOver()
			}
			gc.addScore(gc.comboScore(gc.zone.updateGameZone()))
			gc.writeScoreChan()
			gc.writeZoneChan()
		}()
	}
}

// processing inputs received from users
// it is not necessarily for users to input l r d
// since we can do some mapping in client side
// but the server side is fixed
// inputs(case sensitive):
/*
	"l" 		-> move active piece left,
	"r" 		-> move active piece right,
	"d" 		-> move active piece down,
	"R"		-> rotate
*/
const (
	KeyLeft   = "l"
	KeyRight  = "r"
	KeyDown   = "d"
	KeyRotate = "R"
)

func (gc *gameController) processingInputs() {
	var processFunc = func(input string) {
		gc.Lock()
		defer gc.Unlock()
		switch input {
		case KeyLeft:
			gc.zone.movePieceLeft()
		case KeyRight:
			gc.zone.movePieceRight()
		case KeyDown:
			gc.zone.movePieceDown()
		case KeyRotate:
			gc.zone.rotate()
		}
	}
	for {
		select {
		case input, ok := <-gc.inputChan:
			if !ok {
				warning("input chan is closed?")
			}
			processFunc(input)
		}
	}
}
