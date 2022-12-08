package main

import (
	"app/internal/configs"
	"app/internal/parking/fetch"
	"app/internal/routes"
	"app/pkg/redis"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := configs.GetConfig("config/config.yml")

	fmt.Println(cfg)

	if cfg == nil {
		log.Fatal("Error acure on config init")
	}

	cl, err := redis.NewClient(cfg.Redis.Addr, cfg.Redis.Password)
	if err != nil {
		log.Fatal("Error acure on redis client init: ", err.Error())
	}

	go func() {
		if err := fetch.UpdateAllEntries(cfg.Apitoken, cl); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	app := fiber.New()

	routes.RegisterRoutes(app, cl)

	log.Println("Start server on " + cfg.Listen.BindIp + ":" + cfg.Listen.Port)
	log.Fatal(app.Listen(cfg.Listen.BindIp + ":" + cfg.Listen.Port))
}
