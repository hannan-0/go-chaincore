package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

// Transaction represents a basic blockchain transaction
type Transaction struct {
	ID        string
	From      string
	To        string
	Amount    float64
	Timestamp int64
	Signature string
}

// Mempool holds unconfirmed transactions
var Mempool []Transaction

// NewTransaction creates a new transaction and generates its ID
func NewTransaction(from, to string, amount float64, signature string) Transaction {
	t := Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now().UnixNano(),
		Signature: signature,
	}
	t.ID = t.Hash()
	return t
}

// Hash creates a unique transaction hash (ID)
func (t Transaction) Hash() string {
	data := fmt.Sprintf("%s:%s:%f:%d:%s", t.From, t.To, t.Amount, t.Timestamp, t.Signature)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// ValidateTransaction checks basic rules (e.g., amount > 0, signature present)
func ValidateTransaction(t Transaction) error {
	if t.Amount <= 0 {
		return errors.New("amount must be positive")
	}
	if t.From == "" || t.To == "" {
		return errors.New("sender and receiver must be specified")
	}
	if t.Signature == "" {
		return errors.New("missing signature")
	}
	return nil
}

// AddToMempool validates and appends the transaction to the mempool
func AddToMempool(t Transaction) error {
	if err := ValidateTransaction(t); err != nil {
		return err
	}
	Mempool = append(Mempool, t)
	fmt.Printf("✅ Transaction %s added to mempool\n", t.ID)
	return nil
}

// ClearMempool resets the mempool after mining
func ClearMempool() {
	Mempool = []Transaction{}
	fmt.Println("🧹 Mempool cleared after block addition.")
}
