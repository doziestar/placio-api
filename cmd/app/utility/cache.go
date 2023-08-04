package utility

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	host   string
	db     int
	exp    time.Duration
	client *redis.Client
}

func NewRedisClient(host string, db int) *RedisClient {
	return &RedisClient{
		host: host,
		db:   db,
		exp:  24 * time.Hour,
	}
}

func (r *RedisClient) ConnectRedis() *redis.Client {
	opt, _ := redis.ParseURL("rediss://default:a3677c1a7b84402eb34efd55ad3cf059@golden-colt-33790.upstash.io:33790")
	//client := redis.NewClient(opt)
	newClient := redis.NewClient(opt)

	pong, err := newClient.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Error connecting to redis")
	}
	log.Println("                                                        ")
	log.Println("=====  =====  =====  =====  =====  =====  =====  =======")
	log.Println("=========== Connecting To Redis on port 6379 ===========")
	log.Printf("============== Ready For Cache: %s ===================", pong)
	log.Println("=====  =====  =====  =====  =====  =====  =====  =======")
	log.Println("                                                        ")
	r.client = newClient
	return newClient
}

func (r *RedisClient) SetCache(ctx context.Context, key string, value interface{}) error {
	log.Println("Setting cache for key", key)
	if r.client == nil {
		return errors.New("redis client is nil")
	}
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	log.Println("Setting cache for key", key)
	const CacheDuration = 24 * time.Hour
	return r.client.Set(ctx, key, jsonValue, CacheDuration).Err()
}

func (r *RedisClient) GetCache(ctx context.Context, key string) ([]byte, error) {
	log.Println("Getting cache for key", key)
	if r.client == nil {
		return nil, errors.New("redis client is nil")
	}
	val, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (r *RedisClient) GetOrSetCache(ctx context.Context, key string, value interface{}) ([]byte, error) {
	if r.client == nil {
		return nil, errors.New("redis client is nil")
	}
	val, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil || val == nil {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		err = r.client.Set(ctx, key, jsonValue, 24*time.Hour).Err()
		if err != nil {
			return nil, err
		}
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

func (r *RedisClient) DeleteCache(ctx context.Context, key string) error {
	if r.client == nil {
		return errors.New("redis client is nil")
	}
	return r.client.Del(ctx, key).Err()
}
