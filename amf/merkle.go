package amf

import (
	"crypto/sha256"
	"encoding/hex"
)

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  string
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := &MerkleNode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Hash = hex.EncodeToString(hash[:])
	} else {
		combined := []byte(left.Hash + right.Hash)
		hash := sha256.Sum256(combined)
		node.Hash = hex.EncodeToString(hash[:])
	}
	node.Left = left
	node.Right = right
	return node
}
