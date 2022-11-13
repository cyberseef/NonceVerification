package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

func BlockHash(data string) [32]byte {
	hash := sha256.Sum256([]byte(data))
	return hash
}

func NonceFind(X string, Y [32]byte) int {
	for nonce := 0; nonce < 1000000; nonce++ {
		if NonceVerify(nonce, X, Y) {
			return nonce
		}
	}
	return -1
}

func NonceVerify(nonce int, X string, Y [32]byte) bool {
	hash := BlockHash(string(nonce) + X)
	return hash == Y
}

func main() {
	nonce := rand.Intn(300)
	X := "hello Snakes Friends."
	Y := BlockHash(string(nonce) + X)
	fmt.Println("Nonce Value:", nonce)
	verifyN := NonceVerify(nonce, X, Y)
	fmt.Println("Status:", verifyN)
	findN := NonceFind(X, Y)
	fmt.Println("Find:", findN)
}
