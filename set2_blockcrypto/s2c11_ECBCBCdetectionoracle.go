/* Set 2 Challenge 11 - An ECB/CBC detection oracle */

package set2_blockcrypto

import (
	"crypto/aes"
	"crypto/rand"
	mrand "math/rand"
	"math/big"
	"bytes"
	"time"
	"encoding/hex"

	set1 "wryhder/cryptopals-crypto-challenges/set1_basics"
	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

// Generate a random integer within given range (limits inclusive)
// Used to choose amount of padding for plaintext 
func generateRandInt(lowerLimit, upperLimit int) int {
	mrand.Seed(time.Now().UnixNano())
	randInt := mrand.Intn(upperLimit - lowerLimit + 1) + lowerLimit

	return randInt
}

// Pad plaintext on both ends with 5-10 bytes (count chosen randomly)
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

// Choose whether to encrypt under ECB or CBC (random selection)
func selectECBOrCBC() string {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(2))
	mode := ""

	switch randInt.Int64() {
		case 0: mode = "ECB"
		case 1: mode = "CBC"
	}

	return mode
}

// Encrypt plaintext under ECB or CBC (random selection)
func encryptAES128_ECBOrCBC(plaintext string) (string, string) {
	blockSize := aes.BlockSize
	key := string(utils.GenerateRandomBytes(blockSize))
	paddedPlaintext := string(padStartAndEnd(plaintext))

	mode := ""
	encrypted := ""

	switch selectECBOrCBC() {
		case "ECB":
			mode = "ECB"
			// Add PKCS7padding here; our ECB encrypting function does not pad plaintext
			// and panics if provided with inadequate input (less than multiple of block size)
			paddedPlaintext = string(PKCS7padding([]byte(paddedPlaintext), blockSize))
			encrypted = set1.EncryptAES128_ECB(paddedPlaintext, key)
		case "CBC":
			mode = "CBC"
			IV := utils.GenerateRandomBytes(blockSize)
			encrypted = EncryptAES128_CBC(paddedPlaintext, key, IV)
	}

	return mode, encrypted
}

// Checks if ciphertext is encrypted in ECB mode
func isECBMode(ciphertext string) bool {
	return set1.DetectAES128_ECB(ciphertext)
}

// Detect encryption algorithm used to encrypt ciphertext block
func ECB_CBCDetectionOracle(plaintext string) string {
	_, ciphertext := encryptAES128_ECBOrCBC(plaintext)

	// Convert from base64 to hex since our ECB detection function expects a hex string
	ciphertext = hex.EncodeToString([]byte(utils.DecodeBase64(ciphertext)))

	mode := ""

	if isECBMode(ciphertext) {
		mode = "ECB"
	} else {
		mode = "CBC"
	}

	return mode
}