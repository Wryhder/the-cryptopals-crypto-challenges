/* Set 2 Challenge 16 - CBC bitflipping attacks */

package set2_blockcrypto

import (
	"crypto/aes"
	"encoding/base64"
	"net/url"
	"strings"
	"fmt"

	set1 "wryhder/cryptopals-crypto-challenges/set1_basics"
	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

var BLOCKSIZE = aes.BlockSize
var KEY = string(utils.GenerateRandomBytes(BLOCKSIZE))
var IV = utils.GenerateRandomBytes(blockSize)

func CBCBitflippingAttack_Encrypt(userdata string) string {
	toPrepend := "comment1=cooking%20MCs;userdata="
	toAppend := ";comment2=%20like%20a%20pound%20of%20bacon"

	plaintext := toPrepend + url.QueryEscape(userdata) + toAppend
	ciphertext := EncryptAES128_CBC(plaintext, key, IV)

	return ciphertext
}

// Detects the block size of a cipher
func detectCipherBlockSize_CBCBitflippingAttack() int {
	var ciphertextLengths []int
	var blockSize int

	for n := 0; n < 64; n++ {
		plaintext := strings.Repeat("A", n)
		ciphertext := CBCBitflippingAttack_Encrypt(plaintext)
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

func CBCBitflippingAttack_Decrypt(ciphertext string) (string, error) {
	padded := DecryptAES128_CBC(ciphertext, key, IV)
	plaintext, err := PKCS7PaddingValidation(padded)
    if err != nil {
		return "", err
	}

	return plaintext, nil
}

func CBCBitflippingAttack_IsAdmin(encryptedUserdata string) bool {
	var isAdmin bool

	decrypted, err := CBCBitflippingAttack_Decrypt(encryptedUserdata)
	fmt.Println("decrypted: ", decrypted)
	if err == nil {
		isAdmin = strings.Contains(decrypted, ";admin=true;")
	}

	return isAdmin
}

/* 
https://github.com/FrugalGuy/bitflipper
https://crypto.stackexchange.com/questions/66085/bit-flipping-attack-on-cbc-mode
https://zhangzeyu2001.medium.com/attacking-cbc-mode-bit-flipping-7e0a1c185511
*/
func CBCBitflippingAttack(targetByte, desiredByte byte, targetBytePos int, 
	encryptedUserdata string) string {
	blockSize := detectCipherBlockSize_CBCBitflippingAttack()
	decoded, _ := base64.StdEncoding.DecodeString(encryptedUserdata)
	
	targetOffset := targetBytePos - blockSize
	maskByte := targetByte ^ desiredByte
	xored := maskByte ^ decoded[targetOffset]
	
	// Replace relevant byte and reconstruct ciphertext
	tamperedCiphertext := append(decoded[0:targetOffset], xored)
	tamperedCiphertext = append(tamperedCiphertext, decoded[targetOffset + 1:]...)

	return set1.ByteToBase64(tamperedCiphertext)
}