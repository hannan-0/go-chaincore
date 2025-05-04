# Advanced Blockchain System in Go 🚀

A high-performance blockchain prototype developed in Go as part of **CS4049: Blockchain and Cryptocurrency**. It integrates cutting-edge cryptographic techniques, consensus algorithms, and state optimization methods.

---

## 🧠 Key Innovations

- Adaptive Merkle Forests (AMF) with dynamic sharding and load balancing
- CAP-aware consistency tuning using network telemetry
- Zero-Knowledge Proofs (ZKPs) and cryptographic accumulators
- Hybrid PoW + Delegated BFT consensus
- Byzantine fault tolerance with node scoring and multi-layer defense
- Compact blockchain state compression and pruning
- Homomorphic cross-shard synchronization

---

## 🌟 Features

- Add and process transactions across shards
- Rebalance overloaded shards dynamically
- Run and test CAP tuning module independently
- Deploy and execute simple smart contracts
- Simulate hybrid consensus (PoW + BFT)
- Run zero-knowledge proof validations
- Measure blockchain pruning efficiency

---

## 📁 Project Structure

```bash
go-chaincore/
├── amf/              # Adaptive Merkle Forest and rebalancing logic
├── blockchain/       # Core blockchain data structures and execution
├── consensus/        # Hybrid PoW + dBFT consensus system
├── bft/              # Node scoring and BFT thresholding
├── cap/              # CAP optimization and conflict resolution
├── zkp/              # Zero-knowledge proof and MPC modules
├── crypto/           # Cryptographic primitives (VRF, accumulators)
├── state/            # Blockchain state compression and archival
├── main.go           # Unified CLI for feature execution
├── main_test.go      # Integration test suite
└── README.md         # This file
```
### 🔄 How It Works (Simplified Flow)

User submits a transaction via CLI
AMF module routes it to the best-fit shard
CAP module checks for consistency under network conditions
ZKP verifies any proofs or privacy constraints
Consensus layer confirms and builds a new block
BFT layer ensures fault tolerance and sync
State is pruned and archived using compression strategies

## 🧪 Running the System

### Prerequisites
Go 1.21+
Git
Unix-based system (macOS/Linux/WSL recommended)

### Build & Run

```bash
git clone https://github.com/hannan-0/go-chaincore.git
cd go-chaincore
go run main.go
```

## CLI Usage Examples
| Command        | Description                                |
| -------------- | ------------------------------------------ |
| `-add tx1`     | Add a transaction to shard-1               |
| `-split`       | Trigger shard rebalancing                  |
| `-feature=cap` | Run only the CAP optimization module       |
| `-zkp`         | Run ZKP verification example               |
| `-contract`    | Deploy and execute smart contract          |
| `-mine`        | Mine a new block with mempool transactions |
| `-help`        | View available CLI options                 |

## 🧪 Testing

Run the following to execute the integration tests:
```bash
go test -v
```

This runs all integration tests in main_test.go, covering AMF, CAP, ZKP, BFT, state, and consensus logic.

## 📘 Documentation
For full architecture, theory, and benchmarks, see: docs/report.pdf

## 📈 Performance Overview
| Feature                 | Complexity |
| ----------------------- | ---------- |
| Merkle Proof Generation | `O(log N)` |
| Shard Discovery         | `O(log N)` |
| Cross-Shard Sync        | `O(k)`     |
| Conflict Detection      | `O(n)`     |

## 👨‍🔬 Author
Muhammad Hannan Nadeem
Course: CS4049 – Blockchain and Cryptocurrency
Institution: National University of computer and Emerging Sciences - FAST-NUCES




