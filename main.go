package main

func main() {
	println("\nWelcome to cli-chess\n")

	board := NewBoard()

	board.PlaceStartingPieces()
	board.Draw()
}
