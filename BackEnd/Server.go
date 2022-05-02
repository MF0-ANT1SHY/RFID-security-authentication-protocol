package main

import (
	"fmt"
	"main/Byte"
	"main/Crypto"
	"main/DataBase"
	"net"
	"strconv"
)

var Answer []byte

func main() {
	//listen port
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("LISTEN ERROR!")
		return
	}
	defer listener.Close()
	//wait for client
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connect error!")
			continue
		}
		// process client data
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	ClientAddr := conn.RemoteAddr().String()
	fmt.Printf("===New Connect: %s===\n", ClientAddr)
	//read client context
	Rbuffer := make([]byte, 1024)
	Contextlen, err1 := conn.Read(Rbuffer)
	if err1 != nil {
		fmt.Printf("%s read context error\n", ClientAddr)
	}
	//judge if data is legal
	if Contextlen <= 64 {
		fmt.Println("Data is not legal")
		return
	}
	//processing data
	HofIDT, HofRIDTxorHIDR, random := CutMessage(Rbuffer, Contextlen)
	IDT, state := DataBase.CalibrateIDT(fmt.Sprintf("%x", HofIDT))
	if state != true {
		fmt.Println(fmt.Sprintf("%x", HofIDT), " Not exist!!!")
	} else {
		HofRIDT := Crypto.GenerateORSHA256([]byte(strconv.Itoa(random)), []byte(IDT))
		HofIDR := Crypto.GenerateSHAXOR(HofRIDT, HofRIDTxorHIDR)
		IDR, _ := DataBase.CalibrateIDR(fmt.Sprintf("%x", HofIDR))
		Info := Crypto.GenerateTripleXOR([]byte(IDT), []byte(strconv.Itoa(random)), []byte(IDR))
		fmt.Printf("IDT:%x\nrandom:%x\nIDR:%x\n", []byte(IDT), []byte(strconv.Itoa(random)), []byte(IDR))
		fmt.Printf("===Send To Client:===\n%x", Info)
		conn.Write(Info)
	}
	return
}

func CutMessage(Rbuffer []byte, Contextlen int) (a [32]byte, b [32]byte, c int) {
	a = Byte.Tobyte32(Rbuffer[:32])                     //HofIDT
	b = Byte.Tobyte32(Rbuffer[32:64])                   //HofRIDTxorHIDR
	c, _ = strconv.Atoi(string(Rbuffer[64:Contextlen])) //random
	fmt.Println("===From Client:===")
	fmt.Printf("HofIDT: %x\n", a)
	fmt.Printf("HofRIDTxorHIDR: %x\n", b)
	fmt.Printf("RNum: %d\n\n", c)
	return a, b, c
}
