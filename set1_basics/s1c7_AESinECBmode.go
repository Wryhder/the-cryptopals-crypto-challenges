/* Set 1 Challenge 7 - AES in ECB mode */

package set1_basics

import (
	"crypto/aes"
	"fmt"
	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

// https://stackoverflow.com/questions/24072026/golang-aes-ecb-encryption
// https://github.com/golang/go/issues/5597

func EncryptAES128_ECB(plaintext, key string) string {
	lengthOfPlaintext := len(plaintext)
	blockSize := aes.BlockSize

	if lengthOfPlaintext % blockSize != 0 {
		panic("Length of plaintext must be a multiple of block size")
	}

	block, _ := aes.NewCipher([]byte(key))
	ciphertext := make([]byte, lengthOfPlaintext)

	for begin, end := 0, blockSize; begin < lengthOfPlaintext;
		begin, end = begin+blockSize, end+blockSize {
		block.Encrypt(ciphertext[begin:end], []byte(plaintext[begin:end]))
	}

	encoded := ByteToBase64(ciphertext)
    return encoded
}

func DecryptAES128_ECB(ciphertext, key string) string {
	// Comment out this line when calling this function from other functions
	// that already decoded the ciphertext (such as DecryptAES128_CBC())
	ciphertext = utils.DecodeBase64(ciphertext)
	lengthOfCiphertext := len(ciphertext)
	
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
        fmt.Println("Unable to create new cipher.Block: ", err)
    }
	
	blockSize := aes.BlockSize
	plaintext := make([]byte, lengthOfCiphertext)

	for begin, end := 0, blockSize; begin < lengthOfCiphertext;
		begin, end = begin+blockSize, end+blockSize {
		block.Decrypt(plaintext[begin:end], []byte(ciphertext[begin:end]))
	}

    return string(plaintext)
}