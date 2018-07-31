package api_test

import (
	. "api"
	"net/http/httptest"
	"strings"
	"testing"
)

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
