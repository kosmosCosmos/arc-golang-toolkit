package connect

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisConfiguration struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       int
}

// NewRedisClient creates and returns a new Redis client
func NewRedisClient(config RedisConfiguration) (*redis.Client, error) {
	if config.Port == "" {
		config.Port = "6379"
	}

	if config.DB < 0 {
		return nil, fmt.Errorf("invalid DB number: %d", config.DB)
	}

	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
	})

	// 测试连接
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return client, nil
}
