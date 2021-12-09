package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "go-redis",
	DB: 0,
})

type Book struct {

	Id uint16		`json:"id"`
	BookName string `json:"bookName"`
}

var books []Book = []Book {
	Book {
		Id: 1,
		BookName: "Ona-Tili",
	},
}

func Books (w http.ResponseWriter, r *http.Request) {

	log.Println("Executed Books Controller")

	_, err := client.Ping().Result()


	if err == nil {

		response, err := json.Marshal(books)

		if err == nil {

			client.Set(r.RequestURI, response, 0)
		}
	}

	json.NewEncoder(w).Encode(books)
}