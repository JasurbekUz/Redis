package main

import (
	"log"
	"net/http"
	m "app/middlewares"
	h "app/handlers"
)

func main () {

	http.Handle("/users", m.RedisCache(h.Users))
	http.Handle("/books", m.RedisCache(h.Books))

	log.Println("server is ready at :8080")

	http.ListenAndServe(":8080", nil)
}