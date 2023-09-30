package crdt

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// CRDT G-counter with addition and subtraction
type GCounter struct {
	mu         sync.Mutex
	increments map[string]int
	decrements map[string]int
}

var (
	gCounters       []*GCounter
	redisClientCRDT *redis.Client
	ctx2            = context.Background()
)

func NewGCounter() *GCounter {
	initRedisCRDT()
	gc := GCounter{
		increments: make(map[string]int),
		decrements: make(map[string]int),
	}
	SaveCounter(&gc)
	return &gc
}

func initRedisCRDT() {
	redisClientCRDT = redis.NewClient(&redis.Options{
		Addr:     "localhost:6381",
		Password: "",
		DB:       0,
	})

	pong, err := redisClientCRDT.Ping(ctx2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis vas pozdravlja: " + pong)
}

func SaveCounter(counter *GCounter) error {

	gcID := uuid.New().String()

	serializedData, err := counter.Serialize()
	if err != nil {
		return err
	}

	err = redisClientCRDT.Set(ctx2, gcID, serializedData, 0).Err()
	if err != nil {
		return err
	}

	return nil

}

func GetCounterDB(redisClient *redis.Client, key string) (*GCounter, error) {

	ctx := context.Background()
	serializedData, err := redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	counter, err := Deserialize(serializedData)
	if err != nil {
		return nil, err
	}

	return counter, nil
}

// +1
func (c *GCounter) Increment(objectID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.increments[objectID]++
}

// -1
func (c *GCounter) Decrement(objectID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.decrements[objectID]++
}

// Fetch current G-Counter for this specific object
func (c *GCounter) GetValue(objectID string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	value := c.increments[objectID] - c.decrements[objectID]
	return value
}

func (c *GCounter) Merge(replicas []*GCounter) {
	c.mu.Lock()
	defer c.mu.Unlock()

	mergedIncrements := make(map[string]int)

	for _, replica := range replicas {
		for objectID, increment := range replica.increments {
			mergedIncrements[objectID] += increment
		}
	}

	mergedDecrements := make(map[string]int)

	for _, replica := range replicas {
		for objectID, decrement := range replica.decrements {
			mergedDecrements[objectID] += decrement
		}
	}

	for objectID, increment := range mergedIncrements {
		c.increments[objectID] = increment
	}

	for objectID, decrement := range mergedDecrements {
		c.decrements[objectID] = decrement
	}
}

func (c *GCounter) GetKeyList() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	keys := make([]string, 0, len(c.increments))
	for key := range c.increments {
		keys = append(keys, key)
	}

	return keys
}

func (c *GCounter) Serialize() ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Funkcija za deserijalizaciju GCounter iz JSON-a
func Deserialize(data []byte) (*GCounter, error) {
	var counter GCounter
	err := json.Unmarshal(data, &counter)
	if err != nil {
		return nil, err
	}
	return &counter, nil
}
