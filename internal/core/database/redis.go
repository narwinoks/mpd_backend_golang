package database

import (
	"backend-app/config"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	if !cfg.Redis.Enabled {
		return nil, nil
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// Test connection
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("failed to connect to redis: %v", err)
		return nil, err
	}

	logrus.Infof("Redis connection established: %s:%d", cfg.Redis.Host, cfg.Redis.Port)
	return rdb, nil
}
