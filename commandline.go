package main

import "fmt"

func (cli *CLI)AddBlock(data string)  {
	//cli.bc.AddBlock(data)
}

func (cli *CLI)PrintBlockChain()  {

	it := cli.bc.NewIterator()
	for  {
		block := it.Next()
		fmt.Printf("=====  ====== \n")
		fmt.Printf("前区块哈希: %x\n",block.PrevHash)
		fmt.Printf("当前区块哈希: %x\n",block.Hash)
		fmt.Printf("区块数据: %s\n",block.Transactions[0].TXInputs[0].Sig)

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CLI)Send(from, to string, amount float64, miner, data string)  {
	fmt.Println(from,to,amount,miner,data)
}