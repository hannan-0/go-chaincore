package zkp

import (
	"crypto/sha256"
	"fmt"
)

// Prove simulates generating a zero-knowledge proof
func Prove(secret string) string {
	hash := sha256.Sum256([]byte(secret))
	return fmt.Sprintf("%x", hash)
}

// Verify simulates verifying the ZKP
func Verify(proof, secret string) bool {
	expected := Prove(secret)
	return proof == expected
}
