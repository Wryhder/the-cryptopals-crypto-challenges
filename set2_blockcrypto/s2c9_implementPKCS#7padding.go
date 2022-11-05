/* Set 2 Challenge 9 - Implement PKCS#7 padding */

package set2_blockcrypto

func PKCS7padding(ciphertext []byte, blockSize int) []byte {
	lengthOfCiphertext := len(ciphertext)
	var paddedCiphertext []byte = ciphertext
	toPad := lengthOfCiphertext % blockSize
	padValue := blockSize - toPad

	// if lengthOfBlock is an integer multiple of blockSize, add an extra block
	// Confirms that last byte of the last block is not a pad byte
	if toPad == 0 {
		padValue = blockSize
	}

	for n := 0; n < padValue; n++ {
		paddedCiphertext = append(paddedCiphertext, byte(padValue))
	}

	return  paddedCiphertext
}