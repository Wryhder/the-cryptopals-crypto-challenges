package set2_blockcrypto

import (
	"crypto/aes"
	"fmt"
	"testing"
)

func TestGenerateRandomAESKey(t *testing.T) {
	key := generateRandomBytes(aes.BlockSize)
	fmt.Println(key)

	// Key should be a byte slice, length = blockSize
	// Multiple keys should not be equal (20 keys to test?)
}

func TestPadStartAndEnd(t *testing.T) {
	plaintext := ""
	paddedPlaintext := padStartAndEnd(plaintext)
	fmt.Println(paddedPlaintext)
}

func TestEncryptAES128_ECBOrCBC(t *testing.T) {
	testCases := []string{
		"We may note at the outset the spirit of pessimism which, like the curse on the hoard, pervades the whole.",
		"Animals hoard food so they will have enough to get them through harsh winters.",
		"I am from only discovering that secret hoard of love when it was too late to listen.",
		"Many collectors hoard packages of the cards in the hopes that the next Mickey Mantle will come along and prove just as valuable.",
		"Interesting piece on a viking hoard discovered on a farm in Sweden, which includes many Arabic coins.",
		"However, when the Great Depression hit, President Franklin Roosevelt made it against the law to hoard or be in possession of gold.",
	}

	modes := []string{""}
	for _, c := range testCases {
		fmt.Println(c)
		_, encrypted := encryptAES128_ECBOrCBC("")
		mode := ECB_CBCDetectionOracle(encrypted)
		modes = append(modes, mode)
	}

	modeCount := make(map[string]int)
	for _, mode := range modes {
		modeCount[mode]++
	}

	ECBCount := modeCount["ECB"]
	CBCCount := modeCount["CBC"]
	if ECBCount < 1 && CBCCount < 1 {
			t.Fatalf("Expected plaintext strings encrypted in ECB mode at least once," +
			" got %q; Expected plaintext strings encrypted in CBC mode at least once," +
			" got %q", ECBCount, CBCCount)
		}
}

type TestCase struct {
	input    string
	expected string
}

func TestECB_CBCDetectionOracle(t *testing.T) {
	cases := []TestCase{
		{
			input:    "",
			expected: "",
		},
		{
			input:    "1",
			expected: "1",
		},
		{
			input:    "2",
			expected: "2",
		},
		{
			input:    "1",
			expected: "1",
		},
		{
			input:    "2",
			expected: "2",
		},
	}

	for _, c := range cases {
		actual := ECB_CBCDetectionOracle(c.input)
		expected := c.expected

		if actual != expected {
			t.Fatalf("Expected %q, got %q", expected, actual)
		}
	}
}