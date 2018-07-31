package game_test

import (
	. "game"
	"testing"
)

func Test_NewBoard_Input_3_Should_Be_Board_With_Grid_3x3(t *testing.T) {
	size := 3
	expectedBoard := Board{
		Grid: [][]string{
			[]string{"", "", ""},
			[]string{"", "", ""},
			[]string{"", "", ""},
		},
		Size: 3,
	}
	actualBoard := NewBoard(size)
	for row := 0; row < expectedBoard.Size; row++ {
		for column := 0; column < expectedBoard.Size; column++ {
			if actualBoard.Grid[row][column] != expectedBoard.Grid[row][column] {
				t.Errorf("expected %v but got %v", expectedBoard, actualBoard)
			}
		}
	}
}
