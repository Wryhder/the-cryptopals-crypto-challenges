/* Set 1 Challenge 7 - AES in ECB mode */

package main

import (
	"crypto/aes"
	"fmt"
)

// https://stackoverflow.com/questions/24072026/golang-aes-ecb-encryption
// https://github.com/golang/go/issues/5597

func DecryptAES128_ECB(ciphertext, key string) string {
	decodedStr := decodeBase64(ciphertext)
	lengthOfdecodedStr := len(decodedStr)
	
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
        fmt.Println("Unable to create new cipher.Block: ", err)
    }
	
	blockSize := aes.BlockSize
	plaintext := make([]byte, lengthOfdecodedStr)

	for begin, end := 0, blockSize; begin < lengthOfdecodedStr;
		begin, end = begin+blockSize, end+blockSize {
		block.Decrypt(plaintext[begin:end], []byte(decodedStr[begin:end]))
	}

    return string(plaintext)
}