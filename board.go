package chess

import (
	"fmt"
	"strings"

	"github.com/moreirathomas/go-chess/pkg/color"
	"github.com/moreirathomas/go-chess/pkg/num"
)

type Board struct {
	Square        [64]int    // Square stores the board's state. Each square holds the bitcode of its current piece.
	SquaresToEdge [64][8]int // SquaresToEdge is how many squares separate a given square from the edge in all directions.
	ColorToMove   Color
}

// fileNotation is a representation of the letter part of the squares' notation.
// It is used for efficient slicing at the required index.
var fileNotation = [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}

// NewBoard returns an empty board.
func NewBoard() Board {
	return Board{
		Square: *new([64]int),
	}
}

// squareNotation returns the algebraic notation of a square given its file and rank.
//
// For example, second file and third rank is "b3".
// func squareNotation(file, rank int) string {
// 	return fmt.Sprintf("%s%d", fileNotation[file], rank)
// }

// isWhiteSquare returns true if a square is white given its file and rank.
func isWhiteSquare(file, rank int) bool {
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
	for rank := 8; rank > 0; rank-- {
		i -= 8
		fmt.Printf("%d ", rank)
		for file := 0; file < 8; file++ {
			p := NewPiece(b.Square[i])
			fmt.Printf("%s", squareRepresentation(rank, file, *p))
			i++
		}
		i -= 8
		println()
	}
	print("  ")
	for file := 0; file < 8; file++ {
		fmt.Printf(" %s ", fileNotation[file])
	}
	println()
}

// squareRepresentation returns a printable representation of a square.
func squareRepresentation(rank int, file int, p Piece) string {
	var sb strings.Builder
	var s []string
	var c Color

	if isWhiteSquare(rank, file) {
		c = White
	} else {
		c = Black
	}

	if c == White {
		s = []string{color.White("["), p.Unicode(), color.White("]")}
	} else {
		s = []string{color.Black("["), p.Unicode(), color.Black("]")}
	}

	for _, v := range s {
		sb.WriteString(v)
	}

	return sb.String()
}

// PlacePieces sets the players pieces at their starting position on the board.
func (b *Board) PlacePieces() {
	// i is the starting index for a given set of pieces
	place := func(b *Board, ps [16]*Piece, i int) {
		for _, p := range ps {
			b.Square[i] = p.Bitcode
			i++
		}
	}

	place(b, NewPieceSet(White), 0)
	place(b, NewPieceSet(Black), 48)
}

// // PieceToSquare moves the piece to the given square.
// func (b *Board) PieceToSquare(p Piece, i int) {
// 	b.Square[i] = p.Bitcode
// }

// Move moves the piece on a given square to another.
// It returns the index of the destination square.
func (b *Board) Move(from int, to int) int {
	p := NewPiece(b.Square[from])
	b.Square[from] = 0
	b.Square[to] = p.Bitcode
	return to
}

// PrecomputeMoveData computes and stores on the Board how far each square is
// from the edge of the board.
//
// It is used for quick look-up when computing available moves.
func (b *Board) PrecomputeMoveData() {
	for file := 0; file < 8; file++ {
		for rank := 0; rank < 8; rank++ {
			// From rank 1 (index 0), there are 7 squares moving north before the edge.
			// Respectively for the others directions.
			// When moving diagonally, the number of squares to the edge is the minimum
			// between the remaining squares moving vertically and horizontally.
			north := 7 - rank
			south := rank
			east := 7 - file
			west := file

			i := rank*8 + file

			b.SquaresToEdge[i] = [8]int{
				north,
				south,
				east,
				west,
				num.Min(north, west),
				num.Min(north, west),
				num.Min(north, west),
				num.Min(north, west),
			}
		}
	}
}
