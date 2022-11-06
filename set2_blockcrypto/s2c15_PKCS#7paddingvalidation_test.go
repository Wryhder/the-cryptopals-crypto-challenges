package set2_blockcrypto

import "testing"

func TestPKCS7PaddingValidation(t *testing.T) {
	// Test 1
	expected := "ICE ICE BABY"
	actual, err := PKCS7PaddingValidation("ICE ICE BABY\x04\x04\x04\x04")
	if err != nil {
		t.Fatalf("Expected %q but got error %q", expected, err)
	} else {
		if actual != expected {
		    t.Fatalf("actual %q, expected %q", actual, expected)
		}
	}

	// Test 2
	actual_2, err := PKCS7PaddingValidation("ICE ICE BABY\x05\x05\x05\x05")
    if err == nil {
		t.Fatalf("Expected an error %q but got %q", err, actual_2)
	}

	// Test 3
	actual_3, err := PKCS7PaddingValidation("ICE ICE BABY\x01\x02\x03\x04")
    if err == nil {
		t.Fatalf("Expected an error %q but got %q", err, actual_3)
	}
}