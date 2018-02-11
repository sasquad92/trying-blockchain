package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block is a structure that represents simple
// block unit in a blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// SetHash sets hash for a Block structure
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock is a constructor for a Block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock function creates new
// genesis (first) block in a blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}
