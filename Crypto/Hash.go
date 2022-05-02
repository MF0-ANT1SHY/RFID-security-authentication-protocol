package Crypto

import (
	"crypto/sha256"
	"strconv"
)

func GenerateSHA256(a []byte) [32]byte {
	Res := sha256.Sum256(a)
	return Res
}

func GenerateORSHA256(a, b []byte) [32]byte {
	Aint, _ := strconv.Atoi(string(a))
	Bint, _ := strconv.Atoi(string(b))
	Rint := Aint | Bint
	Res := sha256.Sum256([]byte(strconv.Itoa(Rint)))
	return Res
}

func GenerateSHAXOR(AimStr, Key [32]byte) [32]byte {
	var Res [32]byte
	for i := 0; i < 32; i++ {
		Res[i] = AimStr[i] ^ Key[i]
	}
	return Res
}

func GenerateXOR(AimStr, Key []byte) []byte {
	AimStrLen := len(AimStr)
	KeyLen := len(Key)
	var Res = make([]byte, AimStrLen)
	for i := 0; i < AimStrLen; i++ {
		Res[i] = AimStr[i] ^ Key[i%KeyLen]
	}
	return Res
}

func GenerateTripleXOR(a, b, c []byte) []byte {
	ab := GenerateXOR(a, b)
	abc := GenerateXOR(ab, c)
	return abc
}
