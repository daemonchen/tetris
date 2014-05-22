package tetris

import "fmt"

const maxDotInAPiece = 4

var errExceedsMaxNumOfDots = fmt.Errorf("number of dots in a piece should not be larger than %d", maxDotInAPiece)

type line []bool

func newLine(length int) *line {
	l := make(line, length)
	return &l
}

// ease game zone dots fetching
// should be every careful of this... bug fix takes quite a long time
func (l line) getLine() []bool {
	nl := make(line, len(l))
	copy(nl, l)
	return []bool(nl)
}

// ease game over checking
// for the purpose of checking first line of game zone
func (l line) containsAnyDot() bool {
	for _, v := range l {
		if v {
			return true
		}
	}
	return false
}

func (l line) canClear() bool {
	for _, v := range l {
		if !v {
			return false
		}
	}
	return true
}

func (l *line) placeDots(indexOfDots ...int) error {
	if len(indexOfDots) > maxDotInAPiece {
		return errExceedsMaxNumOfDots
	}
	for _, index := range indexOfDots {
		(*l)[index] = true
	}
	return nil
}
