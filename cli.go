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
	addBlock --data DATA "add data to blockchain"
	printChain			"print all blockchain data"
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
		fmt.Println("打印区块")

		cli.PrintBlockChain()

	default:
		fmt.Println("无效的区块，请检查")

		fmt.Printf(Usage)
	}


	// 执行相应动作
}



