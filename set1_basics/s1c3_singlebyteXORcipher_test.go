package set1_basics

import "testing"
// 

func TestSingleByteXORCipher(t *testing.T){
	problem := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

    _, actual := SingleByteXORCipher(HexToByte(problem))
    expected := "Cooking MC's like a pound of bacon"

    if actual != expected {
        t.Errorf("actual %q, expected %q", actual, expected)
    }
}