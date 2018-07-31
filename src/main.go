package main

import (
	apiLibrary "api"
	"net/http"
)

func main() {
	api := apiLibrary.API{}
	http.HandleFunc("/xo/start", api.StartGameHandler)
	http.Handle("/xogame/", http.StripPrefix("/xogame/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":3000", nil)
}
