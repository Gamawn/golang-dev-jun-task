package main

import (
	"app/internal/parking/fetch"
	"fmt"
	"log"
)

func main() {
	data, err := fetch.GetInfo("apikey")
	if err != nil {
		log.Fatal("Error acure on data fetch: ", err.Error())
	}

	fmt.Println(data)
}
