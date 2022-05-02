package main

import (
	"encoding/hex"
	"fmt"
	"main/Byte"
	"main/Crypto"
	"os"
)

//input: RandomNumber, HRxorIDT
//output: True/False
func main() {
	// RandomNumber and HRxorIDT is argument
	args := os.Args[1:]

	//input
	RNum := args[0]                             //randomnumber
	var ByteIDT = []byte("9p9x1ydp7uue")        //IDT
	var HRxorIDT, _ = hex.DecodeString(args[1]) //HRxorIDT
	fmt.Printf("random:%x\nIDT:%x\n", []byte(RNum), ByteIDT)

	//calculate calibration value
	CaliXOR := Crypto.GenerateXOR(ByteIDT, []byte(RNum))
	CaliHRxorIDT := Crypto.GenerateSHA256(CaliXOR)

	//check if
	if CaliHRxorIDT == Byte.Tobyte32(HRxorIDT) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
		fmt.Printf("%x\n%x\n", CaliHRxorIDT, HRxorIDT)
	}
}
