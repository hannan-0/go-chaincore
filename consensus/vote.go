package consensus

import (
	"fmt"
	"math/rand"
	"time"
)

// CommitteeMember represents a node participating in consensus
type CommitteeMember struct {
	ID      string
	Weight  int // Voting power
	Approve bool
}

// SimulateVoting performs a secure-like multiparty vote on a block proposal
func SimulateVoting(members []CommitteeMember) bool {
	rand.Seed(time.Now().UnixNano())
	totalWeight := 0
	approveWeight := 0

	fmt.Println("🗳 Committee Voting Started")

	for _, member := range members {
		// Simulate voting behavior (random or policy-based)
		member.Approve = rand.Intn(2) == 1

		status := "❌ NO"
		if member.Approve {
			approveWeight += member.Weight
			status = "✅ YES"
		}
		totalWeight += member.Weight

		fmt.Printf("Member %s (%d votes): %s\n", member.ID, member.Weight, status)
	}

	threshold := int(0.66 * float64(totalWeight)) // 66% majority required
	fmt.Printf("🔢 Approve Weight: %d / %d (threshold: %d)\n", approveWeight, totalWeight, threshold)

	return approveWeight >= threshold
}
