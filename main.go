package main

import (
	"fmt"
	"time"

	"blockchain_go/models"
)

// CreateBlockchain creates a new blockchain with the given difficulty
func CreateBlockchain(difficulty int) models.Blockchain {
	genesisBlock := models.Block{
		Hash:      "0",
		Timestamp: time.Now(),
	}
	return models.Blockchain{
		GenesisBlock: genesisBlock,
		Chain:        []models.Block{genesisBlock},
		Difficulty:   difficulty,
	}
}

// AddBlock adds a block to the blockchain
func AddBlock(bc *models.Blockchain, from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := models.Block{
		Data:         blockData,
		PreviousHash: lastBlock.Hash,
		Timestamp:    time.Now(),
	}
	newBlock.Mine(bc.Difficulty)
	bc.Chain = append(bc.Chain, newBlock)
}

// IsValid validates the blockchain
func IsValid(bc *models.Blockchain) bool {
	for i := range bc.Chain[1:] {
		previousBlock := bc.Chain[i]
		currentBlock := bc.Chain[i+1]
		if currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func main() {
	bc := CreateBlockchain(2)

	AddBlock(&bc, "Alice", "Bob", 10)
	AddBlock(&bc, "Bob", "Charlie", 20)
	AddBlock(&bc, "Charlie", "Dave", 30)

	for _, block := range bc.Chain {
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Data: %v\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}

	fmt.Printf("Blockchain valid: %v\n", IsValid(&bc))
}

