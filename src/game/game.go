package game

type Game struct {
	Board   Board
	PlayerX Player
	PlayerO Player
	TurnOf  Player
}
type InfoResponse struct {
	Grid    [][]string `json:"grid"`
	PlayerX Player     `json:"playerX"`
	PlayerO Player     `json:"playerO"`
	TurnOf  Player     `json:"turnOf"`
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

func (g Game) GetInfo() InfoResponse {
	return InfoResponse{
		Grid:    g.Board.Grid,
		PlayerX: g.PlayerX,
		PlayerO: g.PlayerO,
		TurnOf:  g.TurnOf,
	}
}

func CheckVertical(symbol string, row, column int) bool {
	return true
}
