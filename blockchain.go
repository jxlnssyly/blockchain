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

const blockChainDB = "bc.db"
const blockBucket = "blockBucket"

// 定义一个区块链
func NewBlockChain(address string) *BlockChain {
	// 创建一个创世块，并作为第一个区块添加到区块链中

	// 最后一个hash
	var lastHash []byte

	// 打开数据库
	db, err := bolt.Open(blockChainDB, 0600, nil)

	//defer db.Close()

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
			genesisBlock := GenesisBlock(address)
			// hash作为key，block字节流作为hash
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte("LastHash"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("LastHash"))
		}

		return nil
	})
	return &BlockChain{db, lastHash}
}

// 创世块
func GenesisBlock(address string) *Block {
	coinBase := NewCoinbaseTX(address,"Go 创世块")
	return NewBlock([]*Transaction{coinBase}, []byte{})
}

// 添加区块
func (bc *BlockChain) AddBlock(txs []*Transaction) {
	db := bc.db // 数据库

	lastHash := bc.tail //最后一条哈希
	db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应该为空，请检查")
		}
		// 创建新区块
		block := NewBlock(txs,lastHash)

		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHash"), block.Hash)

		// 更新内存中的最后一个区块
		bc.tail = block.Hash
		return nil
	})
}

// 找到指定地址的所有的UTXO
func (bc *BlockChain)FindUTXOs(address string) []TXOutput  {
	var UTXO []TXOutput

	// TODO 找到指定地址的所有的UTXO

	return UTXO
}
