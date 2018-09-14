package main

import (
	"fmt"
	. "github.com/aidamina/cryptocoin/core"
	"encoding/hex"
)



func main (){
	genesis := NewBlock()

	fmt.Println("starting")
	block := NewBlock()
	fmt.Println("block", genesis)

	fmt.Println("block", block)
	fmt.Println("block", hex.EncodeToString(CalculateHash([]byte("test"))))
}