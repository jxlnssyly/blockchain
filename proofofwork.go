package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

// 定义一个工作量证明的结构
type ProofOfWork struct {
	// block
	block *Block
	// 目标值
	// 一个非常大的数，它有很多丰富的方法：比较、赋值方法
	target *big.Int
}

// 提供创建POW的函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block:block,
	}
	// 指定的难度值，现在是一个string类型，需要转换
	targetStr := "0000f000000000000000000000000000000000000000000000000000000000000"
	// 引入的辅助变量，目的是将上面的难度值转换成big.Int
	tmpInt := big.Int{}
	// 将难度赋值给bigInt,指定16进制
	tmpInt.SetString(targetStr,16)
	pow.target = &tmpInt
	return &pow
}

// 提供不断计算hash的函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	var nonce uint64
	block := pow.block
	var hash [32]byte
	for {
		// 拼装数据(区块的数据，还有不断变化的随机数)
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		// 将二维切片数组连接起来，返回一个一维切片
		blockInfo := bytes.Join(tmp, []byte{})

		// 做哈希运算
		hash = sha256.Sum256(blockInfo)
		// 与pow中的target进行比较
		tmpInt := big.Int{}
		// 将我们得到的hash数组转换成bigIng
		tmpInt.SetBytes(hash[:])

		// 比较当前的哈希值与目标哈希值
		if tmpInt.Cmp(pow.target) == -1 { // 找到了

			fmt.Printf("挖矿成功 hash:%x nonce: %d \n",hash,nonce)
			break

		} else { // 没找到
			nonce++
		}
	}

	return hash[:],nonce
}
