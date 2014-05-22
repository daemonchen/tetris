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

// can move down?
// simply check if there is any dot under the piece
func (p *piece) canMoveDown(data [][]bool) bool {
	for _, d := range p.getDots() {
		if data[d.yCoor+1][d.xCoor] {
			return false
		}
	}
	return true
}

// can move left?
func (p *piece) canMoveLeft(data [][]bool) bool {
	for _, d := range p.getDots() {
		// cautious index out of range
		if d.xCoor == 0 {
			return false
		}
		if data[d.yCoor][d.xCoor-1] {
			return false
		}
	}
	return true
}

// can move right?
func (p *piece) canMoveRight(data [][]bool) bool {
	for _, d := range p.getDots() {
		// cautious index out of range
		if d.xCoor >= len(data[0])-1 {
			return false
		}
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
	if ds.hasNegativeDot() {
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
