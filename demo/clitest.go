package main

import (
	"os"
	"fmt"
)

func main2() {
	//len1 := len(os.Args)
	for i, cmd := range os.Args {
		fmt.Printf("arg[%d] : %s\n",i,cmd)
	}
}
