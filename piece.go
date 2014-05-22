package tetris

type piece struct {
	// the dots
	dots *dots
	// the x, y coordinate of the piece
	x, y int
}

func newPiece(x, y int) *piece {
	p := &piece{
		x:    x,
		y:    y,
		dots: newDots(),
	}
	return p
}

// String
func (p *piece) String() string {
	return p.getDots().String()
}

// check if the piece is going to be outbounded on the x axis
func (p *piece) isOutBoundedX(data [][]bool) bool {
	return p.getDots().isOutBoundedX(data)
}

// check if the piece is going to be outbounded on the y axis
func (p *piece) isOutBoundedY(data [][]bool) bool {
	return p.getDots().isOutBoundedY(data)
}

// can move down?
// simply check if there is any dot under the piece
// be cautious of index out of range
func (p *piece) canMoveDown(data [][]bool) bool {
	if p.isOutBoundedY(data) {
		return false
	}
	for _, d := range p.getDots() {
		if data[d.yCoor+1][d.xCoor] {
			return false
		}
	}
	return true
}

// can move left?
func (p *piece) canMoveLeft(data [][]bool) bool {
	if p.isOutBoundedX(data) {
		return false
	}
	for _, d := range p.getDots() {
		if data[d.yCoor][d.xCoor-1] {
			return false
		}
	}
	return true
}

// can move right?
func (p *piece) canMoveRight(data [][]bool) bool {
	if p.isOutBoundedX(data) {
		return false
	}
	for _, d := range p.getDots() {
		if data[d.yCoor][d.xCoor+1] {
			return false
		}
	}
	return true
}

// can rotate?
// simply check overlapping
func (p *piece) canRotate(data [][]bool) bool {
	ds := p.dots.rotate().abs(p.x, p.y)
	if ds.hasNegativeDot() || ds.isOutBoundedX(data) || ds.isOutBoundedY(data) {
		return false
	}
	for _, d := range ds {
		if data[d.yCoor][d.xCoor] {
			return false
		}
	}
	return true
}

// place piece on board
func (p piece) placeOnZone(data [][]bool) [][]bool {
	ds := p.getDots()
	if ds.hasNegativeDot() {
		return data
	}
	for _, d := range ds {
		data[d.yCoor][d.xCoor] = true
	}
	return data
}

// get the absolute location of dots
func (p piece) getDots() dots {
	return p.dots.abs(p.x, p.y)
}

// rotate
func (p *piece) rotate() {
	p.dots.rotate(true)
}

// move down
func (p *piece) moveDown() *piece {
	p.y += 1
	return p
}

// move left
func (p *piece) moveLeft() *piece {
	p.x -= 1
	return p
}

// move right
func (p *piece) moveRight() *piece {
	p.x += 1
	return p
}
