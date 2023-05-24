package set2_blockcrypto

import (
	"bytes"
	"crypto/aes"
	"fmt"
	"testing"

	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

func TestGenerateRandCountBytes(t *testing.T) {
	randBytes1, _ := generateRandCountBytes()
	randBytes2, _ := generateRandCountBytes()

	if bytes.Equal(randBytes1, randBytes2) {
        t.Errorf("should not generate the same byte slice two times in a row")
    }
}

func TestAES128_ECB_SingleKey_WithRandBytes(t *testing.T) {
	plaintext := "The same plaintext encrypted twice should produce same ciphertext."
	ciphertext1 := AES128_ECB_SingleKey_WithRandBytes(plaintext)
	ciphertext2 := AES128_ECB_SingleKey_WithRandBytes(plaintext)

	if ciphertext1 != ciphertext2 {
        t.Errorf("actual %q, expected %q", ciphertext1, ciphertext2)
    }
}

func TestdetectCipherBlockSize_Harder(t *testing.T) {
	actual := detectCipherBlockSize_Harder()
	expected := aes.BlockSize

	if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}

func TestCalcTotalLengthOfUnknownStrings(t *testing.T) {
	prepended := BYTES_TO_PREPEND
	appended := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
		"aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
		"dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
		"YnkK"
	expected := len(string(prepended)) + len(utils.DecodeBase64(appended))
	actual := calcTotalLengthOfUnknownStrings()

	if actual != expected {
        t.Errorf("actual length of unknown strings %d, expected length of unknown strings %d",
		 actual, expected)
    }
}

func TestMarkOffPrependedRandBytes(t *testing.T) {
	prepended := BYTES_TO_PREPEND
	lenPrepended := len(string(prepended))
	fmt.Println(lenPrepended)
	blockSize := detectCipherBlockSize_Harder()

	endOfPrependedBytes := lenPrepended / blockSize
	if lenPrepended % blockSize > 0 {
		endOfPrependedBytes += 1
	}

	expected := endOfPrependedBytes
	actual := markOffPrependedRandBytes()

	if actual != expected {
        t.Errorf("actual %d, expected %d", actual, expected)
    }
}

func TestFindSweetSpot(t *testing.T) {
	optimalPlaintextLength := findSweetSpot()
	if optimalPlaintextLength <= 0 {
        t.Errorf("actual %d, expected optimal plaintext length to be greater than 0",
		 optimalPlaintextLength)
    }
}

func TestByteatatimeECBdecryption_Harder(t *testing.T) {
	actual := ByteatatimeECBdecryption_Harder()
	expected := `Rollin' in my 5.0
With my rag-top down so my hair can blow
The girlies on standby waving just to say hi
Did you stop? No, I just drove by
`

	if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}