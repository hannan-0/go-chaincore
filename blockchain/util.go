package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// HashTransactions computes a single hash for a list of transactions
func HashTransactions(txs []Transaction) string {
	txBytes, err := json.Marshal(txs)
	if err != nil {
		panic(fmt.Sprintf("Error hashing transactions: %v", err))
	}
	hash := sha256.Sum256(txBytes)
	return hex.EncodeToString(hash[:])
}

// ComputeBlockHash generates the hash of a block's metadata
func ComputeBlockHash(index int, timestamp string, txs []Transaction, prevHash string) string {
	blockData := fmt.Sprintf("%d-%s-%s-%s", index, timestamp, HashTransactions(txs), prevHash)
	hash := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hash[:])
}
