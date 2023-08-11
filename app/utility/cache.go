package utility

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

func GetDataFromCache[T any](ctx *gin.Context, cache *RedisClient, service func(ctx2 context.Context, id string) (T, error), id, cacheKey string) error {

	// get data from cache
	bytes, err := cache.GetCache(ctx, cacheKey)
	if err != nil {
		// if the error is redis: nil, just ignore it and fetch from the db
		if err.Error() != "redis: nil" {
			sentry.CaptureException(err)
			return err
		}
	}

	if bytes != nil {
		var data T
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			sentry.CaptureException(err)
			return err
		}
		if fmt.Sprintf("%T", data) == "[]*ent.Business" {
			ctx.JSON(http.StatusOK, gin.H{"businessAccounts": data})
			return nil
		}
		ctx.JSON(http.StatusOK, data)
		return nil
	}

	data, err := service(ctx, id)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	// Cache the data before returning
	cache.SetCache(ctx, cacheKey, data)

	if fmt.Sprintf("%T", data) == "[]*ent.Business" {
		ctx.JSON(http.StatusOK, gin.H{"businessAccounts": data})
		return nil
	}

	ctx.JSON(http.StatusOK, data)
	return nil
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
