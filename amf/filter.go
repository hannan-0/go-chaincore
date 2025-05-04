package amf

import (
	"crypto/sha256"
	"encoding/binary"
)

type BloomFilter struct {
	Size uint
	Bits []bool
	K    uint // number of hash functions
}

func NewBloomFilter(size uint, k uint) *BloomFilter {
	return &BloomFilter{
		Size: size,
		Bits: make([]bool, size),
		K:    k,
	}
}

func (bf *BloomFilter) Add(data []byte) {
	hash := sha256.Sum256(data)
	for i := uint(0); i < bf.K; i++ {
		index := bf.hashIndex(hash[:], i)
		bf.Bits[index%bf.Size] = true
	}
}

func (bf *BloomFilter) Contains(data []byte) bool {
	hash := sha256.Sum256(data)
	for i := uint(0); i < bf.K; i++ {
		index := bf.hashIndex(hash[:], i)
		if !bf.Bits[index%bf.Size] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) hashIndex(hash []byte, i uint) uint {
	salt := make([]byte, 4)
	binary.LittleEndian.PutUint32(salt, uint32(i))
	salted := append(hash, salt...)
	sum := sha256.Sum256(salted)
	return uint(binary.LittleEndian.Uint32(sum[:4]))
}
