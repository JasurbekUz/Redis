package main

import (

	"log"
	"encoding/json"
	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "go-redis",
	DB: 0,
})

type User struct {
	Name string
	Age uint8
}

var user = User {
	Name: "ALi",
	Age: 22,
}

func main () {

	pin, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(user)

	if err = client.Set("user", data, 0).Err(); err != nil {

		log.Fatal(err)
	}

	r_user := User{}

	r_data, _ := client.Get("user").Result()

	err = json.Unmarshal([]byte(r_data), &r_user)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(r_user.Name)
	log.Println(r_user.Age)

	log.Println(pin)
}