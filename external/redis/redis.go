package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gitlab.com/farkroft/auth-service/external/config"
	"gitlab.com/farkroft/auth-service/external/constants"
)

type Repository interface {
	Set(ctx context.Context, key string, value interface{}, exp int, duration string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, keys ...string) error
}

type Redis struct {
	Client *redis.Client
}

func NewRedis(cfg *config.Config) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.GetString(constants.EnvRedisHost),
		Password: cfg.GetString(constants.EnvRedisPass),
		DB:       cfg.GetInt(constants.EnvRedisDB),
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("redis client: %s", err.Error())
	}

	return &Redis{Client: rdb}
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, exp int, duration string) error {
	var dur time.Duration
	switch duration {
	case "hour":
		dur = time.Hour
	case "min":
		dur = time.Minute
	case "sec":
		dur = time.Second
	}

	err := r.Client.Set(ctx, key, value, dur*time.Duration(exp)).Err()

	return err
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	res, err := r.Client.Get(ctx, key).Result()

	return res, err
}

func (r *Redis) Delete(ctx context.Context, keys ...string) error {
	return r.Client.Del(ctx, keys...).Err()
}
