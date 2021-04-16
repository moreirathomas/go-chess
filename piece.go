package main

import "fmt"

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
	// For example, a Black Knight is 16 ^ 3 which equals 19.
	Bitcode int
	Color
	Type
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
// For example, "White King".
func (p Piece) Name() string {
	return fmt.Sprintf("%s %s", p.Color.String(), p.Type.String())
}

func (p Piece) Symbol() string {
	return "♜♞♝♛♚♟♖♘♗♕♔♙"
}
