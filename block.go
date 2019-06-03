package main

import (
	"crypto/sha256"
	"time"
)

// 定义结构
type Block struct {

	// 版本号
	Version uint64

	// 前区块哈希
	PrevHash []byte

	// merkel根
	MerkelRoot []byte

	// 时间戳
	TimeStamp uint64

	// 难度值
	Difficulty uint64

	// 随机数
	Nonce uint64

	// 当前区块哈希,正常比特币区块中，没有当前区块的哈希，我们为了实现方便做了简化
	Hash []byte

	// 数据
	Data []byte
}

// 辅助函数 将uint转成[]byte
func Uint64ToByte(num uint64) []byte {
	return []byte{}
}

// 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block  {
	block := Block{
		Version: 00,
		PrevHash:prevBlockHash,
		MerkelRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Difficulty:0, // 随便填写的无效值
		Nonce:0, // 同上
		Hash: []byte{},
		Data: []byte(data),
	}
	block.SetHash()
	return &block
}

// 3.生成哈希
func (block *Block) SetHash() {

	var blockInfo []byte

	// 1.拼装数据
	blockInfo = append(block.PrevHash,block.Data...) // byte拼接
	blockInfo = append(blockInfo,Uint64ToByte(block.Version)...) // byte拼接
	blockInfo = append(blockInfo,block.PrevHash...) // byte拼接
	blockInfo = append(blockInfo,block.MerkelRoot...) // byte拼接
	blockInfo = append(blockInfo,Uint64ToByte(block.TimeStamp)...) // byte拼接
	blockInfo = append(blockInfo,Uint64ToByte(block.Difficulty)...) // byte拼接
	blockInfo = append(blockInfo,Uint64ToByte(block.Nonce)...) // byte拼接
	blockInfo = append(blockInfo,block.Data...) // byte拼接
	// 2.sha256

	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}