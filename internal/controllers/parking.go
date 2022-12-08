package controllers

import (
	"app/internal/domain"

	"github.com/go-redis/redis/v9"
	fiber "github.com/gofiber/fiber/v2"
)

type ParkingController struct {
	usecase domain.ParkingUsecase
}

func NewParkingController(usecase domain.ParkingUsecase) *ParkingController {
	return &ParkingController{
		usecase: usecase,
	}
}

func (pc *ParkingController) GetParking(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := pc.usecase.GetParking(id)
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "not found",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}
