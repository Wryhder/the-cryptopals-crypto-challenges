package set2_blockcrypto

import (
	"crypto/aes"
	"testing"
)

func TestEncryptAES128_ECB_SingleKey(t *testing.T) {
	plaintext := "The same plaintext encrypted twice should produce same ciphertext."
	ciphertext1 := EncryptAES128_ECB_SingleKey(plaintext)
	ciphertext2 := EncryptAES128_ECB_SingleKey(plaintext)

	if ciphertext1 != ciphertext2 {
        t.Errorf("actual %q, expected %q", ciphertext1, ciphertext2)
    }
}

func TestDetectCipherBlockSize(t *testing.T) {
	actual := detectCipherBlockSize()
	expected := aes.BlockSize

	if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}

func TestDetectEncryptionMode(t *testing.T) {
	actual, _ := detectEncryptionMode(aes.BlockSize)
	expected := "ECB"

	if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}

func TestByteatatimeECBdecryption_Simple(t *testing.T) {
	actual := ByteatatimeECBdecryption_Simple()
	expected := `Rollin' in my 5.0
With my rag-top down so my hair can blow
The girlies on standby waving just to say hi
Did you stop? No, I just drove by
`

	if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}
