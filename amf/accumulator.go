package amf

import (
	"crypto/sha256"
	"encoding/hex"
)

type HashAccumulator struct {
	Value string
	Items [][]byte
}

// NewHashAccumulator initializes an accumulator
func NewHashAccumulator() *HashAccumulator {
	return &HashAccumulator{
		Value: "",
		Items: [][]byte{},
	}
}

// Add inserts an element and updates the accumulator
func (acc *HashAccumulator) Add(data []byte) {
	acc.Items = append(acc.Items, data)
	acc.Value = acc.compute()
}

// compute recomputes the hash of all elements in the accumulator
func (acc *HashAccumulator) compute() string {
	hash := sha256.New()
	for _, item := range acc.Items {
		hash.Write(item)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

// Proof returns a hash commitment (in this simple version)
// Proof returns a hash commitment (in this simple version)
func (acc *HashAccumulator) Proof(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
