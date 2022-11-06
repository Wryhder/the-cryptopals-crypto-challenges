package set2_blockcrypto

import (
	"testing"
	"bytes"
)

func TestPKCS7padding(t *testing.T) {
	// Test 1
	actual_1 := PKCS7padding([]byte("YELLOW SUBMARINE"), 20)
    expected_1 := []byte{89, 69, 76, 76, 79, 87, 32, 83, 85, 66, 77, 65, 82, 73, 78, 69, 4, 4, 4, 4,}

	if !bytes.Equal(actual_1, expected_1) {
		t.Errorf("actual %q, expected %q", actual_1, expected_1)
	}

	// Test 2
	actual_2 := string(PKCS7padding([]byte("YELLOW SUBMARINE"), 20))
    expected_2 := "YELLOW SUBMARINE\x04\x04\x04\x04"

	if actual_2 != expected_2 {
        t.Errorf("actual %q, expected %q", actual_2, expected_2)
    }

	// Test 3
	actual_3 := PKCS7padding([]byte("YELLOW SUBMARINE"), 16)
    expected_3 := []byte{89, 69, 76, 76, 79, 87, 32, 83, 85, 66, 77, 65, 82, 73, 78, 69, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16}

	if !bytes.Equal(actual_3, expected_3) {
		t.Errorf("actual %q, expected %q", actual_3, expected_3)
	}

	// Test 4
	actual_4 := PKCS7padding([]byte("YELLOW S"), 8)
    expected_4 := []byte{89, 69, 76, 76, 79, 87, 32, 83, 8, 8, 8, 8, 8, 8, 8, 8}

	if !bytes.Equal(actual_4, expected_4) {
		t.Errorf("actual %q, expected %q", actual_4, expected_4)
	}

	// Test 5
	actual_5 := PKCS7padding([]byte("YELL"), 8)
    expected_5 := []byte{89, 69, 76, 76, 4, 4, 4, 4}

	if !bytes.Equal(actual_5, expected_5) {
		t.Errorf("actual %q, expected %q", actual_5, expected_5)
	}

	// Test 6
	actual_6 := PKCS7padding([]byte("YELL0"), 8)
    expected_6 := []byte{89, 69, 76, 76, 48, 3, 3, 3}

	if !bytes.Equal(actual_6, expected_6) {
		t.Errorf("actual %q, expected %q", actual_6, expected_6)
	}
}