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
