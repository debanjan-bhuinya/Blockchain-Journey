package main

import (
	"crypto/sha256"
	"fmt"
)

func generateHash(input string) {
	hash := sha256.Sum256([]byte(input))
	fmt.Printf("Data: %-18s | Hash: %x\n", input, hash)
}

func main() {
	fmt.Println("--- Blockchain Immutability Demo ---")
	
	// Original transaction
	generateHash("Alice sends 5 ETH")
	
	// Tampered transaction (changed 5 to 6)
	// Notice how the resulting hash is completely unrecognizable from the first
	generateHash("Alice sends 6 ETH")
}
