package tetris

import "testing"

// testing
func testDots(t *testing.T) {
	ds := newDots()
	t.Log(ds)

	ds2 := ds.abs(1, 1)
	t.Log(ds2)

	if ds.isEqualTo(ds2) {
		t.Errorf("%v should not be equal to %v", ds, ds2)
	}
}

func TestDots(t *testing.T) {
	testDots(t)
}
