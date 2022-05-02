package Byte

func Tobyte32(str []byte) [32]byte {
	var res [32]byte
	for i := 0; i < 32; i++ {
		res[i] = str[i]
	}
	return res
}
