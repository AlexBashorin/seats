package webserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"seats/internal"
	"seats/internal/types"
)

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
}

func Mainpage(w http.ResponseWriter, r *http.Request) {
	// setUser(user)
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
	insertUser := "INSERT INTO users(id, name, compnay) VALUES(?, ?, ?)"
	_, err = db.Exec(insertUser, user.Id, user.Name, user.Company)
	if err != nil {
		log.Fatal(err)
	}
}

func setSeats() ([]types.Amount, error) {
	var spens []types.Amount
	rows, err := internal.ConnectDB().Query("SELECT * FROM amount;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var spe types.Amount
		if err := rows.Scan(&spe.Date, &spe.Sum, &spe.Mov, &spe.Id); err != nil {
			return nil, fmt.Errorf("Parse: Data not found")
		}
		spens = append(spens, spe)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Data not found")
	}

	return spens, nil
}
