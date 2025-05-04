package amf

import (
	"fmt"
	"math/rand"
	"time"
)

// Shard represents a blockchain shard
type MonitoredShard struct {
	ID   string
	Load int // simulated load
}

// MonitorAndBalance simulates dynamic shard balancing
func MonitorAndBalance(shards []MonitoredShard, threshold int) []MonitoredShard {
	rand.Seed(time.Now().UnixNano())
	for i := range shards {
		// Simulate load change
		shards[i].Load = rand.Intn(100)
	}

	for i := range shards {
		if shards[i].Load > threshold {
			fmt.Printf("Shard %s is overloaded (load: %d). Splitting...\n", shards[i].ID, shards[i].Load)
			newShard := MonitoredShard{ID: shards[i].ID + "-b", Load: shards[i].Load / 2}
			shards[i].Load /= 2
			shards = append(shards, newShard)
		} else if shards[i].Load < threshold/4 && len(shards) > 1 {
			fmt.Printf("Shard %s is underloaded (load: %d). Consider merging...\n", shards[i].ID, shards[i].Load)
		}
	}
	return shards
}
