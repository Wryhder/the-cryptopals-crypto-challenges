/* Set 1 Challenge 8 - Detect AES in ECB mode */

package set1_basics

import (
	"bufio"
    "fmt"
    "os"
	"bytes"
)

// Break ciphertext into n-byte blocks (assuming padding is not necessary)
func breakLineIntoBlocks(ciphertext []byte, keySize int) [][]byte {
	lengthOfCiphertext := len(ciphertext)
	var cipherTextBlocks [][]byte

	for i := 0; i < lengthOfCiphertext; i+=keySize {
		cipherTextBlocks = append(cipherTextBlocks, ciphertext[i:i + keySize])
	}

	return cipherTextBlocks
}

func DetectAES128_ECB(filePath string) (int, string) {
	f, _ := os.Open(filePath)
	defer f.Close()

	var repeatedBlocks = make(map[int]int)
	var allResults = make(map[int]map[string]string)
	count := 0
	highestCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// for each line, split text into blocks of 16 bytes
		cipherTextBlocks := breakLineIntoBlocks(HexToByte(line), 16)

		for index, byteBlock := range cipherTextBlocks {
			for n := 0; n < len(cipherTextBlocks); n++ {

				// Don't compare a block with itself
				if index == n {
					continue
				}

				if bytes.Compare(byteBlock, cipherTextBlocks[n]) == 0 {
					repeatedBlocks[count] += 1
				}
			}
		}

		allResults[count] = map[string]string{
			"noOfRepeats": fmt.Sprintf("%v", repeatedBlocks[count]),
			"ciphertext": line,
		}

		currentCount := repeatedBlocks[count]
		if currentCount < highestCount {
			continue
		} else {
			highestCount = currentCount
		}

		count += 1
	}

	// Find line with the most number of repeated strings
	var result string
	for _, lineDetails := range allResults {
		if lineDetails["noOfRepeats"] == fmt.Sprintf("%v", highestCount) {
			result = lineDetails["ciphertext"]
			break
		}
	}
	
	return highestCount, result
}