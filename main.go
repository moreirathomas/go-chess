package main

import "github.com/moreirathomas/go-chess/core"

func main() {
	println("\nWelcome to cli-chess\n")

	board := core.NewBoard()

	board.PlaceStartingPieces()
	board.Draw()

	i := 11

	i = board.Move(i, i+8)
	board.Draw()
	i = board.Move(i, i+8)
	board.Draw()
	i = board.Move(i, i+1)
	board.Draw()
	i = board.Move(i, i-1)
	board.Draw()
	_ = board.Move(i, i-8)
	board.Draw()

	println(core.NewPieceFromCode(board.Square[63]).Name())
}
