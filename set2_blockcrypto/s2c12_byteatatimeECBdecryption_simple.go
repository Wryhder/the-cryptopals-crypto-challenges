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

// TODO:
func EncryptAES128_ECB_SingleKey(plaintext string) string {
	toAppend := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
		"aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
		"dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
		"YnkK"

	plaintext += utils.DecodeBase64(toAppend)
	fmt.Println("length of plaintext + unknown string: ", len(plaintext))

	// Add PKCS7padding here; our ECB encrypting function does not pad plaintext
	// and panics if provided with inadequate input (less than multiple of block size)
	paddedPlaintext := string(PKCS7padding([]byte(plaintext), aes.BlockSize))
	fmt.Println("length of plaintext + unknown string + padding: ", len(paddedPlaintext))
	fmt.Println()
	ciphertext := set1.EncryptAES128_ECB(paddedPlaintext, string(ENCRYPTION_KEY))

	return ciphertext
}

// TODO:
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

// TODO:
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

// TODO:
func generateCiphertextsToMatchAgainst(n int, discoveredStr []string) map[string][]byte {
	var ciphertextsToMatchAgainst = make(map[string][]byte)

	// TODO: Remove after testing
	fmt.Println("generating a map possible ciphertexts...")
	for key := 0; key <= 255; key++ {
		plaintext := strings.Repeat("A", n) + strings.Join(discoveredStr, "") +
			string(uint8(key))
		ciphertext := EncryptAES128_ECB_SingleKey(plaintext)
		decoded, _ := base64.StdEncoding.DecodeString(ciphertext)

		firstBlock := substr(string(decoded), 0, 16)
		ciphertextsToMatchAgainst[plaintext] = decoded
		// TODO: Remove after testing
		fmt.Println(plaintext, ": ", firstBlock)
	}

	return ciphertextsToMatchAgainst
}

// TODO:
func matchOutputToGeneratedCiphertexts(n int, discoveredStr []string,
	ciphertextsToMatchAgainst map[string][]byte) (matched string) {

	blockSize := detectCipherBlockSize()
	plaintext := strings.Repeat("A", n)
	// TODO: Remove after testing
	fmt.Println("plaintext: ", plaintext)
	ciphertext := EncryptAES128_ECB_SingleKey(plaintext)
	// TODO: Remove after testing
	fmt.Println("ciphertext: ", ciphertext)
	decoded, _ := base64.StdEncoding.DecodeString(ciphertext)
	// TODO: Remove after testing
	fmt.Println("decoded: ", decoded)
	
	for key, value := range ciphertextsToMatchAgainst {
		if bytes.Equal(value[:blockSize], decoded[:blockSize]) {
			matched = key
			// TODO: Remove after testing
			fmt.Println("matchedCiphertext", matched)
			fmt.Println()
		}
	}

	return
}

// TODO:
func ByteatatimeECBdecryption_Simple() string {
	fmt.Println("ENCRYPTION_KEY: ", ENCRYPTION_KEY)
	var bytesDiscovered []string

	blockSize := detectCipherBlockSize()
	fmt.Println("blockSize: ", blockSize)
	_, err := detectEncryptionMode(blockSize)
	if err == nil {
		for n := blockSize - 1; n >= 0; n-- {
			// TODO: Remove after testing
			fmt.Println("generating dictionary of ciphertexts to match against...")
			fmt.Println("bytesDiscovered: " + strings.Join(bytesDiscovered, ""))
			ciphertextsToMatchAgainst := 
				generateCiphertextsToMatchAgainst(n, bytesDiscovered)
			// TODO: Remove after testing
			fmt.Println()
			fmt.Println()
			fmt.Println("matching output to generated ciphertexts...")
			matched := matchOutputToGeneratedCiphertexts(n, bytesDiscovered,
				 ciphertextsToMatchAgainst)
			fmt.Println("ENCRYPTION_KEY: ", ENCRYPTION_KEY)
			fmt.Println("matched: ", matched)
			fmt.Println("length of matchedCiphertext: ", len(matched))
			discoveredByte := matched[len(matched)-1:]
			fmt.Println("discoveredByte: ", discoveredByte)
			bytesDiscovered = append(bytesDiscovered, discoveredByte)
			fmt.Println()
			// break
		}
	}

	
	fmt.Println("ENCRYPTION_KEY: ", ENCRYPTION_KEY)
	return "bytesDiscovered: " + strings.Join(bytesDiscovered, "")
}

// ----------------------------------------------------------------------------------

// https://stackoverflow.com/a/56129336
// TODO:
func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
