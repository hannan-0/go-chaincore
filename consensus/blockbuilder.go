package consensus

import (
	"assignment3/blockchain"
	"assignment3/zkp"
	"fmt"
	"time"
)

// Block represents a simplified blockchain block
type Block struct {
	Index        int
	Timestamp    string
	Transactions []blockchain.Transaction
	PrevHash     string
	Hash         string
	Nonce        int
	ZKPProof     string
}

// Blockchain is the ledger (in-memory slice)
var Blockchain []Block

// CreateBlock constructs a new block from a transaction batch
func CreateBlock(txs []blockchain.Transaction, prevHash string) Block {
	index := len(Blockchain)
	timestamp := time.Now().Format(time.RFC3339)
	dataHash := blockchain.HashTransactions(txs)
	proof := zkp.Prove(dataHash)

	newBlock := Block{
		Index:        index,
		Timestamp:    timestamp,
		Transactions: txs,
		PrevHash:     prevHash,
		Hash:         blockchain.ComputeBlockHash(index, timestamp, txs, prevHash),
		ZKPProof:     proof,
	}

	fmt.Printf("🧱 Block #%d created (ZKP: %s...)\n", index, proof[:10])
	return newBlock
}

// AddBlock validates and appends a new block to the chain
func AddBlock(block Block) bool {
	if len(Blockchain) > 0 && block.PrevHash != Blockchain[len(Blockchain)-1].Hash {
		fmt.Println("❌ Invalid previous hash")
		return false
	}

	if !zkp.Verify(block.ZKPProof, blockchain.HashTransactions(block.Transactions)) {
		fmt.Println("❌ Invalid ZKP proof")
		return false
	}

	Blockchain = append(Blockchain, block)
	fmt.Printf("✅ Block #%d appended to the chain\n", block.Index)
	return true
}

// GetLastHash returns the hash of the latest block
func GetLastHash() string {
	if len(Blockchain) == 0 {
		return ""
	}
	return Blockchain[len(Blockchain)-1].Hash
}

// GetBlockchain returns the current chain
func GetBlockchain() []Block {
	return Blockchain
}
