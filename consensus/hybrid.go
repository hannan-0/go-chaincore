package consensus

import (
	"assignment3/amf"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"
)

type ConsensusNode struct {
	ID       string
	Stake    int
	IsLeader bool
}

func HandleCrossShardTx(from, to, payload string) {
	err := amf.SyncState(from, to, payload)
	if err != nil {
		log.Printf("Cross-shard sync failed: %v\n", err)
	} else {
		log.Println("Cross-shard sync successful.")
	}
}

func ProofOfWork(seed string, difficulty int) string {
	var hash [32]byte
	nonce := 0
	for {
		input := fmt.Sprintf("%s%d", seed, nonce)
		hash = sha256.Sum256([]byte(input))
		if hex.EncodeToString(hash[:])[:difficulty] == string(make([]byte, difficulty)) {
			break
		}
		nonce++
	}
	return hex.EncodeToString(hash[:])
}

func RandomizeByStake(nodes []ConsensusNode) ConsensusNode {
	totalStake := 0
	for _, node := range nodes {
		totalStake += node.Stake
	}
	rand.Seed(time.Now().UnixNano())
	threshold := rand.Intn(totalStake)
	running := 0
	for _, node := range nodes {
		running += node.Stake
		if running > threshold {
			return node
		}
	}
	return nodes[0]
}

type dBFTNode struct {
	ID         string
	Reputation int
	Voted      bool
}

func RunDelegatedBFT(nodes []dBFTNode, block *Block) bool {
	voteCount := 0
	for i := range nodes {
		if rand.Intn(100) < nodes[i].Reputation {
			nodes[i].Voted = true
			voteCount++
		}
	}
	log.Printf("[dBFT] Votes received: %d/%d", voteCount, len(nodes))
	return voteCount > len(nodes)/2
}

func HybridConsensus(block *Block, validators []dBFTNode) bool {
	ProofOfWork(fmt.Sprintf("%v", block), 2)
	return RunDelegatedBFT(validators, block)
}

/*
package consensus

import (

	"crypto/rand"
	"math/big"
	"time"

)
*/
type ConsensusEngine struct {
	nodeReputation map[string]float64
}

func NewConsensusEngine() *ConsensusEngine {
	return &ConsensusEngine{
		nodeReputation: make(map[string]float64),
	}
}

// Simulate PoW-based randomness
func (ce *ConsensusEngine) InjectPoWNoise() int64 {
	n, err := crand.Int(crand.Reader, big.NewInt(100000))
	if err != nil {
		log.Println("crypto/rand.Int failed:", err)
		return 0
	}
	return n.Int64()
}

// Reputation-Based Voting (like dBFT)
func (ce *ConsensusEngine) VoteOnBlock(candidates []string) string {
	var best string
	var bestScore float64 = -1.0

	for _, c := range candidates {
		score := ce.nodeReputation[c]
		if score > bestScore {
			best = c
			bestScore = score
		}
	}
	return best
}

func (ce *ConsensusEngine) AdjustReputation(nodeID string, delta float64) {
	ce.nodeReputation[nodeID] += delta
}

func (ce *ConsensusEngine) FinalizeConsensus() bool {
	time.Sleep(500 * time.Millisecond)
	return true
}
