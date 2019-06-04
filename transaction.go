package main

// 定义交易结构
type Transaction struct {
	TXID []byte // 交易ID
	TXIinputs []TXInput // 交易输入数组
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


// 提供创建交易方法



// 创建挖矿交易



// 根据交易调整程序






