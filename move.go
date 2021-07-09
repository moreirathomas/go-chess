package chess

type Move struct {
	StartSquare  int
	TargetSquare int
}

func (b Board) NewMoves() []Move {
	moves := []Move{}

	for i, bitcode := range b.Square {
		p := NewPieceFromCode(bitcode)
		if p.IsColor(b.ColorToMove) {
			// switch p.Type {
			// }
			moves = b.generateSlidingMove(moves, i, *p)
		}
	}

	return moves
}

func (b Board) generateSlidingMove(moves []Move, squareIndex int, piece Piece) []Move {
	for dirIndex := 0; dirIndex < 8; dirIndex++ {
		for n := 0; n < b.SquaresToEdge[squareIndex][dirIndex]; n++ {
			targetIndex := squareIndex + MoveOffsets[dirIndex]*(n+1)

			pieceOnTarget := NewPieceFromCode(b.Square[targetIndex])

			// Friendly piece blocks.
			if pieceOnTarget.IsColor(friendlyColor) {
				break
			}

			moves = append(moves, Move{squareIndex, targetIndex})

			// Enemy piece is captured, we stop.
			if pieceOnTarget.IsColor(enemyColor) {
				break
			}
		}
	}
	return moves
}

// func (b Board) generateShortMove() {}
