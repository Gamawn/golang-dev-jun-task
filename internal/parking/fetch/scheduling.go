package fetch

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

func UpdateAllEntries(apiToken string, client *redis.Client) error {
	for {
		data, err := GetInfo(apiToken)
		if err != nil {
			return err
		}
		if data == nil {
			time.Sleep(time.Second * 5)
			continue
		}

		for _, d := range *data {
			// update data
			e := client.Set(context.Background(), d.ID, d, 0)
			if err := e.Err(); err != nil {
				time.Sleep(time.Second * 2)
			}
		}

		time.Sleep(time.Second * 2)
	}
}
