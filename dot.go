package tetris

import (
	"fmt"
	"math/rand"
	"time"
)

// simply a dot
type dot struct {
	xCoor, yCoor int
}

func newDot(x, y int) *dot {
	return &dot{
		xCoor: x,
		yCoor: y,
	}
}

// string
func (d dot) String() string {
	return fmt.Sprintf("x-coor: %d, y-coor: %d", d.xCoor, d.yCoor)
}

// copy dot
func (d *dot) copyDot(d1 dot) *dot {
	d.xCoor, d.yCoor = d1.xCoor, d1.yCoor
	return d
}

// rotation
// since we should test if the piece can rotate or not
// we first perform rotation without changing the dot value
// is test passed, then perform rotation again with specifying isRef = true
func (d *dot) rotate(center *dot, isRef ...bool) *dot {
	x1, y1 := center.xCoor, center.yCoor
	if len(isRef) >= 1 && isRef[0] {
		d.xCoor, d.yCoor = d.yCoor-y1+x1, x1+y1-d.xCoor
		return d
	}
	d1 := newDot(0, 0).copyDot(*d)
	d1.xCoor, d1.yCoor = d1.yCoor-y1+x1, x1+y1-d1.xCoor
	return d1
}

func (d *dot) add(x, y int, isRef ...bool) *dot {
	if len(isRef) == 1 && isRef[0] {
		d.xCoor += x
		d.yCoor += y
		return d
	}
	return newDot(d.xCoor+x, d.yCoor+y)
}

func (d *dot) moveLeft(isRef ...bool) *dot {
	return d.add(-1, 0, isRef...)
}

func (d *dot) moveRight(isRef ...bool) *dot {
	return d.add(1, 0, isRef...)
}

func (d *dot) moveUp(isRef ...bool) *dot {
	return d.add(0, -1, isRef...)
}

func (d *dot) moveDown(isRef ...bool) *dot {
	return d.add(0, 1, isRef...)
}

func (d dot) isOverlapped(d2 dot) bool {
	return d.xCoor == d2.xCoor && d.yCoor == d2.yCoor
}

func (d dot) isContiguous(d2 dot) bool {
	return d.moveLeft().isOverlapped(d2) || d.moveRight().isOverlapped(d2) || d.moveUp().isOverlapped(d2) || d.moveDown().isOverlapped(d2)
}

var randSeed = rand.New(rand.NewSource(time.Now().UnixNano()))

// generate random dot
func randomDot() *dot {
	// get a random number between 0 and defaultNumOfDotsInAPiece
	x := randSeed.Intn(defaultNumOfDotsInAPiece)
	maxY := defaultNumOfDotsInAPiece - x
	y := randSeed.Intn(maxY)
	return newDot(x, y)
}
