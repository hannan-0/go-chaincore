package amf

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type ShardManager struct {
	shards map[string]*Shard
	mu     sync.RWMutex
}

func NewShardManager() *ShardManager {
	return &ShardManager{
		shards: make(map[string]*Shard),
	}
}

func (sm *ShardManager) AddTransactionToShard(tx string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Assign to least-loaded shard
	var target *Shard
	for _, shard := range sm.shards {
		if target == nil || shard.Load < target.Load {
			target = shard
		}
	}
	if target == nil {
		// Create initial shard if none
		shardID := fmt.Sprintf("%x", sha256.Sum256([]byte(tx)))
		sm.shards[shardID] = &Shard{ID: shardID, Transactions: []string{tx}, Load: 1}
		return
	}

	target.Transactions = append(target.Transactions, tx)
	target.Load++
}

func (sm *ShardManager) RebalanceShards() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	for id, shard := range sm.shards {
		if shard.Load > 10 {
			newShardID := fmt.Sprintf("%x", sha256.Sum256([]byte(id+"-split")))
			newShard := &Shard{ID: newShardID}

			half := len(shard.Transactions) / 2
			newShard.Transactions = shard.Transactions[half:]
			newShard.Load = len(newShard.Transactions)

			shard.Transactions = shard.Transactions[:half]
			shard.Load = len(shard.Transactions)

			sm.shards[newShardID] = newShard
		}
	}
}
