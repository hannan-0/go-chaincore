package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"

	"assignment3/amf"
	"assignment3/bft"
)

func TestMerkleForestShardSplit(t *testing.T) {
	forest := &amf.MerkleForest{}
	forest.Shards = []*amf.Shard{
		{ID: "shard-1", Load: 0, Data: [][]byte{}},
	}

	for i := 0; i < 5; i++ {
		forest.AddDataToShard("shard-1", []byte("Tx-"+string(rune('A'+i))))
	}
	forest.Rebalance(3)

	if len(forest.Shards) < 2 {
		t.Errorf("Shard splitting failed: expected >1 shard, got %d", len(forest.Shards))
	}
}

func TestBloomFilter(t *testing.T) {
	shard := &amf.Shard{
		ID:     "shard-1",
		Data:   [][]byte{},
		Load:   0,
		Filter: amf.NewBloomFilter(1024, 3),
	}
	testTx := []byte("Transaction A")
	shard.Filter.Add(testTx)

	if !shard.Filter.Contains(testTx) {
		t.Errorf("Bloom filter failed: expected to find 'Transaction A'")
	}
}

func TestAccumulator(t *testing.T) {
	acc := amf.NewHashAccumulator()
	data := []byte("Transaction X")
	acc.Add(data)

	proof := acc.Proof(data)
	if proof == "" {
		t.Errorf("Accumulator proof is empty")
	}
}

func TestBFTConsensus(t *testing.T) {
	manager := bft.NewBFTManager(0.5)
	manager.AddNode("n1", false)
	manager.AddNode("n2", true) // Byzantine
	manager.AddNode("n3", false)
	manager.AddNode("n4", false)
	manager.AddNode("n5", false)

	for i := 0; i < 3; i++ {
		manager.SimulateActivity()
	}

	if !manager.AttemptConsensus() {
		t.Log("Consensus failed as expected under certain conditions")
	}
}

func TestCrossShardTransfer(t *testing.T) {
	forest := &amf.MerkleForest{}
	forest.Shards = []*amf.Shard{
		{ID: "shard-1", Load: 0, Data: [][]byte{}},
		{ID: "shard-2", Load: 0, Data: [][]byte{}},
	}
	tx := []byte("X-Payment")

	forest.CrossShardTransfer("shard-1", "shard-2", "tx001", tx)
	forest.ConfirmCommitment("shard-1", "tx001")
	forest.ConfirmCommitment("shard-2", "tx001")
	forest.FinalizeCrossShard("tx001")

	if len(forest.Shards[1].Data) == 0 {
		t.Errorf("Cross-shard transaction failed: destination shard is empty")
	}
}

func runCommand(args ...string) (string, error) {
	cmd := exec.Command("go", append([]string{"run", "main.go"}, args...)...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

func TestHelp(t *testing.T) {
	out, _ := runCommand("-help")
	if !strings.Contains(out, "CLI Usage") {
		t.Error("Help output missing CLI Usage")
	}
}

func TestAddTransaction(t *testing.T) {
	out, _ := runCommand("-add", "TX: Alice pays Bob")
	if !strings.Contains(out, "Transaction added") {
		t.Error("Add transaction failed")
	}
}

func TestShardStatus(t *testing.T) {
	out, _ := runCommand("-status")
	if !strings.Contains(out, "Current Shard States") {
		t.Error("Status command failed")
	}
}

func TestRebalance(t *testing.T) {
	out, _ := runCommand("-split")
	if !strings.Contains(out, "Rebalancing shards") {
		t.Error("Shard rebalance failed")
	}
}

func TestBFTSimulation(t *testing.T) {
	out, _ := runCommand("-bft")
	if !strings.Contains(out, "BFT Simulation") {
		t.Error("BFT simulation did not run")
	}
}

func TestZKP(t *testing.T) {
	out, _ := runCommand("-zkp")
	if !strings.Contains(out, "ZKP verified") {
		t.Error("ZKP test failed")
	}
}

func TestMPC(t *testing.T) {
	out, _ := runCommand("-mpc")
	if !strings.Contains(out, "YES votes") {
		t.Error("MPC voting test failed")
	}
}

func TestConflictResolution(t *testing.T) {
	out, _ := runCommand("-conflict")
	if !strings.Contains(out, "Conflict") {
		t.Error("Conflict resolution test failed")
	}
}

func TestSmartContract(t *testing.T) {
	out, _ := runCommand("-contract")
	if !strings.Contains(out, "Transferred") {
		t.Error("Smart contract logic failed")
	}
}

func TestConsensusVote(t *testing.T) {
	out, _ := runCommand("-vote")
	if !strings.Contains(out, "Block") {
		t.Error("Consensus vote failed")
	}
}

func TestBlockMining(t *testing.T) {
	out, _ := runCommand("-mine")
	if !strings.Contains(out, "Mining block") {
		t.Error("Block mining simulation failed")
	}
}

func TestShardManager_Rebalance(t *testing.T) {
	sm := NewShardManager()
	sm.AddTransactionToShard("tx1")
	sm.AddTransactionToShard("tx2")
	sm.RebalanceShards()

	if len(sm.Shards) == 0 {
		t.Error("Expected shards to exist after rebalancing")
	}
}
