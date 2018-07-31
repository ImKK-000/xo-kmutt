package api

import (
	"encoding/json"
	"game"
	"net/http"
)

type InfoResponse struct {
	Grid    [][]string `json:"grid"`
	PlayerX string     `json:"playerX"`
	PlayerO string     `json:"playerO"`
	Turnof  string     `json:"turnof"`
}
type Players struct {
	PlayerX string `json:"playerX"`
	PlayerO string `json:"playerO"`
}

func GetInfoHandle(write http.ResponseWriter, request *http.Request) {
	infoJson, _ := json.Marshal(GetInfo())
	write.Write(infoJson)
}

func StartGameHandler(write http.ResponseWriter, request *http.Request) {
	var players Players
	err := json.NewDecoder(request.Body).Decode(&players)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
	}
	game.NewGame(players.PlayerX, players.PlayerO)
	write.WriteHeader(http.StatusOK)
}
