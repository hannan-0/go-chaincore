package blockchain

import (
	"errors"
	"fmt"
)

// SmartContract represents a basic contract with executable logic
type SmartContract struct {
	Address string
	Owner   string
	Code    func(args ...string) (string, error) // Executable function
}

// ContractRegistry holds deployed contracts mapped by address
var ContractRegistry = map[string]SmartContract{}

// DeployContract registers a contract with the system
func DeployContract(address, owner string, code func(args ...string) (string, error)) error {
	if _, exists := ContractRegistry[address]; exists {
		return errors.New("contract already exists at this address")
	}
	ContractRegistry[address] = SmartContract{
		Address: address,
		Owner:   owner,
		Code:    code,
	}
	fmt.Printf("🚀 Contract deployed at %s by %s\n", address, owner)
	return nil
}

// ExecuteContract invokes the logic of a deployed contract
func ExecuteContract(address string, args ...string) (string, error) {
	contract, exists := ContractRegistry[address]
	if !exists {
		return "", errors.New("contract not found")
	}
	result, err := contract.Code(args...)
	if err != nil {
		return "", fmt.Errorf("contract execution failed: %v", err)
	}
	fmt.Printf("⚙️  Executed contract at %s with result: %s\n", address, result)
	return result, nil
}
