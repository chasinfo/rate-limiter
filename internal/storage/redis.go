package storage

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStorage(addr string, password string, db int) *RedisStorage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStorage{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *RedisStorage) SetRateLimit(key string, limit int, duration time.Duration) error {
	return r.client.Set(r.ctx, key, limit, duration).Err()
}

func (r *RedisStorage) GetRateLimit(key string) (int, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	var limit int
	_, err = fmt.Sscanf(val, "%d", &limit)
	return limit, err
}

func (r *RedisStorage) DecrementRateLimit(key string) error {
	_, err := r.client.Decr(r.ctx, key).Result()
	return err
}

func (r *RedisStorage) Close() error {
	return r.client.Close()
}