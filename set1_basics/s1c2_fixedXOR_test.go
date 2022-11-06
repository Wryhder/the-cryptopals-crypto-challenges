package set1_basics

import (
	"testing"
	"encoding/hex"
)

func TestFixedXOR(t *testing.T) {
	buffer1 := "1c0111001f010100061a024b53535009181c"
	buffer2 := "686974207468652062756c6c277320657965"
	XORCombination := FixedXOR(HexToByte(buffer1), HexToByte(buffer2))

	actual := hex.EncodeToString(XORCombination)
    expected := "746865206b696420646f6e277420706c6179"

    if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}