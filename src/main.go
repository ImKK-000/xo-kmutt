package main

import "net/http"

func main() {
	http.Handle("/xogame/", http.StripPrefix("/xogame/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":3000", nil)
}
