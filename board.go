package main

import "fmt"

type Square [64]int

type Board struct {
	Square
}

// NewBoard returns an empty board.
func NewBoard() Board {
	return Board{
		Square: *new([64]int),
	}
}

// Draw prints a representation of the current board to the standard output.
func (b Board) Draw() {
	i := 0
	for file := 0; file < 8; file++ {
		for rank := 0; rank < 8; rank++ {
			isWhiteSquare := (file+rank)%2 == 0
			if isWhiteSquare {
				fmt.Printf("[%2d]", b.Square[i])
			} else {
				fmt.Printf("(%2d)", b.Square[i])
			}
			i++
		}
		println()
	}
}

// PieceToSquare moves the piece to the given square. Internally, the board stores the piece's
// bitcode at the given index on the array representing the squares.
func (b *Board) PieceToSquare(p Piece, i int) {
	b.Square[i] = p.Bitcode
}

func (s Square) Notation(i int) string {
	return "yo"
}
