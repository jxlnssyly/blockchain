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
	spentOutputs := make(map[string][]int64)
	// 1.遍历区块链
	// 创建迭代器
	it := bc.NewIterator()
	for {
		block := it.Next()
		// 2.遍历交易

		for _, tx := range block.Transactions {
			fmt.Printf("current txid : %x\n",tx.TXID)
		LABEL:

		// 3.遍历Output,找到和自己相关的utxo(在添加output之前检查一下是否已经消耗过)
			for i, output := range tx.TXOutputs {
				//fmt.Println(i)

				// 将所有消耗过的outputs和当前的所即将添加的output对比一下
				// 如果当前的相同，则跳过，否则继续
				// 如果当前的交易ID，存在于我们已经标识的map，那么说明这个事李四消耗过的output
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j {
							// 当前准备添加output已经消耗过了，不要再加了
							continue LABEL
						}
					}
				}

				// 这个output和我们目标的地址相同，满足条件，加到返回utxo数组中
				if output.PubkeyHash == address {
					UTXO = append(UTXO,output)
				}
			}

			// 如果当前交易是挖矿交易的话，那么不遍历，直接跳过
			if tx.IsCoinbase() {
				// 4.遍历input，找到自己花费过的UTXO集合(把自己消耗的标识出来)
				// 定义Map来保存消费过的output，key是这个output的交易ID，value是这个交易中索引的数组
				for _,input := range tx.TXInputs {
					// 判断一下当前input和目标(李四)是否一致，如果相同，说明这个是李四消耗过的output，就加进来
					if input.Sig == address {
						indexArray := spentOutputs[string(tx.TXID)]
						indexArray = append(indexArray,input.Index)
					}
				}
			}
		}

		if len(block.PrevHash) == 0 {
			fmt.Println("区块遍历完成，退出!")
			break
		}
	}

	return UTXO
}

func (bc *BlockChain)FindNeedUTXOs(from string, amount float64)(map[string][]uint64, float64 ) {
	// 找到的合理的UTXO集合
	var utxos map[string][]uint64
	var calc float64 // 找到的utxos里面包含的钱的总数
	// TODO

	return utxos, calc
}

