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

type testAbsStruct struct {
	dot1, dot2 *dot
	x, y       int
	pass       bool
}

var testAbsData = []testAbsStruct{
	testAbsStruct{newDot(0, 0), newDot(0, 1), 0, 1, true},
	testAbsStruct{newDot(0, 1), newDot(0, 0), 0, -1, true},
	testAbsStruct{newDot(1, 1), newDot(2, 1), 1, 0, true},
}

func testDot(t *testing.T) {
	// test contiguous
	for _, v := range testContiguousData {
		if v.dot1.isContiguous(*v.dot2) != v.pass {
			t.Errorf("something wrong %v %v", v.dot1, v.dot2)
		}
	}
	// test abs
	for _, v := range testAbsData {
		t.Logf("testing if %v + (%d, %d) = %v", v.dot1, v.x, v.y, v.dot2)
		if v.dot1.add(v.x, v.y).isOverlapped(*v.dot2) != v.pass {
			t.Errorf("something wrong with the abs function")
		}
	}
}

func TestDot(t *testing.T) {
	testDot(t)
}
