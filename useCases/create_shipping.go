package usecases

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
)

func CreateShipping(newShipping ShippingDTO) (shippings []ShippingDTO) {
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	result, err := client.Get(context.Background(), "shippings").Result()
	if err != nil {
		log.Error("Cannot create new shipping")
		return
	}

	err = json.Unmarshal([]byte(result), &shippings)
	if err != nil {
		log.Error("Cannot convert shipping")
		return
	}

	shippings = append(shippings, newShipping)

	client.Set(context.Background(), "shippings", shippings, time.Hour*5)
	return
}
