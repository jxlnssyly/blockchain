package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	// 游标，用于不断索引
	currentHashPointer []byte
}

func (bc *BlockChain)NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.db,

		// 最初指向区块的最后一个区块，随着Next的调用，不断变化
		bc.tail,
	}
}

// 返回当前的区块
// 指针前移
func (it *BlockChainIterator)Next() *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("迭代器遍历时，bucket不该为空")
		}

		blockTmp := bucket.Get(it.currentHashPointer)
		block = Deserialize(blockTmp)
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return &block
}





