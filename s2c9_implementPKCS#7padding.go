/* Set 2 Challenge 9 - Implement PKCS#7 padding */

package main

import (
	"strings"
	"fmt"
)

func PKCS7padding(block string, blockSize int) string {
	lengthOfBlock := len(block)
	var paddedBlock, padValue string
	toPad := blockSize - lengthOfBlock

	if toPad != 0 {
		if toPad < 10 {
			padValue = "\\x0" + fmt.Sprintf("%v", toPad)
			paddedBlock = block + strings.Repeat(padValue, toPad)
		} else {
			padValue = "\\x" + fmt.Sprintf("%v", toPad)
			paddedBlock = block + strings.Repeat(padValue, toPad)
		}
	}

	// if lengthOfBlock is an integer multiple of blockSize, add an extra block
	// Confirms that last byte of the last block is not a pad byte
	if toPad == 0 {
		if blockSize < 10 {
			padValue = "\\x0" + fmt.Sprintf("%v", blockSize)
			paddedBlock = block + strings.Repeat(padValue, blockSize)
		} else {
			padValue = "\\x" + fmt.Sprintf("%v", blockSize)
			paddedBlock = block + strings.Repeat(padValue, blockSize)
		}
	}

	return  paddedBlock
}