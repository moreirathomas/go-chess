package main

import "github.com/moreirathomas/go-chess/core"

func main() {
	println("\nWelcome to cli-chess\n")

	board := core.NewBoard()

	board.PlaceStartingPieces()
	board.Draw()
}
