package cap

import "math/rand"

type ConsistencyLevel string

const (
	Strong   ConsistencyLevel = "Strong"
	Eventual ConsistencyLevel = "Eventual"
	Causal   ConsistencyLevel = "Causal"
)

type CAPOrchestrator struct {
	PartitionLikelihood float64
	CurrentLevel        ConsistencyLevel
}

func NewCAPOrchestrator() *CAPOrchestrator {
	return &CAPOrchestrator{
		PartitionLikelihood: 0.0,
		CurrentLevel:        Strong,
	}
}

func (c *CAPOrchestrator) MonitorNetworkConditions() {
	c.PartitionLikelihood = rand.Float64()
	if c.PartitionLikelihood > 0.7 {
		c.CurrentLevel = Eventual
	} else if c.PartitionLikelihood > 0.4 {
		c.CurrentLevel = Causal
	} else {
		c.CurrentLevel = Strong
	}
}

func (c *CAPOrchestrator) GetConsistencyLevel() ConsistencyLevel {
	return c.CurrentLevel
}
