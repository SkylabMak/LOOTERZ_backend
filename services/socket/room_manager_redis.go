// room_manager.go
package socket

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var RedisClient *redis.Client

// InitRedis initializes the Redis client
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-15141.c73.us-east-1-2.ec2.redns.redis-cloud.com:15141", // adjust based on your Redis server
		Password: "wDi6Tv31oPCNZ15ebCrsbmf2zGclRPMk",               // no password set
		DB:       0,                // use default DB
	})

	// Check the Redis connection
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
}

// PublishToRoom publishes a message to the specified room (Redis channel)
func PublishToRoom(roomID string, message string) error {
	err := RedisClient.Publish(ctx, roomID, message).Err()
	if err != nil {
		log.Printf("Error publishing to room %s: %v", roomID, err)
	}
	return err
}
