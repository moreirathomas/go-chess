package main

import (
	"fmt"
)

type Square [64]int

type Board struct {
	Square
}

// files is a representation of the letter part of the squares' notation.
// It is used for efficient slicing at the required index.
const files = "abcdefgh"

// NewBoard returns an empty board.
func NewBoard() Board {
	return Board{
		Square: *new([64]int),
	}
}

// SquareNotation returns the algebraic notation of a square given its file and rank.
//
// For example, second file and third rank is "b3".
func (b Board) SquareNotation(file int, rank int) string {
	return fmt.Sprintf("%s%d", files[file-1:file], rank)
}

// IsWhiteSquare returns whether or not a square is white given its file and rank.
func (b Board) IsWhiteSquare(file int, rank int) bool {
	return (file+rank)%2 == 0
}

// Draw prints a representation of the current board to the standard output.
func (b Board) Draw() {
	// The index is manipulated this way in order to get a representation such as this:
	//
	// 8 [56][57][58][59][60][61][62][63]
	// 7 [48][49][50][51][52][53][54][55]
	// 6 [40][41][42][43][44][45][46][47]
	// 5 [32][33][34][35][36][37][38][39]
	// 4 [24][25][26][27][28][29][30][31]
	// 3 [16][17][18][19][20][21][22][23]
	// 2 [ 8][ 9][10][11][12][13][14][15]
	// 1 [ 0][ 1][ 2][ 3][ 4][ 5][ 6][ 7]
	//    a   b   c   d   e   f   g   h

	i := 64
	for file := 8; file > 0; file-- {
		i -= 8
		for rank := 8; rank > 0; rank-- {
			p := NewPieceFromCode(b.Square[i])
			fmt.Printf("[%2d:%s:%t:%s]", i, b.SquareNotation(file, rank), b.IsWhiteSquare(rank, file), p.Unicode())
			i++
		}
		i -= 8
		println()
	}
	println()
}

// PieceToSquare moves the piece to the given square. Internally, the board stores the piece's
// bitcode at the given index on the array representing the squares.
func (b *Board) PieceToSquare(p Piece, i int) {
	b.Square[i] = p.Bitcode
}

// PlaceStartingPieces places the players pieces at their starting position on the board.
func (b *Board) PlaceStartingPieces() {
	place := func(p PieceSet, b *Board, i int) {
		for _, v := range p {
			b.PieceToSquare(*v, i)
			i++
		}
	}

	place(NewPieceSet(White), b, 0)
	place(NewPieceSet(Black), b, 48)
}
