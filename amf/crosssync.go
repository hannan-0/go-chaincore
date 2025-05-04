package amf

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// CrossShardCommitment represents a cryptographic commitment for syncing between shards
type CrossShardSyncCommitment struct {
	ShardIDFrom string
	ShardIDTo   string
	Payload     string
	Commitment  string
}

// GenerateCommitment creates a hash-based commitment for cross-shard operation
func GenerateCommitment(from, to, payload string) CrossShardSyncCommitment {
	data := from + to + payload
	hash := sha256.Sum256([]byte(data))
	commit := hex.EncodeToString(hash[:])
	return CrossShardSyncCommitment{
		ShardIDFrom: from,
		ShardIDTo:   to,
		Payload:     payload,
		Commitment:  commit,
	}
}

// VerifyCommitment ensures that a given commitment is valid for the payload
func VerifyCommitment(c CrossShardSyncCommitment) bool {
	expected := GenerateCommitment(c.ShardIDFrom, c.ShardIDTo, c.Payload)
	return c.Commitment == expected.Commitment
}

// SyncState simulates partial state transfer between shards with commitment
func SyncState(fromShard, toShard string, payload string) error {
	commit := GenerateCommitment(fromShard, toShard, payload)

	fmt.Printf("[SYNC] From: %s → To: %s\n", fromShard, toShard)
	fmt.Printf("[PAYLOAD] %s\n", payload)
	fmt.Printf("[COMMITMENT] %s\n", commit.Commitment)

	if !VerifyCommitment(commit) {
		return fmt.Errorf("commitment verification failed")
	}
	return nil
}
