package main

import (
	"os"
	"fmt"
)

// 这是一个用来接收命令行参数并且控制区块链操作的文件
type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA "添加区块链"
	printChain			"正向打印区块链"
	printChainR			"反向打印区块链"
	getBalance --address ADDRESS
`


// 接受参数的动作，我们放到一个函数中
func (cli *CLI)Run()  {
	// 得到所有的命令
	args := os.Args
	if len(args) < 2 {
		fmt.Print(Usage)
		return
	}



	// 分析命令
	cmd := args[1]
	switch cmd {
	case "addBlock":
		// 添加区块
		fmt.Println("添加区块")

		// 获取数据
		if len(args) == 4 && args[2] == "--data" {
			// 获取命令的数据
			data := args[3]
			// 添加区块
			cli.AddBlock(data)
		} else {
			fmt.Println("添加区块数据使用参数不当，请检查")
			fmt.Println(Usage)
		}


	case "printChain":
		fmt.Println("正向打印区块")

		cli.PrintBlockChain()
	case "printChainR":
		fmt.Println("反向打印区块")

		cli.PrintBlockChain()
	case "getBalance":
		fmt.Println("获取余额")
		// 获取数据
		if len(args) == 4 && args[2] == "--address" {
			// 获取命令的数据
			address := args[3]
			// 添加区块
			cli.GetBalance(address)
		} else {
			fmt.Println("添加区块数据使用参数不当，请检查")
			fmt.Println(Usage)
		}

	default:
		fmt.Println("无效的区块，请检查")

		fmt.Printf(Usage)
	}


	// 执行相应动作
}

func (cli *CLI) GetBalance(address string) {
	utxos := cli.bc.FindUTXOs(address)
	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Println(address," 余额为: ",total)
}



