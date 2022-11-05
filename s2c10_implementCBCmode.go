/* Set 2 Challenge 10 - Implement CBC mode */

package main

import (
	"bytes"
	"crypto/aes"
)

/*
Approach (encryption):
- Pad plaintext
- Chunkify plaintext
- XOR each block with previous (Fixed XOR), then encrypt (AES in ECB mode)
*/

// Break text into n-byte blocks, padding last block if necessary
func chunkifyText(paddedText []byte, blockSize int)  [][]byte {

	var textBlocks [][]byte
	for i := 0; i < len(paddedText); i+=blockSize {
		textBlocks = append(textBlocks, paddedText[i:i + blockSize])
	}

	return textBlocks
}

// XOR each block with previous, then encrypt
func EncryptAES128_CBC(plaintext, key string, IV []byte) string {
	blockSize := aes.BlockSize
	paddedPlaintext := PKCS7padding([]byte(plaintext), blockSize)
	plainTextBlocks := chunkifyText(paddedPlaintext, blockSize)

	var cipherTextBlocks [][]byte
	var XORed []byte
	var encrypted string
	for index, block := range plainTextBlocks {

		if index == 0 {
			XORed = fixedXOR(IV, block)
		} else {
			previousCiphertextBlock := cipherTextBlocks[index - 1]
			XORed = fixedXOR(previousCiphertextBlock, block)
		}

		encrypted = EncryptAES128_ECB(string(XORed), key)
		cipherTextBlocks = append(cipherTextBlocks, []byte(encrypted))
	}

	return string(bytes.Join(cipherTextBlocks, []byte("")))
}

// Decrypt each block before XORing with previous block to recover plaintext
func DecryptAES128_CBC(ciphertext, key string, IV []byte) string {
	decodedCiphertext := decodeBase64(ciphertext)
	cipherTextBlocks := chunkifyText([]byte(decodedCiphertext), aes.BlockSize)

	var plainTextBlocks [][]byte
	var XORed []byte
	var decrypted string
	for index, block := range cipherTextBlocks {
		decrypted = DecryptAES128_ECB(string(block), key)

		if index == 0 {
			XORed = fixedXOR(IV, []byte(decrypted))
		} else {
			previousCiphertextBlock := cipherTextBlocks[index - 1]
			XORed = fixedXOR(previousCiphertextBlock, []byte(decrypted))
			
		}
		
		plainTextBlocks = append(plainTextBlocks, XORed)
	}

	return string(bytes.Join(plainTextBlocks, []byte("")))
}
