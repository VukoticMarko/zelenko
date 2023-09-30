package crdt

import (
	"sync"
)

// CRDT G-counter with addition and subtraction
type GCounter struct {
	mu         sync.Mutex
	increments map[string]int
	decrements map[string]int
}

var gCounters []*GCounter

func NewGCounter() *GCounter {
	return &GCounter{
		increments: make(map[string]int),
		decrements: make(map[string]int),
	}
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
