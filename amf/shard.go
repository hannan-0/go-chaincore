package amf

import (
	"crypto/sha256"
	"encoding/hex"
)

type CrossShardCommitment struct {
	ID            string
	SourceData    []byte
	DestinationID string
	Hash          string
	Confirmed     bool
}

/*
	type CrossShardCommitment struct {
		ID            string // Keep this for a unique ID if needed
		SourceData    []byte
		DestinationID string
		Hash          string
		Commitment    string // Use this if you want to keep a commitment hash
		Payload       string // For syncing between shards, store payload too
		Confirmed     bool
	}
*/

type Shard struct {
	ID             string
	MerkleRoot     *MerkleNode
	Load           int
	Data           [][]byte
	Filter         *BloomFilter
	PendingCommits map[string]*CrossShardCommitment
	Accumulator    *HashAccumulator
	Nodes          []string
	//
	Transactions []string
}

func (s *Shard) ShouldSplit(threshold int) bool {
	return s.Load > threshold
}

func (s *Shard) ShouldMerge(threshold int) bool {
	return s.Load < threshold
}

func CreateCommitment(id, destination string, data []byte) *CrossShardCommitment {
	hash := sha256.Sum256(append([]byte(id+destination), data...))
	return &CrossShardCommitment{
		ID:            id,
		SourceData:    data,
		DestinationID: destination,
		Hash:          hex.EncodeToString(hash[:]),
		Confirmed:     false,
	}
}
