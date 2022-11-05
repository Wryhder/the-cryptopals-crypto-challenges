/* Set 1 Challenge 2 - Fixed XOR */

package set1_basics

func FixedXOR(buffer1, buffer2 []byte) []byte {
	n := len(buffer1)
	XORCombination := make([]byte, n)
	
	for i := 0; i < n; i++ {
		XORCombination[i] = buffer1[i] ^ buffer2[i]
	}

	return XORCombination
}