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

func StartGameHandler(w http.ResponseWriter, r *http.Request) {
	var players Players
	err := json.NewDecoder(r.Body).Decode(&players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	game.NewGame(players.PlayerX, players.PlayerO)
	w.WriteHeader(http.StatusOK)
}
