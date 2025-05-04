/*
package main

import (
	"fmt"

	"assignment3/amf"
	"assignment3/bft"
)

func testBFT() {
	fmt.Println("\n--- BFT Simulation ---")
	manager := bft.NewBFTManager(0.5)

	manager.AddNode("node-1", false)
	manager.AddNode("node-2", false)
	manager.AddNode("node-3", true) // Byzantine
	manager.AddNode("node-4", false)
	manager.AddNode("node-5", false)

	// Simulate 5 rounds of network activity
	for i := 0; i < 5; i++ {
		fmt.Printf("Round %d:\n", i+1)
		manager.SimulateActivity()
		manager.AttemptConsensus()
		fmt.Println()
	}
}

func main() {
	fmt.Println("--- Adaptive Merkle Forest Demo ---")

	// Initialize a new Merkle Forest
	forest := &amf.MerkleForest{}

	// Create some initial shards
	forest.Shards = []*amf.Shard{
		{ID: "shard-1", Load: 0, Data: [][]byte{}},
		{ID: "shard-2", Load: 0, Data: [][]byte{}},
	}

	// Simulate adding data
	forest.AddDataToShard("shard-1", []byte("Transaction A"))
	forest.AddDataToShard("shard-1", []byte("Transaction B"))
	forest.AddDataToShard("shard-2", []byte("Transaction C"))
	forest.AddDataToShard("shard-2", []byte("Transaction D"))
	forest.AddDataToShard("shard-2", []byte("Transaction E"))

	// Trigger rebalancing based on threshold
	rebalancingThreshold := 3
	forest.Rebalance(rebalancingThreshold)

	// Print Merkle roots
	for _, shard := range forest.Shards {
		fmt.Printf("Shard ID: %s\n", shard.ID)
		if shard.MerkleRoot != nil {
			fmt.Printf("Merkle Root Hash: %s\n", shard.MerkleRoot.Hash)
		} else {
			fmt.Println("Merkle Root: (empty)")
		}
		fmt.Println("Load:", shard.Load)
		fmt.Println()
	}

	// Bloom Filter Proof Test (after rebalancing)
	fmt.Println("--- Bloom Filter Proof Test ---")
	testData := []byte("Transaction E")
	found := false
	for _, shard := range forest.Shards {
		if shard.Filter != nil && shard.Filter.Contains(testData) {
			fmt.Println("Transaction E is *probably* in", shard.ID)
			found = true
		}
	}
	if !found {
		fmt.Println("Transaction E is *definitely not* in any shard")
	}

	// Cross-Shard State Synchronization
	fmt.Println("\n--- Cross-Shard State Transfer ---")
	txData := []byte("Cross-shard Payment")
	forest.CrossShardTransfer("shard-1", "shard-2-1", "tx100", txData)
	forest.ConfirmCommitment("shard-1", "tx100")
	forest.ConfirmCommitment("shard-2-1", "tx100")
	forest.FinalizeCrossShard("tx100")

	// Accumulator Proof
	fmt.Println("\n--- Accumulator Proof ---")
	proof := forest.Shards[0].Accumulator.Proof([]byte("Transaction A"))
	fmt.Println("Accumulator Proof for Transaction A:", proof)
	fmt.Println("Accumulator Root Value:", forest.Shards[0].Accumulator.Value)

	// Run Byzantine Fault Tolerance Simulation
	testBFT()
}





*/

/*
//CLI     1
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"assignment3/amf"
	"assignment3/bft"
	"assignment3/blockchain"
	"assignment3/cap"
	"assignment3/consensus"
	"assignment3/zkp"
)

var forest = &amf.MerkleForest{}

func initShards() {
	forest.Shards = []*amf.Shard{
		{ID: "shard-1", Load: 0, Data: [][]byte{}},
		{ID: "shard-2", Load: 0, Data: [][]byte{}},
	}
}

func printShards() {
	fmt.Println("\n--- Current Shard States ---")
	for _, shard := range forest.Shards {
		fmt.Printf("Shard: %s | Load: %d | Merkle Root: ", shard.ID, shard.Load)
		if shard.MerkleRoot != nil {
			fmt.Println(shard.MerkleRoot.Hash)
		} else {
			fmt.Println("nil")
		}
	}
}

func runBFTSimulation() {
	fmt.Println("\n--- Running BFT Simulation ---")
	manager := bft.NewBFTManager(0.5)
	manager.AddNode("n1", false)
	manager.AddNode("n2", false)
	manager.AddNode("n3", true) // Byzantine
	manager.AddNode("n4", false)
	manager.AddNode("n5", false)

	for i := 0; i < 3; i++ {
		manager.SimulateActivity()
		manager.AttemptConsensus()
	}
}

// A sample contract function
func sampleContract(args ...string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("insufficient arguments")
	}
	// The contract logic here - transferring tokens
	return fmt.Sprintf("📦 Transferred %s tokens to %s", args[1], args[0]), nil
}

func main() {
	add := flag.String("add", "", "Add transaction to shard-1")
	split := flag.Bool("split", false, "Force rebalancing")
	cross := flag.Bool("cross", false, "Simulate cross-shard transfer")
	status := flag.Bool("status", false, "Print current shard state")
	bftSim := flag.Bool("bft", false, "Run BFT simulation")

	flag.Parse()
	initShards()

	if *add != "" {
		forest.AddDataToShard("shard-1", []byte(*add))
		fmt.Println("Transaction added to shard-1:", *add)
	}

	if *split {
		fmt.Println("Forcing rebalancing...")
		forest.Rebalance(3)
	}

	if *cross {
		fmt.Println("Simulating cross-shard transfer from shard-1 to shard-2...")
		tx := []byte("CLI Cross Payment")
		forest.CrossShardTransfer("shard-1", "shard-2", "cli-tx", tx)
		forest.ConfirmCommitment("shard-1", "cli-tx")
		forest.ConfirmCommitment("shard-2", "cli-tx")
		forest.FinalizeCrossShard("cli-tx")
	}

	if *status {
		printShards()
	}

	if *bftSim {
		runBFTSimulation()
	}

	if len(os.Args) == 1 {
		fmt.Println("Usage:\n  -add <tx>  | -split  | -cross  | -status  | -bft")
	}

	shards := []amf.MonitoredShard{
		{ID: "shard-1"}, {ID: "shard-2"},
	}

	go func() {
		ticker := time.NewTicker(15 * time.Second)
		for {
			<-ticker.C
			shards = amf.MonitorAndBalance(shards, 75)
		}
	}()

	blockchain.InitBlockchain()

	for i := 1; i <= 105; i++ {
		blockchain.AddBlock(fmt.Sprintf("Data Block %d", i))
	}

	//cap-conflict
	vc1 := cap.VectorClock{"NodeA": 2, "NodeB": 1}
	vc2 := cap.VectorClock{"NodeA": 2, "NodeB": 1}

	data1 := "short"
	data2 := "longer_data"

	if cap.DetectConflict(vc1, vc2) {
		resolved := cap.ResolveConflict(data1, data2)
		fmt.Println("⚠️ Conflict detected. Resolved to:", resolved)
	} else {
		fmt.Println("✅ No conflict detected.")
	}

	// 🔐 Zero-Knowledge Proof example
	secret := "node-private-key"
	proof := zkp.Prove(secret)
	if zkp.Verify(proof, secret) {
		fmt.Println("✅ ZKP verification successful.")
	} else {
		fmt.Println("❌ ZKP verification failed.")
	}

	// 👥 Multi-Party Computation example
	participants := []zkp.MPCParticipant{
		{"Validator1", 1},
		{"Validator2", 0},
		{"Validator3", 1},
	}
	total := zkp.ComputeSum(participants)
	fmt.Printf("🧮 Secure vote count (YES votes): %d\n", total)

	// Simulate a transaction submission (e.g., Alice sending 50.0 to Bob)
	tx := blockchain.NewTransaction("Alice", "Bob", 50.0, "sig-alice")

	// Add transaction to mempool
	err := blockchain.AddToMempool(tx)
	if err != nil {
		log.Fatal(err)
	}

	// Assuming a block mining or commit logic is here
	// (for the sake of this example, let's simulate block mining)
	fmt.Println("⛏️ Mining block...")

	// Once block is mined/committed, clear the mempool
	blockchain.ClearMempool()

	// Check mempool is empty after mining
	fmt.Printf("Mempool after mining: %v\n", blockchain.Mempool)

	committee := []consensus.CommitteeMember{
		{ID: "NodeA", Weight: 3},
		{ID: "NodeB", Weight: 2},
		{ID: "NodeC", Weight: 1},
	}

	approved := consensus.SimulateVoting(committee)
	if approved {
		fmt.Println("✅ Block approved by consensus!")
	} else {
		fmt.Println("❌ Block rejected by consensus!")
	}

	// Deploy the contract
	blockchain.DeployContract("0xabc123", "Alice", sampleContract)

	// Execute the contract with arguments: receiver ("Bob") and amount ("50")
	result, err := blockchain.ExecuteContract("0xabc123", "Bob", "50")
	if err != nil {
		fmt.Println("❌", err)
	} else {
		fmt.Println("✅", result)
	}

	txs := []blockchain.Transaction{
		{From: "Alice", To: "Bob", Amount: 10},
	}

	prevHash := consensus.GetLastHash()
	block := consensus.CreateBlock(txs, prevHash)
	consensus.AddBlock(block)

}
*/

// CLI for Assignment3 Blockchain System
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"assignment3/amf"
	"assignment3/bft"
	"assignment3/blockchain"
	"assignment3/cap"
	"assignment3/consensus"
	"assignment3/state"
	"assignment3/zkp"
)

var forest = &amf.MerkleForest{}

func initShards() {
	forest.Shards = []*amf.Shard{
		{ID: "shard-1", Load: 0, Data: [][]byte{}},
		{ID: "shard-2", Load: 0, Data: [][]byte{}},
	}
}

func printShardStates() {
	fmt.Println("\n--- Current Shard States ---")
	for _, shard := range forest.Shards {
		fmt.Printf("Shard: %s | Load: %d | Merkle Root: ", shard.ID, shard.Load)
		if shard.MerkleRoot != nil {
			fmt.Println(shard.MerkleRoot.Hash)
		} else {
			fmt.Println("nil")
		}
	}
}

func simulateBFT() {
	fmt.Println("\n--- BFT Simulation ---")
	manager := bft.NewBFTManager(0.5)
	manager.AddNode("n1", false)
	manager.AddNode("n2", false)
	manager.AddNode("n3", true)
	manager.AddNode("n4", false)
	manager.AddNode("n5", false)

	for i := 0; i < 3; i++ {
		manager.SimulateActivity()
		manager.AttemptConsensus()
	}
}

func sampleContract(args ...string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("insufficient arguments")
	}
	return fmt.Sprintf("📦 Transferred %s tokens to %s", args[1], args[0]), nil
}

func monitorShards() {
	shards := []amf.MonitoredShard{{ID: "shard-1"}, {ID: "shard-2"}}
	go func() {
		ticker := time.NewTicker(15 * time.Second)
		for {
			<-ticker.C
			shards = amf.MonitorAndBalance(shards, 75)
		}
	}()
}

func handleZKP() {
	fmt.Println("\n🔐 Zero-Knowledge Proof")
	secret := "node-private-key"
	proof := zkp.Prove(secret)
	if zkp.Verify(proof, secret) {
		fmt.Println("✅ ZKP verified")
	} else {
		fmt.Println("❌ ZKP verification failed")
	}
}

func handleMPC() {
	fmt.Println("\n👥 Multi-Party Computation Voting")
	participants := []zkp.MPCParticipant{
		{"Validator1", 1},
		{"Validator2", 0},
		{"Validator3", 1},
	}
	total := zkp.ComputeSum(participants)
	fmt.Printf("🧮 YES votes: %d\n", total)
}

func handleConflict() {
	fmt.Println("\n⚔️ Conflict Resolution via Vector Clocks")
	vc1 := cap.VectorClock{"NodeA": 2, "NodeB": 1}
	vc2 := cap.VectorClock{"NodeA": 2, "NodeB": 1}
	if cap.DetectConflict(vc1, vc2) {
		res := cap.ResolveConflict("short", "longer_data")
		fmt.Println("⚠️ Conflict resolved to:", res)
	} else {
		fmt.Println("✅ No conflict")
	}
}

func handleContract() {
	fmt.Println("\n📜 Smart Contract Execution")
	blockchain.DeployContract("0xabc123", "Alice", sampleContract)
	result, err := blockchain.ExecuteContract("0xabc123", "Bob", "50")
	if err != nil {
		fmt.Println("❌", err)
	} else {
		fmt.Println("✅", result)
	}
}

func handleConsensusVote() {
	fmt.Println("\n🗳️ Committee Voting")
	committee := []consensus.CommitteeMember{
		{"NodeA", 3, true}, {"NodeB", 2, false}, {"NodeC", 1, true},
	}
	if consensus.SimulateVoting(committee) {
		fmt.Println("✅ Block approved!")
	} else {
		fmt.Println("❌ Block rejected!")
	}
}

func handleBlockchainOps() {
	fmt.Println("\n⛏️ Blockchain Initialization & Mining")
	blockchain.InitBlockchain()
	for i := 1; i <= 5; i++ {
		blockchain.AddBlock(fmt.Sprintf("Block Data %d", i))
	}
	tx := blockchain.NewTransaction("Alice", "Bob", 50.0, "sig-alice")
	if err := blockchain.AddToMempool(tx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("⛏️ Mining block...")
	blockchain.ClearMempool()
	fmt.Printf("Mempool: %v\n", blockchain.Mempool)
}

func runAMF() {
	fmt.Println("[AMF] Initializing Adaptive Merkle Forest and Rebalancer...")
	shardManager := amf.NewShardManager()
	shardManager.AddTransactionToShard("tx1")
	shardManager.AddTransactionToShard("tx2")
	shardManager.AddTransactionToShard("tx3")
	shardManager.RebalanceShards()
	fmt.Println("[AMF] Rebalancing complete.")
}

func runCAP() {
	fmt.Println("[CAP] Starting CAP optimizer...")
	optimizer := cap.NewCAPOptimizer()
	optimizer.UpdateTelemetry(0.6, 0.4)
	fmt.Println("[CAP] Current Consistency:", optimizer.GetConsistencyLevel())
	fmt.Println("[CAP] Retry Timeout:", optimizer.GetRetryTimeout())
}

func runConsensus() {
	fmt.Println("[Consensus] Initializing hybrid consensus protocol...")
	engine := consensus.NewConsensusEngine()
	engine.AdjustReputation("NodeA", 0.9)
	engine.AdjustReputation("NodeB", 0.6)
	rand := engine.InjectPoWNoise()
	fmt.Println("[Consensus] PoW-based randomness:", rand)
	selected := engine.VoteOnBlock([]string{"NodeA", "NodeB"})
	fmt.Println("[Consensus] Block approved by:", selected)
	fmt.Println("[Consensus] Finalizing consensus:", engine.FinalizeConsensus())
}

func runState() {
	fmt.Println("[State] Archiving blockchain state...")
	archive := state.NewStateArchive()
	node := state.StateNode{Key: "balance:Alice", Value: "100"}
	archive.CompressAndStore(node)
	hash, exists := archive.RetrieveCompressedHash("balance:Alice")
	if exists {
		fmt.Println("[State] Compressed Hash:", hash)
	} else {
		fmt.Println("[State] No hash found.")
	}
}

func runAll() {
	runAMF()
	runCAP()
	runConsensus()
	runState()
	fmt.Println("✅ All modules executed successfully.")
	time.Sleep(1 * time.Second)
}

func main() {
	add := flag.String("add", "", "Add transaction to shard-1")
	split := flag.Bool("split", false, "Rebalance shards")
	cross := flag.Bool("cross", false, "Simulate cross-shard transfer")
	status := flag.Bool("status", false, "Show shard states")
	bftSim := flag.Bool("bft", false, "Run BFT simulation")
	zproof := flag.Bool("zkp", false, "Run ZKP example")
	mpc := flag.Bool("mpc", false, "Run MPC voting example")
	conflict := flag.Bool("conflict", false, "Test conflict resolution")
	contract := flag.Bool("contract", false, "Deploy & execute smart contract")
	consVote := flag.Bool("vote", false, "Run consensus vote via MPC")
	block := flag.Bool("mine", false, "Simulate block mining and mempool")
	help := flag.Bool("help", false, "Display help menu")
	feature := flag.String("feature", "all", "Feature to run: amf | cap | consensus | state | all")

	flag.Parse()
	initShards()
	monitorShards()

	didRunSomething := false

	switch {
	case *help || len(os.Args) == 1:
		fmt.Println("\n📘 CLI Usage:")
		fmt.Println("  -add <tx>       Add transaction to shard-1")
		fmt.Println("  -split          Force shard rebalancing")
		fmt.Println("  -cross          Simulate cross-shard transfer")
		fmt.Println("  -status         Print shard states")
		fmt.Println("  -bft            Simulate Byzantine Fault Tolerance")
		fmt.Println("  -zkp            Run Zero-Knowledge Proof verification")
		fmt.Println("  -mpc            Run Multi-Party Computation voting")
		fmt.Println("  -conflict       Test vector clock conflict resolution")
		fmt.Println("  -contract       Deploy and execute a smart contract")
		fmt.Println("  -vote           Simulate committee voting via MPC")
		fmt.Println("  -mine           Submit tx, mine block, clear mempool")
		fmt.Println("  -help           Show this menu")
	case *add != "":
		didRunSomething = true
		forest.AddDataToShard("shard-1", []byte(*add))
		fmt.Println("✅ Transaction added to shard-1:", *add)
	case *split:
		didRunSomething = true
		fmt.Println("🔁 Rebalancing shards...")
		forest.Rebalance(3)
	case *cross:
		didRunSomething = true
		fmt.Println("🔀 Cross-shard transfer shard-1 ➝ shard-2")
		tx := []byte("CLI Cross Payment")
		forest.CrossShardTransfer("shard-1", "shard-2", "cli-tx", tx)
		forest.ConfirmCommitment("shard-1", "cli-tx")
		forest.ConfirmCommitment("shard-2", "cli-tx")
		forest.FinalizeCrossShard("cli-tx")
	case *status:
		didRunSomething = true
		printShardStates()
	case *bftSim:
		didRunSomething = true
		simulateBFT()
	case *zproof:
		didRunSomething = true
		handleZKP()
	case *mpc:
		didRunSomething = true
		handleMPC()
	case *conflict:
		didRunSomething = true
		handleConflict()
	case *contract:
		didRunSomething = true
		handleContract()
	case *consVote:
		didRunSomething = true
		handleConsensusVote()
	case *block:
		didRunSomething = true
		handleBlockchainOps()

	default:
		fmt.Println("❗ Invalid flag. Use -help to view options.")
	}

	if !didRunSomething {
		switch *feature {
		case "amf":
			runAMF()
		case "cap":
			runCAP()
		case "consensus":
			runConsensus()
		case "state":
			runState()
		default:
			runAll()
		}
	}

	//
	//
	// --- Adaptive Merkle Forest + Dynamic Rebalancing ---
	fmt.Println("[AMF] Initializing Adaptive Merkle Forest and Rebalancer...")
	shardManager := amf.NewShardManager()
	shardManager.AddTransactionToShard("tx1")
	shardManager.AddTransactionToShard("tx2")
	shardManager.AddTransactionToShard("tx3")
	shardManager.RebalanceShards()
	fmt.Println("[AMF] Rebalancing complete.")

	// --- CAP Theorem Adaptive Optimizer ---
	fmt.Println("[CAP] Starting CAP optimizer...")
	optimizer := cap.NewCAPOptimizer()
	optimizer.UpdateTelemetry(0.6, 0.4)
	fmt.Println("[CAP] Current Consistency:", optimizer.GetConsistencyLevel())
	fmt.Println("[CAP] Retry Timeout:", optimizer.GetRetryTimeout())

	// --- Hybrid Consensus Engine ---
	fmt.Println("[Consensus] Initializing hybrid consensus protocol...")
	engine := consensus.NewConsensusEngine()
	engine.AdjustReputation("NodeA", 0.9)
	engine.AdjustReputation("NodeB", 0.6)
	rand := engine.InjectPoWNoise()
	fmt.Println("[Consensus] PoW-based randomness:", rand)
	selected := engine.VoteOnBlock([]string{"NodeA", "NodeB"})
	fmt.Println("[Consensus] Block approved by:", selected)
	fmt.Println("[Consensus] Finalizing consensus:", engine.FinalizeConsensus())

	// --- State Compression and Archival ---
	fmt.Println("[State] Archiving blockchain state...")
	archive := state.NewStateArchive()
	node := state.StateNode{Key: "balance:Alice", Value: "100"}
	archive.CompressAndStore(node)
	hash, exists := archive.RetrieveCompressedHash("balance:Alice")
	if exists {
		fmt.Println("[State] Compressed Hash:", hash)
	} else {
		fmt.Println("[State] No hash found.")
	}

	fmt.Println("✅ All modules executed successfully.")
	time.Sleep(1 * time.Second)

}
