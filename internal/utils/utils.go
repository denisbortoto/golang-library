package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "library"
	sslmode  = "disable"
)

func Connect() *sql.DB {
	/* Connect function is used to open a connection with the database */
	URL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", URL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

/* func ErrorResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	ToJson(w, struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})
}
*/
func ToJson(w http.ResponseWriter, data interface{}) {
	/* ToJson function is used to avoid code repetition.
	This function set the Content-Type and Encode the data to json format */
	w.Header().Set("Content-type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
