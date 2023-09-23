package repository

import (
	"context"
	"fmt"
	"zelenko-backend/model"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

type GreenScoreRepository interface {
	AddOne(model.GreenObject) model.GreenObject
	GetAttributeForObject(objectID string, attribute string) int64
}

type greenScoreRepository struct {
}

func NewGreenScoreRepository() GreenScoreRepository {
	initRedis()
	return &greenScoreRepository{}
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis vas pozdravlja: " + pong)
}

func (*greenScoreRepository) AddOne(greenObject model.GreenObject) model.GreenObject {

	SetAttributeForObject(greenObject.ID.String(), "Verification", greenObject.GreenScore.Verification)
	return greenObject
}

// objectID -> uuid; attribute -> greenScore.Verification; value -> number from G-Counter
func SetAttributeForObject(objectID string, attribute string, value int) error {
	err := redisClient.HSet(ctx, objectID, attribute, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func (*greenScoreRepository) GetAttributeForObject(objectID string, attribute string) int64 {
	val, err := redisClient.HGet(ctx, objectID, attribute).Int64()
	if err != nil {
		fmt.Println("Error fetching from Redis!")
		return 0
	}
	return val
}
