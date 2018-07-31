package api_test

import (
	. "api"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_GetInfoHandle_Should_Be_Response_Grid_With_PlayerX_kad_PlayerO_lek_Turnof_kad(t *testing.T) {
	expected := `{
		"grid":[
			["X","X",""],
			["O","O",""],
			["","",""],
		]
		"playerX":"กาด",
		"playerO":"เล็ก",
		"turnof":"กาด",
	}`

	request := httptest.NewRequest("GET", "xo/info", nil)
	responseRecorder := httptest.NewRecorder()

	GetInfoHandler(responseRecorder, request)
	response := responseRecorder.Result()
	actuat, _ := ioutil.ReadAll(request.Body)

	if expected != string(actuat) {
		t.Errorf("expectted is %s but got %s", expected, string(actuat))
	}
}

func Test_StartGameHandler_Input_PlayerX_Kad_PlayerO_Lek_Should_Be_Http_Status_200(t *testing.T) {
	players := `{"playerX" : "กาด", "playerO" : "เล็ก"}`
	expectedStatus := 200

	req := httptest.NewRequest("POST", "/xo/start", strings.NewReader(players))
	w := httptest.NewRecorder()
	StartGameHandler(w, req)
	response := w.Result()
	actualStatus := response.StatusCode

	if actualStatus != expectedStatus {
		t.Errorf("expected %d but got %d", expectedStatus, actualStatus)
	}
}
