package main

func main() {
	println("Welcome to cli-chess")
	println("♜♞♝♛♚♟♖♘♗♕♔♙")

	board := NewBoard()
	setStartingBoard(&board)
	board.Draw()
}

func getStartingPieceSet(c Color) PieceSet {
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

func placeStartingPieceSet(p PieceSet, b *Board, i int) {
	for _, v := range p {
		b.PieceToSquare(*v, i)
		i++
	}
}

func setStartingBoard(b *Board) {
	placeStartingPieceSet(getStartingPieceSet(White), b, 0)
	placeStartingPieceSet(getStartingPieceSet(Black), b, 48)
}
