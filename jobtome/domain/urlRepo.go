package domain

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type UrlRepo struct {
}
type StorageService struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewUrlRepoInit() *UrlRepo {
	return &UrlRepo{}
}

type UrlRepoInt interface {
	GetKey(value Urls) (Urls, error)
	SetKey(value Urls, expiration time.Duration) (Urls, error)
	GetAllValues() map[string]string
}

func (u *UrlRepo) initializeRedis() *StorageService {
	var ctx = context.Background()
	storageService := &StorageService{}
	// dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	c := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	if err := c.Ping(ctx).Err(); err != nil {
		log.Println("Unable to connect to redis " + err.Error())
		return nil
	}
	storageService.redisClient = c
	storageService.ctx = ctx
	return storageService
}

//GetKey get value from redis
func (u *UrlRepo) GetKey(value Urls) (Urls, error) {
	var urlReq Urls
	client := u.initializeRedis().redisClient
	val, err := client.Get(u.initializeRedis().ctx, value.ShortUrl).Result()
	if err == redis.Nil || err != nil {
		return urlReq, err
	}
	err = json.Unmarshal([]byte(val), &urlReq)
	if err != nil {
		return urlReq, err
	}

	return urlReq, nil
}

//SetKey save or upd value on redis
func (u *UrlRepo) SetKey(value Urls, expiration time.Duration) (Urls, error) {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return Urls{}, err
	}
	client := u.initializeRedis().redisClient
	err = client.Set(u.initializeRedis().ctx, value.ShortUrl, cacheEntry, expiration).Err()
	if err != nil {
		return Urls{}, err
	}
	return value, nil
}

//GetAllValues retrieves all values from redis
func (u *UrlRepo) GetAllValues() map[string]string {
	client := u.initializeRedis().redisClient
	iterKeys := client.Scan(u.initializeRedis().ctx, 0, "*", 0).Iterator()
	allValues := make(map[string]string)
	for iterKeys.Next(u.initializeRedis().ctx) {
		val, _ := client.Get(u.initializeRedis().ctx, iterKeys.Val()).Result()
		allValues[iterKeys.Val()] = val
	}

	return allValues
}
