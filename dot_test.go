package tetris

import "testing"

// testing
type testContiguousStruct struct {
	dot1, dot2 *dot
	pass       bool
}

var testContiguousData = []testContiguousStruct{
	testContiguousStruct{newDot(0, 0), newDot(0, 1), true},
	testContiguousStruct{newDot(0, 0), newDot(1, 0), true},
	testContiguousStruct{newDot(0, 0), newDot(0, 2), false},
	testContiguousStruct{newDot(0, 0), newDot(1, 1), false},
	testContiguousStruct{newDot(1, 1), newDot(0, 1), true},
	testContiguousStruct{newDot(1, 1), newDot(2, 1), true},
}

func testDot(t *testing.T) {
	// test contiguous
	for _, v := range testContiguousData {
		if v.dot1.isContiguous(*v.dot2) != v.pass {
			t.Errorf("something wrong %v %v", v.dot1, v.dot2)
		}
	}
	// test generate dots
	for i := 0; i < 100; i++ {
		ds := newDots()
		// t.Logf("dots: %v with center %v", ds, ds.center())
		ds.rotate()
		if ds.hasNegativeDot() {
			t.Logf("dots after rotation: %v", ds)
		}
		ds.rotate()
		if ds.hasNegativeDot() {
			t.Logf("dots after rotation: %v", ds)
		}
		ds.rotate()
		if ds.hasNegativeDot() {
			t.Logf("dots after rotation: %v", ds)
		}
		ds.rotate()
		if ds.hasNegativeDot() {
			t.Logf("dots after rotation: %v", ds)
		}
	}
}

func TestDot(t *testing.T) {
	testDot(t)
}
