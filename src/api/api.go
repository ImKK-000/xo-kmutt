package api

import (
	"encoding/json"
	"game"
	"net/http"
)

type Players struct {
	PlayerX string `json:"playerX"`
	PlayerO string `json:"playerO"`
}
type API struct {
	Game game.Game
}

func (api API) GetInfoHandler(write http.ResponseWriter, request *http.Request) {
	infoJson, _ := json.Marshal(api.Game.GetInfo())
	write.Write(infoJson)
}

func (api *API) StartGameHandler(write http.ResponseWriter, request *http.Request) {
	var players Players
	err := json.NewDecoder(request.Body).Decode(&players)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
	}
	api.Game = game.NewGame(players.PlayerX, players.PlayerO)
	write.WriteHeader(http.StatusOK)
}
