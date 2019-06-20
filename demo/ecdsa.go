package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 演示如何使用ecdsa 生成公钥私钥
// 签名和校验
func main1() {
	// 创建曲线
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve,rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	// 生成公钥
	pubKey := privateKey.PublicKey
	data := "helloworld"
	hash := sha256.Sum256([]byte(data))

	// 签名
	r, s, err := ecdsa.Sign(rand.Reader,privateKey,hash[:])
	if err != nil {
		log.Panic(err)
	}

	// 把 r,s进行序列化传输
	signature := append(r.Bytes(), s.Bytes()...)


	// 1.定义2个辅助的big.Int
	r1 := big.Int{}
	s1 := big.Int{}

	// 2.拆分signature,平局分,前半部分给r,后半部分给s
	r1.SetBytes(signature[0 :len(signature)/2])
	s1.SetBytes(signature[len(signature)/2 :])
	// 校验
	res := ecdsa.Verify(&pubKey, hash[:], &r1, &s1)

	fmt.Println(res)



}
