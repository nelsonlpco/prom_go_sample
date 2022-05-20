package domain

import "fmt"

func XOR(b1 uint8, b2 uint8) (result uint8) {
	strb1 := fmt.Sprintf("%08b", b1)
	strb2 := fmt.Sprintf("%08b", b2)

	for i := 0; i < 8; i++ {
		sb1 := strb1[i] == '1'
		sb2 := strb2[i] == '1'

		var rb uint8 = 0
		if sb1 && !sb2 || !sb1 && sb2 {
			rb = 1
		}

		result <<= 1
		result |= rb
	}

	return result
}

func Cript(msg string, key string) (product string) {
	size := len(msg)
	for i := 0; i < size; i++ {
		r := XOR(msg[i], key[i])
		product = fmt.Sprintf("%s%c", product, r)
	}
	return product
}

func Decript(key string, product string) (msg string) {
	size := len(key)
	for i := 0; i < size; i++ {
		r := XOR(key[i], product[i])
		msg = fmt.Sprintf("%s%c", msg, r)
	}
	return msg
}
