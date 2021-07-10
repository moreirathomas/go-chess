package chess_test

import (
	"testing"

	"github.com/moreirathomas/go-chess"
)

func TestBoard(t *testing.T) {
	// White is first, Black second in the array representation.
	// When drawing the board, we loop in a reversed order.
	expected := [64]int{
		13, 11, 12, 14, 9, 12, 11, 13,
		10, 10, 10, 10, 10, 10, 10, 10,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		18, 18, 18, 18, 18, 18, 18, 18,
		21, 19, 20, 22, 17, 20, 19, 21,
	}

	board := chess.NewBoard()
	board.PlacePieces()

	for i, p := range board.Square {
		e := chess.NewPiece(expected[i])
		if p != e.Bitcode {
			t.Fatalf("bad piece in squares array at index %d, want %d got %d", i, e, p)
		}
	}
}
