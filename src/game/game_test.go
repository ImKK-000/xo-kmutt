package game_test

import (
	. "game"
	"testing"
)

func Test_NewGame_Input_PlayerX_Kad_PlayerO_Lek_Should_Be_Game_With_Board_And_Players(t *testing.T) {
	playerX := "กาด"
	playerO := "เล็ก"
	expectedGame := Game{
		Board: Board{
			Grid: [][]string{
				[]string{"", "", ""},
				[]string{"", "", ""},
				[]string{"", "", ""},
			},
			Size: 3,
		},
		PlayerX: Player{
			Name:   playerX,
			Symbol: "X",
		},
		PlayerO: Player{
			Name:   playerO,
			Symbol: "O",
		},
	}

	actualGame := NewGame(playerX, playerO)

	if actualGame.PlayerX != expectedGame.PlayerX ||
		actualGame.PlayerO != expectedGame.PlayerO {
		t.Errorf("expected %v but got %v", expectedGame, actualGame)
	}
	for row := 0; row < expectedGame.Board.Size; row++ {
		for column := 0; column < expectedGame.Board.Size; column++ {
			if actualGame.Board.Grid[row][column] != expectedGame.Board.Grid[row][column] {
				t.Errorf("expected %v but got %v", expectedGame, actualGame)
			}
		}
	}
}

func Test_GetInfo_Should_Be_Grid_And_Player_Name_kad_And_Player_Name_lek_And_Turnof_kad(t *testing.T) {
	playerX := Player{
		Name:   "กาด",
		Symbol: "X",
	}
	playerO := Player{
		Name:   "เล็ก",
		Symbol: "O",
	}
	board := Board{
		Grid: [][]string{
			[]string{"X", "X", "X"},
			[]string{"O", "O", ""},
			[]string{"", "", ""},
		},
		Size: 3,
	}
	Game := Game{
		Board:   board,
		PlayerX: playerX,
		PlayerO: playerO,
		TurnOf:  playerX,
	}
	expected := InfoResponse{
		Grid:    board.Grid,
		PlayerX: playerX,
		PlayerO: playerO,
		TurnOf:  playerX,
	}

	actualInfo := Game.GetInfo()

	if expected.PlayerX != actualInfo.PlayerX {
		t.Errorf("expected PlayerX is %s but got %s", expected.PlayerX, actualInfo.PlayerX)
	}
	if expected.PlayerO != actualInfo.PlayerO {
		t.Errorf("expected PlayerO is %s but got %s", expected.PlayerO, actualInfo.PlayerO)
	}
	if expected.TurnOf != actualInfo.TurnOf {
		t.Errorf("expected TurnOf is %s but got %s", expected.TurnOf, actualInfo.TurnOf)
	}
	for row := 0; row < board.Size; row++ {
		for column := 0; column < board.Size; column++ {
			if actualInfo.Grid[row][column] != expected.Grid[row][column] {
				t.Errorf("expected %v but got %v", expected, actualInfo)
			}
		}
	}
}
