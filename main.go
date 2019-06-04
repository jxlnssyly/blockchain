package main

func main() {
	bc := NewBlockChain()

	bc.AddBlock("班长向班花赚了50枚比特币")
	bc.AddBlock("班长又向班花赚了50枚比特币")
	/*

for i, block := range bc.blocks {
fmt.Printf("===== 当前区块高度: %d ====== \n",i)
fmt.Printf("前区块哈希: %x\n",block.PrevHash)
fmt.Printf("当前区块哈希: %x\n",block.Hash)
fmt.Printf("区块数据: %s\n",block.Data)
}
*/
}
