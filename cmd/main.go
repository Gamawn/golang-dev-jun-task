package main

import (
	"app/pkg/redis"
	"context"
	"log"
)

func main() {
	// for i := 0; i < 100; i++ {
	// 	data, err := fetch.GetInfo("7a5a1ed1dc55436d803a02bd9505d70e")
	// 	if err != nil {
	// 		log.Fatal("Error acure on data fetch: ", err.Error())
	// 	}
	// 	log.Println(data)
	// }

	cl, err := redis.NewClient()
	if err != nil {
		log.Fatal("Error acure on redis client init: ", err.Error())
	}

	if cmd := cl.Ping(context.Background()); cmd.Err() != nil {
		log.Fatal("Error acure on redis ping: ", cmd.Err().Error())
	}

}
