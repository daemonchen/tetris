package tetris

import "container/list"

type gameZone struct {
	// height width of game area
	height int
	width  int

	// the x-coor that generate new active piece
	midPoint int

	// double linked list should make implementation simpler
	// although a well designed slice should have better performance
	lines *list.List

	// active and next piece
	activePiece, nextPiece *piece
}

func newGameZone(height, width int) *gameZone {
	gz := &gameZone{
		width:    width,
		height:   height,
		midPoint: width/2 - 1,
		lines:    list.New(),
	}
	return gz.init()
}

func (gz *gameZone) init() *gameZone {
	// initialize the game zone
	for i := 0; i < gz.height; i++ {
		gz.lines.PushFront(newLine(gz.width))
	}
	// initialize piece generate
	gz.nextPiece = newPiece(gz.midPoint, 0)
	return gz.generateNextPiece()
}

// generate next piece
func (gz *gameZone) generateNextPiece() *gameZone {
	gz.activePiece = gz.nextPiece
	gz.nextPiece = newPiece(gz.midPoint, 0)
	return gz
}

// is game over
// simply check if the first line contains any dot
func (gz *gameZone) isGameOver() bool {
	return gz.lines.Front().Value.(*line).containsAnyDot()
}

// is the activePiece able to rotate?
func (gz *gameZone) canRotate(data [][]bool) bool {
	return gz.activePiece.canRotate(data)
}

// rotate the activePiece
func (gz *gameZone) rotate() {
	if gz.canRotate(gz.dotsOnLines()) {
		gz.activePiece.rotate()
	}
}

// is the activePiece able to move down?
func (gz *gameZone) canMoveDown(data [][]bool) bool {
	return gz.activePiece.canMoveDown(data)
}

// move the active piece down,
// update the lines otherwise
func (gz *gameZone) movePieceDown() bool {
	if gz.canMoveDown(gz.dotsOnLines()) {
		gz.activePiece.moveDown()
		return true
	}
	return false
}

// is the activePiece able to move left?
func (gz *gameZone) canMoveLeft(data [][]bool) bool {
	return gz.activePiece.canMoveLeft(data)
}

// move the active piece left
func (gz *gameZone) movePieceLeft() {
	if gz.canMoveLeft(gz.dotsOnLines()) {
		gz.activePiece.moveLeft()
	}
}

// is the activePiece able to move right?
func (gz *gameZone) canMoveRight(data [][]bool) bool {
	return gz.activePiece.canMoveRight(data)
}

// move the active piece right
func (gz *gameZone) movePieceRight() {
	if gz.canMoveRight(gz.dotsOnLines()) {
		gz.activePiece.moveRight()
	}
}

// dotsOnLines is for checking if we are able to
// move piece down, left, right or roate
//
// dotsOnZone is used for rendering on browser
//
// 2D array indicating the game zone data
// [height][width]bool -> more generally, [y][x]bool

func (gz gameZone) dotsOnLines() [][]bool {
	data := make([][]bool, gz.height)
	index := 0
	for e := gz.lines.Front(); e != nil; e = e.Next() {
		data[index] = e.Value.(*line).getLine()
		index++
	}
	return data
}

func (gz gameZone) dotsOnZone() [][]bool {
	return gz.activePiece.placeOnZone(gz.dotsOnZone())
}

// clear the lines filled with dots
func (gz *gameZone) clearLines() (num int) {
	for e := gz.lines.Back(); e != nil; {
		if e.Value.(*line).canClear() {
			num++
			gz.addNewLine()
			if e.Prev() != nil {
				e = e.Prev()
				gz.lines.Remove(e.Next())
				continue
			} else {
				gz.lines.Remove(e)
				break
			}
		}
		e = e.Prev()
	}
	return
}

// new line added to the front
func (gz *gameZone) addNewLine() {
	gz.lines.PushFront(newLine(gz.width))
}

// convert the piece into lines
func (gz *gameZone) convPieceIntoLines() *gameZone {
	for _, d := range gz.activePiece.getDots() {
		for l := gz.lines.Front(); l != nil; l = l.Next() {
			d.yCoor--
			if d.yCoor >= 0 {
				continue
			}
			if err := l.Value.(*line).placeDots(d.xCoor); err != nil {
				warning("can not place dots on the line: %v", err)
			}
		}
	}
	return gz
}

// update game zone, return score earned by user
func (gz *gameZone) updateGameZone() int {
	if gz.movePieceDown() {
		return 0
	}
	return gz.convPieceIntoLines().generateNextPiece().clearLines()
}
