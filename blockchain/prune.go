package blockchain

import (
	"crypto/sha256"
	"fmt"
)

// StateBlock represents a simplified state snapshot
type StateBlock struct {
	Index int
	Data  string
	Hash  string
}

// GenerateSnapshot creates a hashed snapshot of a state block
func GenerateSnapshot(index int, data string) StateBlock {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d-%s", index, data)))
	return StateBlock{
		Index: index,
		Data:  data,
		Hash:  fmt.Sprintf("%x", hash),
	}
}

// PruneStates simulates pruning older blocks and storing snapshot
func PruneStates(chain []StateBlock, keepLast int) []StateBlock {
	if len(chain) <= keepLast {
		return chain
	}
	fmt.Println("🔒 Archiving old state snapshots...")
	archived := chain[:len(chain)-keepLast]
	for _, snap := range archived {
		fmt.Printf("🗃 Archived Snapshot: Index=%d Hash=%s\n", snap.Index, snap.Hash)
	}
	return chain[len(chain)-keepLast:]
}
