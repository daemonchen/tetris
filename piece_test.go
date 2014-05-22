package tetris

import "testing"

// test
func testPiece(t *testing.T) {
	p := newPiece(10, 0)
	t.Log(p)
}

func TestPiece(t *testing.T) {
	testPiece(t)
}
