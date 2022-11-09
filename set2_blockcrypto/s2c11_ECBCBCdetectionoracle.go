/* Set 2 Challenge 11 - An ECB/CBC detection oracle */

package set2_blockcrypto

import (
	"crypto/aes"
	"crypto/rand"
	mrand "math/rand"
	"math/big"
	"bytes"
	"time"
	"fmt"

	set1 "wryhder/cryptopals-crypto-challenges/set1_basics"
)

// Generate random bytes of specified size/length
func generateRandomBytes(size int) []byte {
	randBytes := make([]byte, size)
	_, err := rand.Read(randBytes)
	if err != nil {
		fmt.Println("error: ", err)
		return nil
	}
	return randBytes
}

// Generate a random integer within given range (limits inclusive)
func generateRandInt(lowerLimit, upperLimit int) int {
	mrand.Seed(time.Now().UnixNano())
	randInt := mrand.Intn(upperLimit - lowerLimit + 1) + lowerLimit

	return randInt
}

func padStartAndEnd(plaintext string) []byte {
	padLength := generateRandInt(5, 10)
	padValue := byte(padLength)
	padBytes := bytes.Repeat([]byte{padValue}, padLength)

	var paddedPlaintext []byte
	// Prepend pad bytes
	paddedPlaintext = append(padBytes, []byte(plaintext)...)
	// Append pad bytes
	paddedPlaintext = append(paddedPlaintext, padBytes...)

	return paddedPlaintext
}

func selectECBOrCBC() string {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(2))
	mode := ""

	switch randInt.Int64() {
		case 0: mode = "ECB"
		case 1: mode = "CBC"
	}

	return mode
}

func encryptAES128_ECBOrCBC(plaintext string) (string, string) {
	blockSize := aes.BlockSize
	key := string(generateRandomBytes(blockSize))
	paddedPlaintext := string(padStartAndEnd(plaintext))

	mode := ""
	encrypted := ""

	switch selectECBOrCBC() {
		case "ECB":
			mode = "ECB"
			encrypted = set1.EncryptAES128_ECB(paddedPlaintext, key)
		case "CBC":
			mode = "CBC"
			IV := generateRandomBytes(blockSize)
			encrypted = EncryptAES128_CBC(paddedPlaintext, key, IV)
	}

	return mode, encrypted
}

func ECB_CBCDetectionOracle(block string) string {
	// Analyze ciphertext block for the encryption algorithm used
	return ""
}