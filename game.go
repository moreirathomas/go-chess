package chess

var (
	friendlyColor Color
	enemyColor    Color
)

// ChooseColor sets the friendly and enemy colors' value for the game.
func ChooseColor(c Color) {
	friendlyColor = c
	switch c {
	case White:
		enemyColor = Black
	case Black:
		enemyColor = White
	}
}
