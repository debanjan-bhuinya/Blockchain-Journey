# Day 4: Proof of Work (Mining)

## Objective
Implement a Proof of Work consensus algorithm to understand how blockchains prevent network spam and tampering.

## What I did today
* Added a `Nonce` property to the Block struct.
* Built a `mineBlock` loop that repeatedly calculates hashes until it finds one that starts with a specific number of zeroes (based on a difficulty target).
* Successfully mined two blocks, simulating the computational work required to secure a blockchain network.

## How to run the code
`go run main.go`
