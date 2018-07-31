package game

type Game struct {
	Board   Board
	PlayerX Player
	PlayerO Player
	TurnOf  Player
}

func NewGame(playerXName, playerOName string) Game {
	playerX := NewPlayer(playerXName, "X")
	playerO := NewPlayer(playerOName, "O")
	board := NewBoard(3)
	return Game{
		PlayerX: playerX,
		PlayerO: playerO,
		Board:   board,
		TurnOf:  playerX,
	}
}

func CheckVertical(symbol string, row, column int) bool {
	return true
}
