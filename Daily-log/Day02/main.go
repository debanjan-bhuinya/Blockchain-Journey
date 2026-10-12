package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	fmt.Println("--- Asymmetric Cryptography & Wallets ---")

	// 1. Generate Private Key (The secret you never share)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1. Private Key (DO NOT SHARE):\n   %x\n\n", privateKey.D)

	// 2. Derive Public Key (Can be shared with anyone)
	publicKey := privateKey.PublicKey
	fmt.Printf("2. Public Key (X, Y coordinates):\n   X: %x\n   Y: %x\n\n", publicKey.X, publicKey.Y)

	// 3. Generate Wallet Address (Hash of the public key)
	pubKeyBytes := append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)
	addressHash := sha256.Sum256(pubKeyBytes)
	
	fmt.Printf("3. Simulated Wallet Address:\n   0x%x\n", addressHash[:20])
}
