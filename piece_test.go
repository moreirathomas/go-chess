package chess_test

import (
	"testing"

	"github.com/moreirathomas/go-chess"
)

func TestPiece(t *testing.T) {
	expected := map[string]int{
		"White None":   8,
		"White King":   9,
		"White Pawn":   10,
		"White Knight": 11,
		"White Bishop": 12,
		"White Rook":   13,
		"White Queen":  14,
		"Black None":   16,
		"Black King":   17,
		"Black Pawn":   18,
		"Black Knight": 19,
		"Black Bishop": 20,
		"Black Rook":   21,
		"Black Queen":  22,
	}

	for name, bitcode := range expected {
		p := chess.NewPiece(bitcode)
		if p.Name() != name {
			t.Fatalf("bad piece for bitcode %d, want %s got %s", bitcode, name, p.Name())
		}
	}
}
