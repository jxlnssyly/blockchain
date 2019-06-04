package main

import "fmt"

func main() {
	bc := NewBlockChain()

	bc.AddBlock("班长向班花赚了50枚比特币")
	bc.AddBlock("班长又向班花赚了50枚比特币")

	it := bc.NewIterator()
	for  {
		block := it.Next()
		fmt.Printf("=====  ====== \n")
		fmt.Printf("前区块哈希: %x\n",block.PrevHash)
		fmt.Printf("当前区块哈希: %x\n",block.Hash)
		fmt.Printf("区块数据: %s\n",block.Data)

		if len(block.PrevHash) == 0 {
			break
		}
	}


}
