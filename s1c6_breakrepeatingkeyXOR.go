/* Set 1 Challenge 6 - Break repeating-key XOR */

package main

import (
	"fmt"
	"encoding/base64"
	"strings"
	"sort"
)

// var testStr1 string = "this is a test"
// var testStr2 string = "wokka wokka!!!"
var sampleFile = "./data/s1c6_encodedrepeatingkeyXORsample.txt"

// Decodes Base64 to string
func decodeBase64(str string) string {
	decodedStr, e := base64.StdEncoding.DecodeString(str)
    if e != nil {
        fmt.Println(e)
    }
    
	return string(decodedStr)
}

// Converts string to binary format
func stringToBin(str string) string {
	var binEquivalent string
    for _, char := range str {
        binEquivalent = fmt.Sprintf("%s%.8b",  binEquivalent, char)
    }
    return binEquivalent
}

// Computes Hamming distance between two strings (in binary format) of equal length
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

// Computes the sum of an array
func sum(array []float64) float64 {
	result := 0.0

	for _, value := range array {  
		result += value  
	}

	return result  
}

// Guess the length of the key used to encrypted text with repeating-key XOR
func guessKeySize(text string, lowerLimitGuessRange, upperLimitGuessRange int) int {
	var keySizesAndAverageNormalizedHds = make(map[int]float64)
	lengthOfText := len(text)
	
	/* For each keySize, find the hamming distance between every two adjacent blocks.
	Then normalize each distance and average the normalized distances. */
	for possibleKeySize := lowerLimitGuessRange; possibleKeySize <= upperLimitGuessRange; possibleKeySize++ {
		var normalizedHds []float64

		for count := 0; count < lengthOfText; count+=possibleKeySize {
			// This checks for the end of the text, when there isn't at least two blocks
			// of possibleKeySize-length left (cannot find hamming distance of less than
			// two full length blocks); Ignore any dangling bits.
			if len(text[count:]) <= possibleKeySize * 2 {
				break
			} else {
				// For example: block[0:8, 8:16], block[8:16, 16:24], block[16:24, 24:32],
				// block[24:32, 32:40], ..., until end of text (as long as the loop index
				// increments by possibleKeySize each time, such as `count+=possibleKeySize`)
				firstBlock := text[count:count+possibleKeySize]
				secondBlock := text[count+possibleKeySize:count+possibleKeySize*2]

				hammingDistance := HammingDistance(firstBlock, secondBlock)
				normalizedHd := hammingDistance / possibleKeySize
				normalizedHds = append(normalizedHds, float64(normalizedHd))
			}
		}

		averageNormalizedHd := sum(normalizedHds) / float64(len(normalizedHds))
		keySizesAndAverageNormalizedHds[possibleKeySize] = averageNormalizedHd
	}
	
	/* Find and return the best keySize */
	
	// The list of averageNormalizedHds (for all key sizes) will be sorted
	// to get the smallest value.
	var averageNormalizedHds []float64
	for _, averageNormalizedHd := range keySizesAndAverageNormalizedHds {
		averageNormalizedHds = append(averageNormalizedHds, averageNormalizedHd)
    }
	sort.Float64s(averageNormalizedHds[:])
	
	// Get all key size matching the minimum normalized hd
	var bestKeySize int
	for possibleKeySize, normalizedHd := range keySizesAndAverageNormalizedHds {
		if normalizedHd == averageNormalizedHds[0] {
			bestKeySize = possibleKeySize
		}

    }
	fmt.Println(keySizesAndAverageNormalizedHds)

	return bestKeySize
}

// Break ciphertext into blocks of keySize length
func breakTextIntoBlocks(decodedStr string, keySize int) []string {
	lengthOfdecodedStr := len(decodedStr)

	// Break text into blocks of keySize length
	var cipherTextBlocks []string

	for i := 0; i < lengthOfdecodedStr; i+=keySize {
		// This check prevents a possible `runtime error: slice bounds out of range`
		// which may occur when the final block is less than keySize
		if (lengthOfdecodedStr - i) < keySize {
			cipherTextBlocks = append(cipherTextBlocks, decodedStr[i:lengthOfdecodedStr])
		} else {
			cipherTextBlocks = append(cipherTextBlocks, decodedStr[i:i + keySize])
		}
	}

	return cipherTextBlocks
}

// Transpose ciphertext blocks: make a block that is the first byte of every block,
// and a block that is the second byte of every block, and so on
func transposeBlocks(cipherTextBlocks []string) map[int][]string {
	var transposedBlocks = make(map[int][]string)

	/* 
	Looping through cipherTextBlocks in the outer loop rather than the inner loop as in:
	```
	for currentByte := 0; currentByte < keySize; currentByte++ {
	 	for _, block := range cipherTextBlocks {
	```
	prevents a possible `slice bounds out of range` error that may occur when
	len(last block in cipherTextBlocks) is less than keySize (since the loop may continue on
	as if the last block were of the same length as the previous ones even when it isn't).

	Using len(block) instead in the inner loop below ensures the inner loop
	only runs for the exact length of the block, without expecting a specific number
	of bytes (keySize) and ignoring keySize altogether. 
	*/
	for _, block := range cipherTextBlocks {
		numOfBytesInBlock := len(block)

		for currentByte := 0; currentByte < numOfBytesInBlock; currentByte++ {
			transposedBlocks[currentByte] = append(transposedBlocks[currentByte], string(block[currentByte]))
		}	
	}

	return transposedBlocks
}

func BreakRepeatingKeyXOR(text string, lowerLimitGuessRange, upperLimitGuessRange int) string {
	decodedStr := decodeBase64(text)
	keySize := guessKeySize(decodedStr, lowerLimitGuessRange, upperLimitGuessRange)

	// Break text into blocks of keySize length
	cipherTextBlocks := breakTextIntoBlocks(decodedStr, keySize)
	
	// Transpose and attempt to solve ciphertext blocks
	transposedBlocks := transposeBlocks(cipherTextBlocks)
	var joinedBlockStr, result string

	// Solve each block in transposedBlocks as if it was single-character XOR
	for _, block := range transposedBlocks {
		joinedBlockStr = strings.Join(block[:], "")
		result = SingleByteXORCipher([]byte(joinedBlockStr))
		fmt.Println()
		fmt.Println(result)
	}

	return ""
}