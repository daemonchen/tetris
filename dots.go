package tetris

import "fmt"

const defaultNumOfDotsInAPiece = 4

// dots is a collection of dot
type dots [defaultNumOfDotsInAPiece]*dot

// generate random dots
func newDots() *dots {
	var dots dots
	for i := 0; i < defaultNumOfDotsInAPiece; i++ {
		if i == 0 {
			dots[i] = randomDot()
			continue
		}
		var f = func() (*dot, bool) {
			d := randomDot()
			numOfCon := 0
			for j := 0; j < i; j++ {
				if dots[j].isOverlapped(*d) {
					return nil, false
				}
				if d.isContiguous(*dots[j]) {
					numOfCon++
				}
			}
			if numOfCon == 0 {
				return nil, false
			}
			return d, true
		}
		for {
			d, isNew := f()
			if isNew {
				dots[i] = d
				break
			}
		}
	}
	return &dots
}

// string
func (ds dots) String() (str string) {
	for i, v := range ds {
		str += fmt.Sprintf("dot%d %v\n", i, v)
	}
	return
}

// calculate the absolute location of the dots
func (ds dots) abs(x, y int) dots {
	for i, _ := range ds {
		ds[i] = ds[i].add(x, y)
	}
	return ds
}

// dots' center
func (ds dots) center() *dot {
	var x, y int
	for _, v := range ds {
		x += v.xCoor*2 + 1
		y += v.yCoor*2 + 1
	}
	return newDot(x/defaultNumOfDotsInAPiece/2, y/defaultNumOfDotsInAPiece/2)
}

// dots copy
func (ds *dots) copyDots(ds1 *dots) *dots {
	for i, d := range ds1 {
		ds[i] = d
	}
	return ds
}

// dots rotate
func (ds *dots) rotate(isRef ...bool) *dots {
	center := ds.center()
	if len(isRef) >= 1 && isRef[0] {
		for _, d := range ds {
			d.rotate(center, true)
		}
	}
	ds1 := &dots{}
	for _, d := range ds1.copyDots(ds) {
		d.rotate(center, false)
	}
	return ds1
}

// if the dots goes outbound?
func (ds dots) isOutBoundedY(data [][]bool) bool {
	lHeight := len(data)
	for _, v := range ds {
		if v.yCoor >= lHeight-1 {
			return true
		}
	}
	return false
}

// if the dots goes outbound in x-axis
func (ds dots) isOutBoundedX(data [][]bool) bool {
	lWidth := len(data[0])
	for _, v := range ds {
		if v.xCoor >= lWidth-1 || v.xCoor <= 0 {
			return true
		}
	}
	return false
}

// some dots have negative y-coor?
func (ds *dots) hasNegativeDot() bool {
	for _, d := range ds {
		if d.yCoor < 0 {
			return true
		}
	}
	return false
}

// is dots equal to another one
func (ds dots) isEqualTo(ds1 dots) bool {
	for i, v := range ds {
		if !v.isContiguous(*(ds1[i])) {
			return false
		}
	}
	return true
}
