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
