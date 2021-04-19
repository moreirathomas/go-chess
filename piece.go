package main

import (
	"fmt"
)

type Type int

const (
	None Type = iota
	King
	Pawn
	Knight
	Bishop
	Rook
	Queen
)

type Color int

const (
	White Color = iota + 8
	Black       = White + 8
)

type Piece struct {
	// The unique code associated with the piece's color and type combination
	// using bitwise XOR operation.
	//
	// For example, a White Pawn is 8 ^ 2 which equals 10.
	Bitcode int
	Color
	Type
}

// unicodes is a map between a piece's bitcode and its unicode representation.
var unicodes = map[int]string{
	8: " ", 9: "♚", 10: "♟", 11: "♞", 12: "♝", 13: "♜", 14: "♛",
	16: " ", 17: "♔", 18: "♙", 19: "♘", 20: "♗", 21: "♖", 22: "",
}

type PieceSet [16]*Piece

// NewPiece returns a new Piece given its color and type.
func NewPiece(c Color, t Type) *Piece {
	return &Piece{
		Bitcode: int(c) ^ int(t),
		Color:   c,
		Type:    t,
	}
}

// NewPieceFromCode returns a new Piece given its bitcode.
func NewPieceFromCode(b int) *Piece {
	// A piece having a bitcode smaller than 16 is a white piece, as White Queen is 14 and Black None is 16.
	if b < 16 {
		return NewPiece(White, Type(8^b))
	} else {
		return NewPiece(Black, Type(16^b))
	}
}

// Name returns the piece's color and type combination as a human readable name.
//
// For example, "White Pawn".
func (p Piece) Name() string {
	return fmt.Sprintf("%s %s", p.Color.String(), p.Type.String())
}

// Unicode returns the piece's unicode representation.
//
// For example, "♟" for a White Pawn.
func (p Piece) Unicode() string {
	return fmt.Sprintf("%1s", unicodes[p.Bitcode])
}

// NewPieceSet returns a set of the 16 starting pieces for a given color.
// The pieces a ordered as their starting position on the board would be.
func NewPieceSet(c Color) PieceSet {
	king := NewPiece(c, King)
	pawn := NewPiece(c, Pawn)
	knight := NewPiece(c, Knight)
	bishop := NewPiece(c, Bishop)
	rook := NewPiece(c, Rook)
	queen := NewPiece(c, Queen)

	pieces := [16]*Piece{
		rook, knight, bishop, king, queen, bishop, knight, rook,
		pawn, pawn, pawn, pawn, pawn, pawn, pawn, pawn,
	}

	// Blacks pieces are inverted
	if c > White {
		for i, j := 0, len(pieces)-1; i < j; i, j = i+1, j-1 {
			pieces[i], pieces[j] = pieces[j], pieces[i]
		}
	}
	return pieces
}
