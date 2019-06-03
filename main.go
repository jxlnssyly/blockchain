package main

import (
	"fmt"
	"crypto/sha256"
)

// 定义结构
type Block struct {
	// 前区块哈希
	PrevHash []byte

	// 当前区块哈希
	Hash []byte

	// 数据
	Data []byte
}

// 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block  {
	block := Block{
		PrevHash:prevBlockHash,
		Hash: []byte{}, // TODO 修改
		Data: []byte(data),
	}
	block.SetHash()
	return &block
}

// 3.生成哈希
func (block *Block) SetHash() {
	// 1.拼装数据
	blockInfo := append(block.PrevHash,block.Data...) // byte拼接
	// 2.sha256

	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

// 引入区块链
type BlockChain struct {
	// 定一个区块链数组
	blocks []*Block
}

// 定义一个区块链
func NewBlockChain() *BlockChain {
	// 创建一个创世块，并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()

	return &BlockChain{
		blocks:[]*Block{genesisBlock},

	}
}

// 创世块
func GenesisBlock() *Block {
	return NewBlock("Go 创世块",[]byte{})

}


func main() {
	bc := NewBlockChain()

	for i, block := range bc.blocks {
		fmt.Printf("===== 当前区块高度: %d ====== \n",i)
		fmt.Printf("前区块哈希: %x\n",block.PrevHash)
		fmt.Printf("当前区块哈希: %x\n",block.Hash)
		fmt.Printf("区块数据: %s\n",block.Data)
	}


}
