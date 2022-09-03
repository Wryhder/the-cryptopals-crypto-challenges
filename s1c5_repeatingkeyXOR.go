/* Set 1 Challenge 5 - Implement repeating-key XOR */

package main

import (
	"strings"
	"encoding/hex"
)

// Sample string for testing
// var openingStanza string = `Burning 'em, if you ain't quick and nimble
// I go crazy when I hear a cymbal`

func RepeatingKeyXOR(plaintext, secretKey string) string {
	lengthOfPlaintext := len(plaintext)
	lengthOfSecretKey := len(secretKey)
	var encodedCipher string
	splitSecret := strings.Split(secretKey, "")

	// Unless the length of the plaintext is perfectly divisible by the length of the secret key,
	// there will be leftover characters. So, we repeat the secret (e.g "DUH") until the length
	// of the repeated secret (e.g "DUHDUH") is as long as or nearly as long as the length of the
	// plaintext. Any unaccounted characters are gotten and stored in `secondPartOfRepeatedKey`
	firstPartOfRepeatedKey := strings.Repeat(secretKey, lengthOfPlaintext/lengthOfSecretKey)
	secondPartOfRepeatedKey := strings.Join(splitSecret[:lengthOfPlaintext % lengthOfSecretKey], "")

	// The two parts of the repeated secret (if there were indeed leftover characters
	// as explained above) are combined and converted to bytes for the XOR operation
	repeatedKey := firstPartOfRepeatedKey + secondPartOfRepeatedKey

	// Each character in the plaintext is XORed with a byte of the key
	cipher := fixedXOR([]byte(plaintext), []byte(repeatedKey))
	encodedCipher = hex.EncodeToString(cipher)

	return encodedCipher
}