package webserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"seats/internal"
	// "seats/internal/types"
)

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
}

func Mainpage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test1"))
}

// set user
func SetUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	db := internal.ConnectDB()
	insertUser := "INSERT INTO users(id, name, company) VALUES(?, ?, ?)"
	_, err = db.Exec(insertUser, user.Id, user.Name, user.Company)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte("write: " + user.Name))
}
