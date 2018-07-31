package game

type Player struct {
	Name   string
	Symbol string
}
type Board struct {
	Grid [][]string
	Size int
}
type Game struct {
	Board   Board
	PlayerX Player
	PlayerO Player
	TurnOf  Player
}

func NewGame(playerXName, playerOName string) Game {
	playerX := NewPlayer(playerXName, symbol)
	playerO := NewPlayer(playerOName, symbol)
	board := NewBoard(size)
	return Game{
		PlayerX: playerX,
		PlayerO: playerO,
		Board:   board,
		TurnOf:  playerX,
	}
}
