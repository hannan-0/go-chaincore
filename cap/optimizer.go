package cap

import (
	"log"
	"math/rand"
	"time"
)

type CAPOptimizer struct {
	partitionLikelihood float64
	currentLevel        ConsistencyLevel
}

func NewCAPOptimizer() *CAPOptimizer {
	return &CAPOptimizer{
		partitionLikelihood: 0.0,
		currentLevel:        Strong,
	}
}

func (c *CAPOptimizer) EvaluateNetwork() {
	// Simulated telemetry
	c.partitionLikelihood = rand.Float64()

	if c.partitionLikelihood > 0.7 {
		c.currentLevel = Eventual
	} else if c.partitionLikelihood > 0.4 {
		c.currentLevel = Causal
	} else {
		c.currentLevel = Strong
	}

	log.Printf("[CAPOptimizer] Partition probability: %.2f, Set consistency: %v", c.partitionLikelihood, c.currentLevel)
}

func (c *CAPOptimizer) GetConsistencyLevel() ConsistencyLevel {
	return c.currentLevel
}

func (c *CAPOptimizer) UpdateTelemetry(networkDelay float64, packetLoss float64) {
	c.partitionLikelihood = (networkDelay + packetLoss) / 2
	if c.partitionLikelihood > 0.5 {
		c.currentLevel = Eventual
	} else {
		c.currentLevel = Strong
	}
}

func (c *CAPOptimizer) GetRetryTimeout() time.Duration {
	if c.currentLevel == Strong {
		return 3 * time.Second
	}
	return 1 * time.Second
}

func (c *CAPOptimizer) PredictPartition() bool {
	return rand.Float64() < c.partitionLikelihood
}
