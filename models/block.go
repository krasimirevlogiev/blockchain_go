package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	PreviousHash string                 `json:"previous_hash"`
	Hash         string                 `json:"hash"`
	Timestamp    time.Time              `json:"timestamp"`
	Data         map[string]interface{} `json:"data"`
	PoW          int                    `json:"pow"`
}

// CalculateHash calculates the hash of the block
func (b *Block) CalculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PreviousHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.PoW)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

// Mine mines the block with the given difficulty
func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.PoW++
		b.Hash = b.CalculateHash()
	}
}


