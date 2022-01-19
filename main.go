package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string) {

	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {

	chain := InitBlockchain()

	chain.AddBlock("--FirtsBlock after genesis--")
	chain.AddBlock("--Second Block after genesis--")
	chain.AddBlock("Third Block after genesis--")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n\n", block.Hash)
	}
}
