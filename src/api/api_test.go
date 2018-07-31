package api_test

import (
	"io/ioutil"
	"net/http/httptest"
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
