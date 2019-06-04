package main

import (
	"github.com/boltdb/bolt"
	"fmt"
	"log"
)

// 引入区块链
type BlockChain struct {
	// 定一个区块链数组
	//blocks []*Block
	db   *bolt.DB
	tail []byte // 存储最后一个区块的哈希
}

const blockChainDB = "blockchain.db"
const blockBucket = "blockBucket"

// 定义一个区块链
func NewBlockChain() *BlockChain {
	// 创建一个创世块，并作为第一个区块添加到区块链中

	//return &BlockChain{
	//	blocks:[]*Block{genesisBlock},
	//
	//}

	// 最后一个hash
	var lastHash []byte

	// 打开数据库
	db, err := bolt.Open(blockChainDB, 0600, nil)

	defer db.Close()

	if err != nil {
		fmt.Println("打开数据库失败")
		return nil
	}

	// 将要操作数据库(改写)
	// 写数据
	db.Update(func(tx *bolt.Tx) error {
		// 找到抽屉,如果没有就创建
		bucket := tx.Bucket([]byte(blockBucket))

		if bucket == nil { // 没有抽屉，需要创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
			}
			genesisBlock := GenesisBlock()
			// hash作为key，block字节流作为hash
			bucket.Put(genesisBlock.Hash, genesisBlock.toByte())
			bucket.Put([]byte("LastHash"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			bucket.Get([]byte("LastHash"))
		}

		return nil
	})
	return &BlockChain{db, lastHash}
}

// 创世块
func GenesisBlock() *Block {
	return NewBlock("Go 创世块", []byte{})
}

// 添加区块
func (bc *BlockChain) AddBlock(data string) {
	/*

		// 获取前区块的哈希
		lastBlock := bc.blocks[len(bc.blocks) - 1]
		prevHash := lastBlock.Hash

		// a.创建新的区块
		block := NewBlock(data, prevHash)

		// b.添加到区块链数组中
		bc.blocks = append(bc.blocks, block)
		*/
}
