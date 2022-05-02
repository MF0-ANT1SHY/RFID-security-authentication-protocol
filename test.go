package main

import (
	"fmt"
	"main/Crypto"
)

func main() {
	//str := "662a3c7c3a3b32777c703038"
	//byteStr, _ := hex.DecodeString(str)
	//fmt.Printf("%x\n", byteStr)
	idt := []byte("9p9x1ydp7uue")
	idr := []byte("nh679sg5x6wl")
	random := []byte("123321123321")
	res := Crypto.GenerateTripleXOR(idr, idt, random)
	ridt := Crypto.GenerateXOR(res, idr)
	fmt.Printf("%x",Crypto.GenerateSHA256(ridt))
}
