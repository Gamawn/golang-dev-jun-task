package routes

import (
	"app/internal/controllers"
	"app/internal/domain"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, redis *redis.Client) {
	controllers := controllers.NewParkingController(domain.NewParkingRepository(redis))
	app.Get("/parking/:id", controllers.GetParking)
}
