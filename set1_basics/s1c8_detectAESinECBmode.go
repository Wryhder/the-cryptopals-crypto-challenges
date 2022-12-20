/* Set 1 Challenge 8 - Detect AES in ECB mode */

package set1_basics

import (
	"bufio"
	"os"
	"bytes"
	"strconv"

	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

func DetectAES128_ECB_File(filePath string) string {
	f, _ := os.Open(filePath)
	defer f.Close()

	var allResults = make(map[int]map[string]string)
	count := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		isECB := DetectAES128_ECB(line)

		allResults[count] = map[string]string{
			"isECB": strconv.FormatBool(isECB),
			"ciphertext": line,
		}

		count += 1
	}

	// Find line encrypted in ECB mode
	var ECBEncryted string

	for _, lineDetails := range allResults {
		mode, _ := strconv.ParseBool(lineDetails["isECB"])
		if mode {
			ECBEncryted = lineDetails["ciphertext"]
			break
		}
	}

	return ECBEncryted
}

func DetectAES128_ECB(ciphertext string) (bool) {
	var repeatedBlocks = make(map[int]int)

	cipherTextBlocks := utils.ChunkifyText(HexToByte(ciphertext), 16)
	for index, byteBlock := range cipherTextBlocks {
		for n := 0; n < len(cipherTextBlocks); n++ {

			// Don't compare a block with itself
			if index == n {
				continue
			}

			if bytes.Compare(byteBlock, cipherTextBlocks[n]) == 0 {
				repeatedBlocks[index] += 1
			}
		}
	}

	// Find any repeated block(s)
	isECB := false
	for _, freq := range repeatedBlocks {
		if freq >= 1 {
			isECB = true
		}
	}
	
	return isECB
}