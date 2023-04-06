/* Set 2 Challenge 13 - ECB cut-and-paste */

package set2_blockcrypto

import (
	"fmt"
	"strings"
    "net/mail"
	"crypto/aes"
	"encoding/base64"

	set1 "wryhder/cryptopals-crypto-challenges/set1_basics"
	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

var blockSize = aes.BlockSize
var key = string(utils.GenerateRandomBytes(blockSize))

type User struct {
  email      string
  uid    int
  role   string
}

func parse(str string) map[string]string {
    keyValuePairs := strings.Split(str, "&")
    var mapObj = make(map[string]string)

    for _, pair := range keyValuePairs {
        keyValPair := strings.Split(pair, "=")
        mapObj[keyValPair[0]] = keyValPair[1]
    }
    
    return mapObj
}

func cleanUserInput(userInput string) string {
    findMetachars := func(c rune) bool {
        return c == '&' || c == '='
    }

    splitStr := strings.FieldsFunc(userInput, findMetachars)
    return splitStr[0]
}

func validateMail(address string) (string, bool) {
    addr, err := mail.ParseAddress(address)
    if err != nil {
        return "", false
    }
    return addr.Address, true
}

var userID int

func encodeUserProfile(userInput string) string {
    cleanedInput := cleanUserInput(userInput)
    var encoded string

    if email, valid := validateMail(cleanedInput); valid {
        userID += 1

        profile := User{
            email: email,
            uid: userID,
            role: "user",
        }

        encoded = "email=" + profile.email + 
        "&uid=" + fmt.Sprint(profile.uid) + 
        "&role=user"
    }
    
    return encoded
}

func encryptUserProfile(encodedProfile string) string {
    // Add PKCS7padding here; our ECB encrypting function does not pad plaintext
	// and panics if provided with inadequate input (less than multiple of block size)
	paddedPlaintext := string(PKCS7padding([]byte(encodedProfile), aes.BlockSize))
	ciphertext := set1.EncryptAES128_ECB(paddedPlaintext, string(key))

	return ciphertext
}

func decryptAndParseUserProfile(encryptedProfile string) (map[string]string, error) {
    padded := set1.DecryptAES128_ECB(encryptedProfile, key)
    plaintext, err := PKCS7PaddingValidation(padded)
    if err != nil {
		return nil, err
	}

	return parse(plaintext), nil
}

/* 
Modifies ciphertexts to create an admin role.
- Creates three profiles using tweaked emails such that we have the
    following approx. block structures, 16-bytes each:
    "email=foobar@zaz .com&uid=1&role= user............"
    "email=foobar@co. admin&uid=1&role =user..........."
    "email=foo@bar.co  &uid=1&role=user {last blocks including padding}"
- The produced ciphertexts are combined based on the following arrangement
    to form one final 'valid' ciphertext:
    "email=foobar@zaz .com&uid=1&role= admin&uid=1&role email=foobar@zaz 
        {final blocks including padding}"
*/
func ECBCutAndPaste() map[string]string {
    var tamperedCiphertext []byte

    profile1 := encodeUserProfile("foobar@zaz.com")
    ciphertext1 := encryptUserProfile(profile1)
    base64_decoded1, _ := base64.StdEncoding.DecodeString(ciphertext1)
    firstBlock := base64_decoded1[:blockSize]
    firstAndSecondBlocks := base64_decoded1[:blockSize*2]

    profile2 := encodeUserProfile("foobar@co.admin")
    ciphertext2 := encryptUserProfile(profile2)
    base64_decoded2, _ := base64.StdEncoding.DecodeString(ciphertext2)
    secondBlock := base64_decoded2[blockSize:blockSize*2]

    profile3 := encodeUserProfile("foo@bar.co")
    ciphertext3 := encryptUserProfile(profile3)
    base64_decoded3, _ := base64.StdEncoding.DecodeString(ciphertext3)
    finalBlocks := base64_decoded3[blockSize*2:]

    tamperedCiphertext = append(firstAndSecondBlocks, secondBlock...)
    tamperedCiphertext = append(tamperedCiphertext, firstBlock...)
    tamperedCiphertext = append(tamperedCiphertext, finalBlocks...)

    base64_encoded := set1.ByteToBase64(tamperedCiphertext)
    decryptedProfile, _ := decryptAndParseUserProfile(base64_encoded)

    return decryptedProfile
}

func ECBCutAndPaste_IsAdmin() bool {
	decryptedProfile := ECBCutAndPaste()

    role := decryptedProfile["role"]
    isAdmin := role == "admin"

	return isAdmin
}