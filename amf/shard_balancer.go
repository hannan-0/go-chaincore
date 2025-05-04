package amf

import (
	"log"
	"sync"
)

type ShardBalancer struct {
	shards map[string]*Shard
	mu     sync.RWMutex
}

func NewShardBalancer() *ShardBalancer {
	return &ShardBalancer{
		shards: make(map[string]*Shard),
	}
}

func (sb *ShardBalancer) AddShard(id string, initialLoad int, nodes []string) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.shards[id] = &Shard{
		ID:    id,
		Load:  initialLoad,
		Nodes: nodes,
	}
}

func (sb *ShardBalancer) EvaluateAndRebalance() {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	for id, shard := range sb.shards {
		if shard.Load > 1000 {
			log.Printf("[ShardBalancer] Splitting overloaded shard: %s", id)
			sb.splitShard(shard)
		} else if shard.Load < 300 && len(sb.shards) > 1 {
			log.Printf("[ShardBalancer] Merging underutilized shard: %s", id)
			sb.mergeShard(id)
		}
	}
}

func (sb *ShardBalancer) splitShard(oldShard *Shard) {
	newID := oldShard.ID + "_split"
	newShard := &Shard{
		ID:    newID,
		Load:  oldShard.Load / 2,
		Nodes: []string{},
	}
	oldShard.Load /= 2
	sb.shards[newID] = newShard
}

func (sb *ShardBalancer) mergeShard(shardID string) {
	// Simple strategy: merge with a random neighbor
	for id, s := range sb.shards {
		if id != shardID {
			s.Load += sb.shards[shardID].Load
			delete(sb.shards, shardID)
			break
		}
	}
}
