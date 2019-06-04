package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"fmt"
)

type Person struct {
	Name string
	age uint
}

func main() {
	// 定义一个结构Person
	var xiaoming Person
	xiaoming.Name = "小明"
	xiaoming.age = 99

	// 编码数据放到buffer
	var buffer bytes.Buffer

	// 定义一个编码器，使用编码器进行编码
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&xiaoming)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("编码后:%x\n",buffer.Bytes())
	var xm Person
	// 定义一个解码器，使用解码器解码
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	decoder.Decode(&xm)

	fmt.Println(xm)

}
