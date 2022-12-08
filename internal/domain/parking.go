package domain

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type ParkingUsecase interface {
	GetParking(key string) (interface{}, error)
}

type ParkingRepository struct {
	redis *redis.Client
}

func NewParkingRepository(redis *redis.Client) *ParkingRepository {
	return &ParkingRepository{
		redis: redis,
	}
}

func (pr *ParkingRepository) GetParking(key string) (interface{}, error) {
	return pr.redis.Get(context.Background(), key).Result()
}
