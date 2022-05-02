package main

import (
	"fmt"
	"main/Crypto"
	"os"
)

//input: Randomnumber
//output: hash value of IDT, hash value of IDT||R

const IDT = "9p9x1ydp7uue" //12

func init() {
	if len(os.Args) != 2 {
		panic("===incorrect number of arg!===")
	}
}

func main() {
	args := os.Args[1:]

	//RandomNumber is argument
	RNum := args[0]        //randomnumber
	ByteIDT := []byte(IDT) //IDT

	//caluculate the two mentioned hash values
	HofRIDT := Crypto.GenerateORSHA256([]byte(RNum), ByteIDT) //hash of randomnumber||IDT
	HofIDT := Crypto.GenerateSHA256(ByteIDT)                  //hash of IDT

	//Print(Send) hash values
	fmt.Printf("%x\n", HofRIDT)
	fmt.Printf("%x\n", HofIDT)
}

//go run IDT.go 123321
//R||I:   a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76
//I:      2ba0018dd2c706619a72ab42b880cf05bec8afc141f68a670211fb52bf2079a2
