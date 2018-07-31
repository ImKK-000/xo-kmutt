package api

import (
	"encoding/json"
	"net/http"
)

type InfoResponse struct {
	Grid    [][]string `json:"grid"`
	PlayerX string     `json:"playerX"`
	PlayerO string     `json:"playerO"`
	Turnof  string     `json:"turnof"`
}

func GetInfoHandle(write http.ResponseWriter, request *http.Request) {
	infoJson, _ := json.Marshal(GetInfo())
	write.Write(infoJson)
}
