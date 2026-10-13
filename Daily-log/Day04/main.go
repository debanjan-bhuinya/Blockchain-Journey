package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

// 1. Add 'Nonce' to our block structure
type Block struct {
	Timestamp int64
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int // The counter we change to find the correct hash
}

// 2. Hash function now includes the Nonce
func calculateHash(timestamp int64, data string, prevHash string, nonce int) string {
	record := fmt.Sprintf("%d%s%s%d", timestamp, data, prevHash, nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 3. The Mining Function (Proof of Work)
func mineBlock(timestamp int64, data string, prevHash string, difficulty int) Block {
	// Create a string of target zeroes (e.g., difficulty 4 = "0000")
	targetPrefix := strings.Repeat("0", difficulty)
	nonce := 0
	hash := ""

	fmt.Printf("Mining block with data: '%s'...\n", data)

	// Keep looping and guessing the nonce until the hash starts with our target zeroes
	for {
		hash = calculateHash(timestamp, data, prevHash, nonce)
		if strings.HasPrefix(hash, targetPrefix) {
			fmt.Printf("Block Mined! Nonce found: %d\n", nonce)
			break
		}
		nonce++ // Guess again
	}

	return Block{timestamp, data, prevHash, hash, nonce}
}

func main() {
	fmt.Println("--- Proof of Work (Mining) Demo ---")

	// Higher difficulty = more zeroes = takes longer to mine
	difficulty := 4 

	// Mine the Genesis Block
	genesisBlock := mineBlock(time.Now().Unix(), "Genesis Block", "0", difficulty)
	fmt.Printf("Hash: %s\n\n", genesisBlock.Hash)

	// Mine Block 2
	block2 := mineBlock(time.Now().Unix(), "Alice sends 10 ETH to Bob", genesisBlock.Hash, difficulty)
	fmt.Printf("Hash: %s\n\n", block2.Hash)
}
