package blockchain

import (
	"fmt"
	"sync"
	"time"
)

// Global blockchain state
var (
	Blockchain     []Block
	stateSnapshots []StateBlock
	mu             sync.Mutex
	maxBlocks      = 100 // Keep last 100 blocks
)

// InitBlockchain creates the genesis block
func InitBlockchain() {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "", // Will be calculated
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)
	stateSnapshots = append(stateSnapshots, GenerateSnapshot(0, genesisBlock.Data))
}

// AddBlock creates a new block and adds it to the blockchain
func AddBlock(data string) {
	mu.Lock()
	defer mu.Unlock()

	prevBlock := Blockchain[len(Blockchain)-1]
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = CalculateHash(newBlock)
	Blockchain = append(Blockchain, newBlock)

	// Save state snapshot and prune old states
	stateSnapshots = append(stateSnapshots, GenerateSnapshot(newBlock.Index, data))
	stateSnapshots = PruneStates(stateSnapshots, maxBlocks)

	fmt.Printf("✅ Block %d added. Blockchain length: %d\n", newBlock.Index, len(Blockchain))
}
