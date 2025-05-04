package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

func GenerateVRFSeed(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func SelectVRFLeader(seed string, totalNodes int) int {
	hash := sha256.Sum256([]byte(seed))
	num := big.NewInt(0).SetBytes(hash[:])
	return int(num.Int64() % int64(totalNodes))
}
