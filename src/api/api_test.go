package api_test

import (
	. "api"
	"game"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_GetInfoHandle_Should_Be_Response_Grid_With_PlayerX_kad_PlayerO_lek_Turnof_kad(t *testing.T) {
	playerX := game.Player{
		Name:   "กาด",
		Symbol: "X",
	}
	playerO := game.Player{
		Name:   "เล็ก",
		Symbol: "O",
	}
	board := game.Board{
		Grid: [][]string{
			[]string{"X", "X", "X"},
			[]string{"O", "O", ""},
			[]string{"", "", ""},
		},
		Size: 3,
	}
	Game := game.Game{
		Board:   board,
		PlayerX: playerX,
		PlayerO: playerO,
		TurnOf:  playerX,
	}
	expected := `{"grid":[["X","X","X"],["O","O",""],["","",""]],"playerX":{"name":"กาด","symbol":"X"},"playerO":{"name":"เล็ก","symbol":"O"},"turnOf":{"name":"กาด","symbol":"X"}}`

	request := httptest.NewRequest("GET", "/xo/info", nil)
	responseRecorder := httptest.NewRecorder()

	api := API{Game: Game}
	api.GetInfoHandler(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if expected != string(actual) {
		t.Errorf("expectted is\n %s but got\n %s", expected, string(actual))
	}
}

func Test_StartGameHandler_Input_PlayerX_Kad_PlayerO_Lek_Should_Be_Http_Status_200(t *testing.T) {
	players := `{"playerX" : "กาด", "playerO" : "เล็ก"}`
	expectedStatus := 200

	req := httptest.NewRequest("POST", "/xo/start", strings.NewReader(players))
	w := httptest.NewRecorder()
	api := API{}
	api.StartGameHandler(w, req)
	response := w.Result()
	actualStatus := response.StatusCode

	if actualStatus != expectedStatus {
		t.Errorf("expected %d but got %d", expectedStatus, actualStatus)
	}
}
