package game_test

import (
	. "game"
	"testing"
)

func Test_NewPlayer_Input_Name_kad_Symbol_X_Should_Be_Player_With_Name_kad_Symbol_X(t *testing.T) {
	name := "กาด"
	symbol := "X"
	expected := Player{
		Name:   "กาด",
		Symbol: "X",
	}

	actual := NewPlayer(name, symbol)

	if expected != actual {
		t.Errorf("expected is %v but got %v", expected, actual)
	}
}
func Test_NewPlayer_Input_Name_lek_Symbol_O_Should_Be_Player_With_Name_lek_Symbol_O(t *testing.T) {
	name := "เล็ก"
	symbol := "O"
	expected := Player{
		Name:   "เล็ก",
		Symbol: "O",
	}

	actual := NewPlayer(name, symbol)

	if expected != actual {
		t.Errorf("expected is %v but got %v", expected, actual)
	}
}
