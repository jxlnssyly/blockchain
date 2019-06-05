package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

// 挖矿奖励
const Reward  = 12.5

// 定义交易结构
type Transaction struct {
	TXID []byte // 交易ID
	TXInputs []TXInput // 交易输入数组
	TXOutputs []TXOutput
}

// 定义交易输入
type TXInput struct {
	// 引用的交易ID
	TXid []byte
	// 引用的output索引值
	Index int64
	// 解锁脚本,我们用地址来模拟
	Sig string
}

type TXOutput struct {
	// 转账金额
	Value float64
	// 锁定脚本,我们用地址来模拟
	PubkeyHash string
}

// 设置交易ID
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

// 判断是否为挖矿交易
func (tx *Transaction)IsCoinbase() bool  {
	// 1.交易input只有一个

	if len(tx.TXInputs) == 1 {

		input := tx.TXInputs[0]
		// 2.交易ID为空
		// 3.交易的index为-1
		if !bytes.Equal(input.TXid,[]byte{}) || input.Index != -1{
			return false
		}
	}

	return true
}

// 提供创建交易方法



// 创建挖矿交易
func NewCoinbaseTX(address string, data string) *Transaction {
	//Transaction
	// 挖矿交易的特点：
	// 1、只有一个input
	// 2、无需引用交易id
	// 3、无需引用index

	// 矿工由于挖矿时无需指定签名，所以这个sig字段可以由矿工自由填写
	input := TXInput{[]byte{},-1,data}

	output := TXOutput{Reward,address}

	// 对于挖矿交易来说，只有一个input和一个output
	tx := Transaction{[]byte{},[]TXInput{input},[]TXOutput{output}}

	tx.SetHash()
	return &tx
}



// 根据交易调整程序






