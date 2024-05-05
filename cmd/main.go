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

	log.Fatal(http.ListenAndServe(":93", mux))
}
