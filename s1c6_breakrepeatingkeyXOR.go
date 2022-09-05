/* Set 1 Challenge 6 - Break repeating-key XOR */

package main

import (
	"fmt"
)

func stringToBin(str string) string {
	var binEquivalent string
    for _, char := range str {
        binEquivalent = fmt.Sprintf("%s%.8b",  binEquivalent, char)
    }
    return binEquivalent
}

func HammingDistance(str1, str2 string) int {
	str1Bin := stringToBin(str1)
	str2Bin := stringToBin(str2)
	XORResult := make([]rune, len(str1Bin))

	for i := 0; i < len(str1Bin); i++ {
		XORResult[i] = []rune(str1Bin)[i] ^ []rune(str2Bin)[i]
	}

	var distance int = 0
	for i := 0; i < len(XORResult); i++ {
		if XORResult[i] == 1 {
			distance += 1
		}
	}

	return distance
}
