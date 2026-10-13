package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// 1. Define what a Block looks like
type Block struct {
	Timestamp int64
	Data      string
	PrevHash  string // This is the glue that holds the chain together!
	Hash      string
}

// 2. A helper function to create the fingerprint (hash) of a block
func calculateHash(timestamp int64, data string, prevHash string) string {
	// Combine all the block's data into one single string
	record := fmt.Sprintf("%d%s%s", timestamp, data, prevHash)
	
	// Hash it using SHA-256
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	fmt.Println("--- Building a Blockchain ---")

	// --- THE GENESIS BLOCK (Block 1) ---
	// The first block has no previous block, so its PrevHash is usually just "0"
	genesisTime := time.Now().Unix()
	genesisData := "Genesis Block - The beginning of the chain!"
	genesisPrevHash := "0"
	genesisHash := calculateHash(genesisTime, genesisData, genesisPrevHash)

	block1 := Block{genesisTime, genesisData, genesisPrevHash, genesisHash}
	
	fmt.Printf("BLOCK 1 (Genesis)\nData: %s\nHash: %s\n\n", block1.Data, block1.Hash)

	// --- THE SECOND BLOCK ---
	block2Time := time.Now().Unix()
	block2Data := "Alice sends 10 ETH to Bob"
	
	// THE MAGIC HAPPENS HERE: We grab Block 1's hash and store it in Block 2
	block2PrevHash := block1.Hash 
	
	block2Hash := calculateHash(block2Time, block2Data, block2PrevHash)

	block2 := Block{block2Time, block2Data, block2PrevHash, block2Hash}
	
	fmt.Printf("BLOCK 2\nData: %s\nPrevious Hash: %s\nHash: %s\n", block2.Data, block2.PrevHash, block2.Hash)
}
