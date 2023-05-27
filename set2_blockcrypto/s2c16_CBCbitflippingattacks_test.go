package set2_blockcrypto

import (
	"crypto/aes"
	"testing"
)

func TestCBCBitflippingAttack_Encrypt(t *testing.T) {
	encrypted := CBCBitflippingAttack_Encrypt("data;admin=true")
	isAdmin := CBCBitflippingAttack_IsAdmin(encrypted)
	if isAdmin {
		t.Fatalf("Expected user input to be escaped but user can " +
			"exploit input to elevate permissions (set admin=true)")
	}
}

func TestDetectCipherBlockSize_CBCBitflippingAttack(t *testing.T) {
	actual := detectCipherBlockSize_CBCBitflippingAttack()
	expected := aes.BlockSize

	if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}

func TestManipulateCiphertextToElevatePermissions(t *testing.T) {
	plaintext := "xdataAadminZtrue"
	encrypted := CBCBitflippingAttack_Encrypt(plaintext)
	tamperedCiphertext1 := CBCBitflippingAttack([]byte("A")[0], []byte(";")[0], 37,
	 encrypted)
	tamperedCiphertext2 := CBCBitflippingAttack([]byte("Z")[0], []byte("=")[0], 43,
	 tamperedCiphertext1)
	
	isAdmin := CBCBitflippingAttack_IsAdmin(tamperedCiphertext2)

	if !isAdmin {
		t.Fatalf("Expected elevated permissions (set admin=true) " +
			"but got admin=%t", isAdmin)
	}
}