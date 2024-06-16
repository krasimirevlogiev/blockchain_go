package models

// Blockchain represents the blockchain itself
type Blockchain struct {
	GenesisBlock Block   `json:"genesis_block"`
	Chain        []Block `json:"chain"`
	Difficulty   int     `json:"difficulty"`
}

