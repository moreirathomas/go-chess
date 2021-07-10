package chess

import (
	"fmt"
)

// Type is the type of the piece.
type Type int

const (
	None Type = iota // None represents the absence of a piece.
	King
	Pawn
	Knight
	Bishop
	Rook
	Queen
)

// Color is the color of the piece.
type Color int

const (
	White Color = iota + 8
	Black       = White + 8
)

type Piece struct {
	// Bitcode is the unique code associated with the piece's
	// color and type combination using bitwise XOR operation.
	//
	// For example, a White Pawn is 8 ^ 2 which equals 10.
	Bitcode int
}

// unicodes is a map between a piece's bitcode and its unicode representation.
var unicodes = map[int]string{
	8: " ", 9: "♚", 10: "♟", 11: "♞", 12: "♝", 13: "♜", 14: "♛",
	16: " ", 17: "♔", 18: "♙", 19: "♘", 20: "♗", 21: "♖", 22: "♕",
}

// NewPiece returns a new Piece given its bitcode.
func NewPiece(b int) *Piece {
	// Default any invalid bitcode to None.
	if b != 0 && (b < 8 || b > 22 || b == 15) {
		return &Piece{Bitcode: int(None)}
	}
	return &Piece{Bitcode: b}
}

// Color returns the color of the piece.
func (p *Piece) Color() Color {
	if p.Bitcode < 16 {
		return White
	}
	return Black
}

// Type returns the type of the piece.
func (p *Piece) Type() Type {
	return Type(p.Color() ^ Color(p.Bitcode))
}

// NewPieceExplicit returns a new Piece given its color and its type.
func NewPieceExplicit(c Color, t Type) *Piece {
	return &Piece{Bitcode: int(c) ^ int(t)}
}

// Name returns the piece as a human readable string.
//
// For example, "White Pawn".
func (p Piece) Name() string {
	return fmt.Sprintf("%s %s", p.Color().String(), p.Type().String())
}

// Unicode returns the piece's unicode representation.
//
// For example, "♟" for a White Pawn.
func (p Piece) Unicode() string {
	return fmt.Sprintf("%1s", unicodes[p.Bitcode])
}

// IsColor returns true if a piece is a given color.
func (p Piece) IsColor(c Color) bool {
	return p.Color() == c
}

// NewPieceSet returns a set of the 16 starting pieces for a given color.
// The pieces are ordered by their starting position on the board.
func NewPieceSet(c Color) [16]*Piece {
	var set [16]*Piece

	king := NewPieceExplicit(c, King)
	pawn := NewPieceExplicit(c, Pawn)
	knight := NewPieceExplicit(c, Knight)
	bishop := NewPieceExplicit(c, Bishop)
	rook := NewPieceExplicit(c, Rook)
	queen := NewPieceExplicit(c, Queen)

	kingRank := [8]*Piece{rook, knight, bishop, queen, king, bishop, knight, rook}
	pawnRank := [8]*Piece{pawn, pawn, pawn, pawn, pawn, pawn, pawn, pawn}

	if c == White {
		_ = append(append(set[:0], kingRank[:]...), pawnRank[:]...)
	} else {
		_ = append(append(set[:0], pawnRank[:]...), kingRank[:]...)
	}

	return set
}
