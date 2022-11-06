package set1_basics

import (
	"testing"
	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

func TestBreakRepeatingKeyXOR(t *testing.T) {
	sampleFile := "../data/s1c6_encodedrepeatingkeyXORsample.txt"

	actual := BreakRepeatingKeyXOR((utils.ReadTextFile(sampleFile)), 2, 40)
    expected := "Terminator X: Bring the noise"

    if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}