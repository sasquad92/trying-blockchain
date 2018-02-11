package main

import (
	"time"
)

// Block is a structure that represents simple
// block unit in a blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock is a constructor for a Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock function creates new
// genesis (first) block in a blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}
