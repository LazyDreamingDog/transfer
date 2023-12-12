package core

import "transfer/core/types"

type BlockChain struct {
	blocks []*types.Block
}

// AddBlock adds a block to the blockchain
func (bc *BlockChain) AddBlock(block *types.Block) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	// newBlock := NewBlock(data, prevBlock.Hash)
	// bc.blocks = append(bc.blocks, newBlock)

}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *types.Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockChain creates a new BlockChain with genesis Block
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
