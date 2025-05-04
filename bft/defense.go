package bft

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	ID          string
	Reputation  float64
	IsByzantine bool
	LastActive  time.Time
}

type BFTManager struct {
	Nodes                 []*Node
	MinConsensusThreshold float64
}

// NewBFTManager initializes the BFT system with default threshold
func NewBFTManager(threshold float64) *BFTManager {
	return &BFTManager{
		MinConsensusThreshold: threshold,
	}
}

// AddNode to the network
func (m *BFTManager) AddNode(id string, byzantine bool) {
	node := &Node{
		ID:          id,
		Reputation:  1.0, // start neutral
		IsByzantine: byzantine,
		LastActive:  time.Now(),
	}
	m.Nodes = append(m.Nodes, node)
}

// UpdateReputation based on behavior
func (m *BFTManager) UpdateReputation(id string, success bool) {
	for _, node := range m.Nodes {
		if node.ID == id {
			if success {
				node.Reputation += 0.1
			} else {
				node.Reputation -= 0.2
			}
			if node.Reputation < 0 {
				node.Reputation = 0
			}
			return
		}
	}
}

// GetTrustedNodes returns nodes above the threshold
func (m *BFTManager) GetTrustedNodes() []*Node {
	var trusted []*Node
	for _, node := range m.Nodes {
		if node.Reputation >= m.MinConsensusThreshold {
			trusted = append(trusted, node)
		}
	}
	return trusted
}

// AttemptConsensus simulates reaching consensus
func (m *BFTManager) AttemptConsensus() bool {
	trusted := m.GetTrustedNodes()
	trustRatio := float64(len(trusted)) / float64(len(m.Nodes))
	fmt.Printf("Trusted node ratio: %.2f\n", trustRatio)

	if trustRatio >= 0.67 {
		fmt.Println("Consensus reached ✅")
		return true
	}
	fmt.Println("Consensus failed ❌")
	return false
}

// SimulateActivity randomly updates node reputations
func (m *BFTManager) SimulateActivity() {
	for _, node := range m.Nodes {
		if node.IsByzantine {
			m.UpdateReputation(node.ID, false)
		} else {
			success := rand.Float64() < 0.95
			m.UpdateReputation(node.ID, success)
		}
	}
}
