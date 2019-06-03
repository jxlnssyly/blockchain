package main

import (
	"github.com/boltdb/bolt"
	"fmt"
	"log"
)

func main() {
	// 打开数据库
	db, err := bolt.Open("test.db",0600,nil)
	if err != nil {
		fmt.Println("打开数据库失败")
		return
	}

	// 将要操作数据库(改写)
	db.Update(func(tx *bolt.Tx) error {
		// 找到抽屉,如果没有就创建
		bucket := tx.Bucket([]byte("b1"))

		if bucket == nil { // 没有抽屉，需要创建
			 bucket,err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic(err)
			}

		}
		// 准备写数据
		bucket.Put([]byte("111"),[]byte("hello"))
		bucket.Put([]byte("222"),[]byte("word"))


		return nil
	})

	// 写数据

	// 读数据
}
