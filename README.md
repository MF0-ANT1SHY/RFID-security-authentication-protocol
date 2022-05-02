# RFID-security-authentication-protocol

> *A more full-featured RFID authentication protocol, the theory of this project comes from: https://cloud.tencent.com/developer/article/1035161#undefined*

## Quick Start

To use this project, you first need to start the backend program, which listens on local port 8000 by default and communicates with IDR.go using the TCP protocol. The port can be changed by modifying the source code:

```
cd ../BackEnd
go run Server.go
```

Enter the IDT folder and enter a random number into IDT. go ( the number of bits of the random number needs to be the same as the number of bits of the id of IDT and IDR, the default is 12 bits) to complete the operation of IDT to generate hash values: 

```shell
cd ./IDT
go run IDT.go [12-bit random number]
```

The program generates two strings ( **hash( random number|| IDT)** and **hash( IDT) , which are passed as arguments** to IDR.go:

```shell
cd ../IDR
go run IDR.go [Arg1] [Arg2]
```

IDR.go will complete a series of operations and communicate with the backend using TCP protocol to achieve data validation and receive the **IDT ⊕ IDR⊕ R** sent back from the backend, passing this data to IDT.go **in the form of argument** to complete the final validation: 

```shell
cd ../IDT
go run IDT2.go [random number mentioned above] [Arg2]
```

IDT2.go will output the final validation result: true or false
