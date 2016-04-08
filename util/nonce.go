package util

import "math/rand"

const nonceChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Nonce(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = nonceChars[rand.Int63() % int64(len(nonceChars))]
	}
	return b
}

