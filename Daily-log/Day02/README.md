# Day 2: Public/Private Key Cryptography

## Objective
Understand asymmetric cryptography and how blockchain wallet addresses are derived.

## What I did today
* Wrote a Go script using `crypto/ecdsa` to generate an Elliptic Curve private key.
* Derived the corresponding public key coordinates.
* Simulated a wallet address by hashing the public key bytes using SHA-256 and truncating it to 20 bytes.

## How to run the code
`go run main.go`
