package handlers

import (

	"net/http"
	"encoding/json"
)

type User struct {

	Id uint16		`json:"id"`
	Username string `json:"username"`
}

var users []User

func Users (w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(users)
}