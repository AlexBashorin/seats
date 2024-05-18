package main

import (
	"log"
	"net/http"
	"seats/internal/webserver"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", webserver.Mainpage)
	mux.HandleFunc("/setUser", webserver.SetUser)

	log.Fatal(http.ListenAndServe(":94", mux))
}
