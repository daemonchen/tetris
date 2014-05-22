package tetris

import "time"

type timer struct {
	tickInterval int
	ticker       *time.Timer
}

// tick interval default 500 ms
const defaultInitTickInterval = 500

func newTimer(tickInterval int) *timer {
	return &timer{
		tickInterval: tickInterval,
		ticker:       time.NewTimer(time.Millisecond * time.Duration(tickInterval)),
	}
}

// wait until tick
func (t *timer) wait() {
	select {
	// use ok here in case the timer is already stop
	case _, ok := <-t.ticker.C:
		if !ok {
			warning("the timer is stop")
			t.ticker = time.NewTimer(time.Millisecond * time.Duration(t.tickInterval))
			return
		}
		t.ticker.Reset(time.Millisecond * time.Duration(t.tickInterval))
	}
}

// reset tick interval
func (t *timer) resetInterval(interval int) {
	t.tickInterval = interval
}
