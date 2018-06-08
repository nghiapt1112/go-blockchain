package chain

import (
    "../block"
)

type Blockchain struct {
    blocks []*block.Block
}

func (bc *Blockchain) GetBlock() []*block.Block {
    return bc.blocks;
}
func (b *Blockchain) AddBlock(data string) {
    prevBlock := b.blocks[len(b.blocks)-1]
    newBlock := block.NewBlock(data, prevBlock.Hash)
    b.blocks = append(b.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
    return &Blockchain{[]*block.Block{block.NewGenesisBlock()}}
}
