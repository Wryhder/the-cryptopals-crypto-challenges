/* Set 2 Challenge 15 - PKCS#7 padding validation */

package set2_blockcrypto

import "fmt"

func PKCS7PaddingValidation(plaintext string) (string, error) {
	lengthOfText := len(plaintext)
	textInBytes := []byte(plaintext)

	lastByte := textInBytes[lengthOfText - 1]
	paddingStart := lengthOfText - int(lastByte)
	padBytes := textInBytes[paddingStart:]
	for _, b := range padBytes {
		if b != lastByte {
			return "", fmt.Errorf("invalid padding")
		}
	}

	originalText := plaintext[0:paddingStart]
	return originalText, nil
}