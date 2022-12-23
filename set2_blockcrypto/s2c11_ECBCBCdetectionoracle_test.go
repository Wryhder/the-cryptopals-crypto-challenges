package set2_blockcrypto

import (
	"bytes"
	"crypto/aes"
	"strings"
	"testing"
)

func TestGenerateRandomAESKey(t *testing.T) {
	blockSize := aes.BlockSize

	type TestCase struct {
		input          []byte
		expectedLength int
	}

	// Key should be a byte slice, length = blockSize
	cases := []TestCase{
		{
			input:          generateRandomBytes(blockSize),
			expectedLength: blockSize,
		},
		{
			input:          generateRandomBytes(blockSize / 2),
			expectedLength: blockSize / 2,
		},
		{
			input:          generateRandomBytes(blockSize * 2),
			expectedLength: blockSize * 2,
		},
	}

	for _, c := range cases {
		actualLength := len(c.input)
		expectedLength := c.expectedLength

		if actualLength != expectedLength {
			t.Fatalf("Expected %q, got %q", expectedLength, actualLength)
		}
	}

	// Test that generated keys are random
	generatedKeys := [][]byte{cases[0].input} // initialize slice with one key for the loop below
	for i := 1; i < 20; i++ {
		generatedKeys = append(generatedKeys, generateRandomBytes(blockSize))

		currentKey := generatedKeys[i]
		previousKey := generatedKeys[i-1]

		if bytes.Compare(currentKey, previousKey) == 0 {
			t.Fatalf("Expected random, non-repeating keys, "+
				"got identical keys: %q and %q", currentKey, previousKey)
		}
	}
}

func TestPadStartAndEnd(t *testing.T) {
	plaintext := "The afternoon is as hot as a fireplace."

	for i := 0; i < 10; i++ {
		paddedPlaintext := padStartAndEnd(plaintext)
		lengthOfPaddedPlaintext := len(paddedPlaintext)
		padding := (lengthOfPaddedPlaintext - len(plaintext)) / 2

		// Test amount of padding added
		if padding < 5 || padding > 10 {
			t.Fatalf("Expected padding to be between 5 and 10 bytes, got %q", padding)
		}

		// Validate padding
		extractedPlaintext := paddedPlaintext[padding : lengthOfPaddedPlaintext-padding]
		if string(extractedPlaintext) != plaintext {
			t.Fatalf("Invalid padding. Expected %q after stripping padding, got %q",
				plaintext, extractedPlaintext)
		}
	}
}

func TestEncryptAES128_ECBOrCBC(t *testing.T) {
	plaintext := "The afternoon is as hot as a fireplace."
	modes := []string{""}

	for i := 0; i < 5; i++ {
		mode, _ := encryptAES128_ECBOrCBC(plaintext)
		modes = append(modes, mode)
	}

	modeCount := make(map[string]int)
	for _, mode := range modes {
		modeCount[mode]++
	}

	ECBCount := modeCount["ECB"]
	CBCCount := modeCount["CBC"]
	if ECBCount < 1 && CBCCount < 1 {
		t.Fatalf("Expected plaintext strings encrypted in ECB mode at least once,"+
			" got %q; Expected plaintext strings encrypted in CBC mode at least once,"+
			" got %q", ECBCount, CBCCount)
	}
}

func TestECB_CBCDetectionOracle(t *testing.T) {
	actual := ECB_CBCDetectionOracle(strings.Repeat("A", 64))
	expectedOption1 := "ECB"
	expectedOption2 := "CBC"

	if actual != expectedOption1 && actual != expectedOption2 {
		t.Fatalf("Expected either %q or %q, got %q", expectedOption1, expectedOption2, actual)
	}
}
