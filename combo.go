package tetris

import "time"

const comboInterval = 1e9

type combo struct {
	score     int
	comboTime int64
}

func (c *combo) comboWith(c1 combo) *combo {
	if c.comboTime-c1.comboTime <= comboInterval {
		c.score += c1.score
	}
	return c
}

type combos []*combo

func newCombos() *combos {
	c := make(combos, 0)
	return &c
}

func (cs *combos) newCombo(score int) {
	*cs = append(*cs, &combo{score: score, comboTime: time.Now().UnixNano()})
}

func (cs combos) comboScore() int {
	l := len(cs)
	score := cs[l-1].score
	for i := l - 1; i > 0; i-- {
		if cs[i].comboTime-cs[i-1].comboTime <= comboInterval {
			score += cs[i-1].score
		} else {
			break
		}
	}
	return score
}
