package main

import (
	"encoding/hex"
	"fmt"
	"main/Byte"
	"main/Crypto"
	"net"
	"os"
	"strconv"
)

//input: hash of idt, hash of idt||randomnumber
//output: hash of idt, (hash of idt||r)XOR(hash of idr), randomnumber

//input: idt XOR idr XOR randomnumber
//output: hash of randomnumberXORidt

const IDR = "nh679sg5x6wl"

var R = 123456789123
var HofRIDT []byte
var HofIDT []byte

func init() {
	if len(os.Args) != 3 {
		panic("===incorrect number of arg!===")
	}
}

func main() {
	//read arguments
	args := os.Args[1:]
	ByteIDR := []byte(IDR)                 //IDR
	HofRIDT, _ = hex.DecodeString(args[0]) //hash of idt||randomnumber
	HofIDT, _ = hex.DecodeString(args[1])  //hash of idt
	fmt.Printf("===receive===\nHofRIDT:\t%x\nHofIDT:\t\t%x\n", HofRIDT, HofIDT)

	//generate (hash of idr)XOR(hash of idt||randomnumber)
	HofIDR := Crypto.GenerateSHA256(ByteIDR)                                //hash of idr
	HofRIDTxorHIDR := Crypto.GenerateSHAXOR(Byte.Tobyte32(HofRIDT), HofIDR) // (...)XOR(...)

	//generate info
	Info := append(HofIDT, HofRIDTxorHIDR[:]...)
	Info = append(Info, []byte(strconv.Itoa(R))...)
	fmt.Printf("===send===\nHofIDT:\t\t%x\nHofRIDTxorHIDR:\t%x\nR:\t\t%d\n", HofIDT, HofRIDTxorHIDR, R)

	//send info to the backend
	//connect to the backend
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	//send info
	_, err = conn.Write(Info)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
	fmt.Printf("write to server = %x\n", Info)

	//get reply
	reply := make([]byte, 1024)
	Replylen, err1 := conn.Read(reply)
	if err1 != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
	fmt.Printf("===recieve from backend===\nreply from server=%s\n", fmt.Sprintf("%x", reply[:Replylen]))

	//processing reply
	triplexor := reply[:Replylen]
	IDTxorR := Crypto.GenerateXOR(triplexor, ByteIDR)
	res := Crypto.GenerateSHA256(IDTxorR)
	fmt.Printf("IDR:%x\nSend to IDT: %x\n", ByteIDR, res)
	conn.Close()
}

//go run IDR.go a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76 2ba0018dd2c706619a72ab42b880cf05bec8afc141f68a670211fb52bf2079a2
//===receive===
//HofRIDT:        a320480f534776bddb5cdb54b1e93d210a3c7d199e80a23c1b2178497b184c76
//HofIDT:         2ba0018dd2c706619a72ab42b880cf05bec8afc141f68a670211fb52bf2079a2
//===send===
//HofIDT:         2ba0018dd2c706619a72ab42b880cf05bec8afc141f68a670211fb52bf2079a2
//HofRIDTxorHIDR: bf7cc49f99dd75671d7d0074d91e10c93fb19f38a5104d849fac207cfb5f4994
//R:              123321
