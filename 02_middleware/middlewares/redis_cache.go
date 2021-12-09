package middlewares

import (
	"log"
	"net/http"
	"github.com/go-redis/redis"
)

var client = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	Password: "go-redis",
	DB: 0,
})

func RedisCache (fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

		_, err := client.Ping().Result()

		if err != nil {

			fn(w, r)
			return
		}

		key := r.RequestURI

		response , err := client.Get(key).Result()

		if err == nil && r.Method == "GET" {

			log.Println("Executed Middleware")
		
			w.Write([]byte(response))
			return
		}
		
		fn(w, r)
	})
}