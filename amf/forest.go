package amf

import (
	"fmt"
)

type MerkleForest struct {
	Shards []*Shard
}

func (mf *MerkleForest) AddDataToShard(shardID string, data []byte) {
	for _, shard := range mf.Shards {
		if shard.ID == shardID {
			shard.Data = append(shard.Data, data)

			if shard.Filter == nil {
				shard.Filter = NewBloomFilter(1024, 3)
			}
			shard.Filter.Add(data)

			if shard.Accumulator == nil {
				shard.Accumulator = NewHashAccumulator()
			}
			shard.Accumulator.Add(data)

			shard.Load++
			mf.recalculateMerkle(shard)
			return
		}
	}
}

func (mf *MerkleForest) recalculateMerkle(shard *Shard) {
	var nodes []*MerkleNode
	for _, d := range shard.Data {
		nodes = append(nodes, NewMerkleNode(nil, nil, d))
	}
	for len(nodes) > 1 {
		var level []*MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			if i+1 == len(nodes) {
				level = append(level, nodes[i])
			} else {
				level = append(level, NewMerkleNode(nodes[i], nodes[i+1], nil))
			}
		}
		nodes = level
	}
	if len(nodes) > 0 {
		shard.MerkleRoot = nodes[0]
	}
}

func (mf *MerkleForest) Rebalance(threshold int) {
	var newShards []*Shard
	shardCount := len(mf.Shards)

	for _, shard := range mf.Shards {
		if shard.ShouldSplit(threshold) && len(shard.Data) > 1 {
			fmt.Println("Splitting shard:", shard.ID)

			// Split data
			mid := len(shard.Data) / 2
			leftData := shard.Data[:mid]
			rightData := shard.Data[mid:]

			leftShard := &Shard{
				ID:   shard.ID + "-1",
				Data: leftData,
				Load: len(leftData),
			}
			// 🔸 Initialize and populate Bloom filter for left shard
			leftShard.Filter = NewBloomFilter(1024, 3)
			for _, d := range leftData {
				leftShard.Filter.Add(d)
			}
			// 🔸 Initialize and populate accumulator for left shard
			leftShard.Accumulator = NewHashAccumulator()
			for _, d := range leftData {
				leftShard.Accumulator.Add(d)
			}
			mf.recalculateMerkle(leftShard)

			rightShard := &Shard{
				ID:   shard.ID + "-2",
				Data: rightData,
				Load: len(rightData),
			}
			// 🔸 Initialize and populate Bloom filter for right shard
			rightShard.Filter = NewBloomFilter(1024, 3)
			for _, d := range rightData {
				rightShard.Filter.Add(d)
			}
			// 🔸 Initialize and populate accumulator for right shard
			rightShard.Accumulator = NewHashAccumulator()
			for _, d := range rightData {
				rightShard.Accumulator.Add(d)
			}
			mf.recalculateMerkle(rightShard)

			newShards = append(newShards, leftShard, rightShard)
		} else if shard.ShouldMerge(threshold) && shardCount > 1 {
			fmt.Println("Merge logic not implemented yet for:", shard.ID)
			newShards = append(newShards, shard)
		} else {
			newShards = append(newShards, shard)
		}
	}

	mf.Shards = newShards
}

func (mf *MerkleForest) CrossShardTransfer(fromID, toID, txID string, data []byte) {
	var fromShard, toShard *Shard
	for _, shard := range mf.Shards {
		if shard.ID == fromID {
			fromShard = shard
		}
		if shard.ID == toID {
			toShard = shard
		}
	}

	if fromShard == nil || toShard == nil {
		fmt.Println("Invalid shard IDs")
		return
	}

	commit := CreateCommitment(txID, toID, data)

	if fromShard.PendingCommits == nil {
		fromShard.PendingCommits = make(map[string]*CrossShardCommitment)
	}
	if toShard.PendingCommits == nil {
		toShard.PendingCommits = make(map[string]*CrossShardCommitment)
	}

	fromShard.PendingCommits[txID] = commit
	toShard.PendingCommits[txID] = commit
	fmt.Println("Commitment created:", commit.Hash)
}

func (mf *MerkleForest) ConfirmCommitment(shardID, txID string) {
	for _, shard := range mf.Shards {
		if shard.ID == shardID {
			if commit, ok := shard.PendingCommits[txID]; ok {
				commit.Confirmed = true
				fmt.Printf("Shard %s confirmed transaction %s\n", shardID, txID)
			}
		}
	}
}

func (mf *MerkleForest) FinalizeCrossShard(txID string) {
	var source, dest *Shard
	for _, shard := range mf.Shards {
		if commit, ok := shard.PendingCommits[txID]; ok {
			if shard.ID == commit.DestinationID {
				dest = shard
			} else {
				source = shard
			}
		}
	}

	if source != nil && dest != nil &&
		source.PendingCommits[txID].Confirmed &&
		dest.PendingCommits[txID].Confirmed {

		dest.Data = append(dest.Data, source.PendingCommits[txID].SourceData)
		dest.Load++
		mf.recalculateMerkle(dest)
		fmt.Println("Cross-shard transaction finalized.")
		delete(source.PendingCommits, txID)
		delete(dest.PendingCommits, txID)
	} else {
		fmt.Println("Cross-shard not ready for finalization.")
	}
}
