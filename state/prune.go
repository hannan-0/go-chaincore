package state

import (
	"crypto/sha256"
	"fmt"
	"log"
	"sync"
)

type StateSnapshot struct {
	BlockHeight int
	StateRoot   []byte
	Archived    bool
}

type Pruner struct {
	stateHistory []*StateSnapshot
	lock         sync.Mutex
}

func NewPruner() *Pruner {
	return &Pruner{
		stateHistory: []*StateSnapshot{},
	}
}

func (p *Pruner) AddSnapshot(height int, root []byte) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.stateHistory = append(p.stateHistory, &StateSnapshot{BlockHeight: height, StateRoot: root, Archived: false})
}

func (p *Pruner) Prune(threshold int) {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, snapshot := range p.stateHistory {
		if snapshot.BlockHeight < threshold && !snapshot.Archived {
			log.Printf("[Pruner] Archiving snapshot at height %d", snapshot.BlockHeight)
			snapshot.Archived = true
			// Could write to disk, IPFS, or other archival service
		}
	}
}

func (p *Pruner) GetLatestUnarchived() *StateSnapshot {
	p.lock.Lock()
	defer p.lock.Unlock()
	for i := len(p.stateHistory) - 1; i >= 0; i-- {
		if !p.stateHistory[i].Archived {
			return p.stateHistory[i]
		}
	}
	return nil
}

type StateNode struct {
	Key   string
	Value string
	Hash  string
}

type StateArchive struct {
	Archived map[string]string
}

func NewStateArchive() *StateArchive {
	return &StateArchive{
		Archived: make(map[string]string),
	}
}

func (sa *StateArchive) CompressAndStore(node StateNode) {
	data := node.Key + node.Value
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
	sa.Archived[node.Key] = hash
}

func (sa *StateArchive) RetrieveCompressedHash(key string) (string, bool) {
	hash, exists := sa.Archived[key]
	return hash, exists
}
