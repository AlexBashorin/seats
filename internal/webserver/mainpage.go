package webserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"seats/internal"
	"seats/internal/types"
)

func Mainpage(w http.ResponseWriter, r *http.Request) {
	spensData, err := getSpens()
	if err != nil {
		log.Fatal(err)
	}

	jsdata, err := json.Marshal(spensData)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsdata)
}

func getSpens() ([]types.Amount, error) {
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
