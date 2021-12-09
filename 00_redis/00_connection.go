package main

import (
	"log"
	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "go-redis",
	DB: 0,
})

func main () {

	pong, err := client.Ping().Result()

	defer client.Close()

	if err != nil {

		log.Fatalf("connection error: %v", err)
	}

	if err = client.Set("color", "white", 0).Err(); err != nil {

		log.Fatal(err)
	}

	log.Println(pong)

	color, err := client.Get("color").Result()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(color)

}

