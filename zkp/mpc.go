package zkp

import "fmt"

// MPCParticipant simulates a node in an MPC protocol
type MPCParticipant struct {
	ID    string
	Value int
}

// ComputeSum simulates secure summation using MPC
func ComputeSum(participants []MPCParticipant) int {
	sum := 0
	for _, p := range participants {
		sum += p.Value
		fmt.Printf("Participant %s contributes %d\n", p.ID, p.Value)
	}
	return sum
}
