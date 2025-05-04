package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp time.Time
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func CalculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s%d", block.Index, block.Timestamp.String(), block.Data, block.PrevHash, block.Nonce)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func GenerateBlock(oldBlock Block, data string, nonce int) Block {
	newBlock := Block{
		Index:     oldBlock.Index + 1,
		Timestamp: time.Now(),
		Data:      data,
		PrevHash:  oldBlock.Hash,
		Nonce:     nonce,
	}
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock
}
