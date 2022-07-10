/* Set 1 Challenge 2 - Fixed XOR */

package main

func fixedXOR(buffer1, buffer2 []byte) []byte {
	n := len(buffer1)
	XORCombination := make([]byte, n)
	
	for i := 0; i < n; i++ {
		XORCombination[i] = buffer1[i] ^ buffer2[i]
	}

	return XORCombination
}