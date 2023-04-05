/* Set 2 Challenge 12 - Byte-at-a-time ECB decryption (Simple) */

package set2_blockcrypto

import (
	"crypto/aes"
	"encoding/hex"
	"encoding/base64"
	"fmt"
	"bytes"
	"strings"

	set1 "wryhder/cryptopals-crypto-challenges/set1_basics"
	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

var ENCRYPTION_KEY []byte = utils.GenerateRandomBytes(aes.BlockSize)
var lengthOfCiphertext int
var endOfCipherTextReached bool

// Encrypts buffers under ECB mode using a consistent but unknown key
// on each function run
func EncryptAES128_ECB_SingleKey(plaintext string) string {
	toAppend := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
		"aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
		"dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
		"YnkK"

	plaintext += utils.DecodeBase64(toAppend)

	// Add PKCS7padding here; our ECB encrypting function does not pad plaintext
	// and panics if provided with inadequate input (less than multiple of block size)
	paddedPlaintext := string(PKCS7padding([]byte(plaintext), aes.BlockSize))
	ciphertext := set1.EncryptAES128_ECB(paddedPlaintext, string(ENCRYPTION_KEY))

	return ciphertext
}

// Detects the block size of a cipher
func detectCipherBlockSize() int {
	var ciphertextLengths []int
	var blockSize int

	for n := 0; n < 64; n++ {
		plaintext := strings.Repeat("A", n)
		ciphertext := EncryptAES128_ECB_SingleKey(plaintext)
		decoded := utils.DecodeBase64(ciphertext)

		ciphertextLengths = append(ciphertextLengths, (len(decoded)))

		if len(ciphertextLengths) > 1 &&
			ciphertextLengths[n] != ciphertextLengths[n-1] {
			// current length minus previous length
			blockSize = ciphertextLengths[n] - ciphertextLengths[n-1]
			break
		}
	}
	return blockSize
}

// Detects whether or not encryption is AES-ECB
func detectEncryptionMode(blockSize int) (string, error) {
	ciphertext := EncryptAES128_ECB_SingleKey(strings.Repeat("A", blockSize*4))
	// Convert from base64 to hex since our ECB detection function expects a hex string
	ciphertext = hex.EncodeToString([]byte(utils.DecodeBase64(ciphertext)))
	mode := ""

	if isECBMode(ciphertext) {
		mode = "ECB"
	} else {
		return "", fmt.Errorf("error detecting encryption mode")
	}

	return mode, nil
}

// Generates a dictionary of every possible last byte by feeding different
// strings to the oracle
func generateCiphertextsToMatchAgainst(n int, decrypted []string) map[string][]byte {
	var ciphertextsToMatchAgainst = make(map[string][]byte)

	for key := 0; key <= 255; key++ {
		plaintext := strings.Repeat("A", n) + strings.Join(decrypted, "") +
			string(uint8(key))
		ciphertext := EncryptAES128_ECB_SingleKey(plaintext)
		// not using utility function because we want to work with raw bytes
		// directly to avoid encoding-related errors
		decoded, _ := base64.StdEncoding.DecodeString(ciphertext)
		ciphertextsToMatchAgainst[plaintext] = decoded
	}

	return ciphertextsToMatchAgainst
}

// Matches actual output of cipher against generated ciphertexts in a bid to decrypt
// each byte of an unknown string appended to a plaintext by the cipher 
func matchOutputToGeneratedCiphertexts(n, start, end int,
	ciphertextsToMatchAgainst map[string][]byte) (matched string) {

	plaintext := strings.Repeat("A", n)
	ciphertext := EncryptAES128_ECB_SingleKey(plaintext)

	// This piece of code is pretty much useless, except that I need the
	// endOfCipherTextReached check to get the loop in ByteatatimeECBdecryption_Simple
	// to run at all.
	lengthOfCiphertext = len(ciphertext)
	if end >= lengthOfCiphertext {
		endOfCipherTextReached = true
	}

	decoded, _ := base64.StdEncoding.DecodeString(ciphertext)
	for key, value := range ciphertextsToMatchAgainst {
		if bytes.Equal(value[start:end], decoded[start:end]) {
			matched = key
		}
	}

	return
}

// Decrypts each byte of an unknown string appended to a plaintext by a cipher (AES-ECB)
func ByteatatimeECBdecryption_Simple() string {
	var decryptedBytes []string
	blockSize := detectCipherBlockSize()
	
	_, err := detectEncryptionMode(blockSize)
	if err == nil {
		out:
		// The ciphertext is not directly accessible within this scope so I'm relying on 
		// global variables (endOfCipherTextReached) to get this loop to run.
		for start, end := 0, blockSize; !endOfCipherTextReached;
			start, end = start + blockSize, end + blockSize {
			for n := blockSize - 1; n >= 0; n-- {
				ciphertextsToMatchAgainst := 
					generateCiphertextsToMatchAgainst(n, decryptedBytes)
				matched := matchOutputToGeneratedCiphertexts(n, start, end,
					ciphertextsToMatchAgainst)
				// It's assumed we've hit the padding
				if (len(matched) == 0) {
					break out
				}

				discoveredByte := matched[len(matched)-1:]
				decryptedBytes = append(decryptedBytes, discoveredByte)
			}
		}
	}

	// decryptedBytes will include one padding character which we need to chop off
	return strings.Join(decryptedBytes[:len(decryptedBytes) - 1], "")
}
