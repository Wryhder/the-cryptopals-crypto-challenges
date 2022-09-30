/* Set 1 Challenge 6 - Break repeating-key XOR */

package main

import (
	"fmt"
	"encoding/base64"
	// "bytes"
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
func HammingDistance(str1Bin, str2Bin string) int {
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

// Find the minimum value in an integer slice
func min(slice []int) int {
	var minVal int

    if len(slice) == 0 {
        panic("Empty slice, darling. You don chop so?")
    }

	for i, e := range slice {
		if i == 0 || e < minVal {
			minVal = e
		}
	}

	return minVal
}

var BYTESIZE int = 8

// Guess keySize (length of the key) used to perform repeating key XOR
func guessKeySize(text string, lowerLimitGuessRange, upperLimitGuessRange int) int {
	var hd, normalizedHd int
	var listOfNormalizedHds [39]int
	var listOfKeySizeAndNormalizedHd = make(map[int]int)
	var keySize int
	
	var count int = 0
	for possibleKeySize := lowerLimitGuessRange; possibleKeySize <= upperLimitGuessRange; possibleKeySize++ {
		// For each keySize, take the first and second keySize worth of bytes, and find the edit distance between them. Normalize this result by dividing by keySize.
		hd = HammingDistance(stringToBin(text)[:((BYTESIZE * possibleKeySize) - 1)], stringToBin(text)[(BYTESIZE * possibleKeySize):(((BYTESIZE * possibleKeySize) * 2) - 1)])

		// Normalize calculated hamming distance by dividing by keySize
		normalizedHd = hd / possibleKeySize
		listOfNormalizedHds[count] = normalizedHd
		listOfKeySizeAndNormalizedHd[possibleKeySize] = normalizedHd

		count++
	}
	
	for possibleKeySize, val := range listOfKeySizeAndNormalizedHd {
		if val == min(listOfNormalizedHds[:]) {
			keySize = possibleKeySize
		}
    }

	return keySize
}

func BreakRepeatingKeyXOR(text string, lowerLimitGuessRange, upperLimitGuessRange int) int {
	decodedStr := decodeBase64(text)
	keySize := guessKeySize(decodedStr, lowerLimitGuessRange, upperLimitGuessRange)
	binStr := stringToBin(decodedStr)
	lengthOfBinString := len(binStr)
	blockSize := keySize * BYTESIZE

	// Break text into blocks of keySize length
	var cipherTextBlocks []string

	for i := 0; i < lengthOfBinString; i+=blockSize {
		// This check prevents `runtime error: slice bounds out of range`
		// which happens when the final block is less than blockSize
		// e.g, slice bounds out of range [:23040] with length 23008
		if (lengthOfBinString - i) < blockSize {
			cipherTextBlocks = append(cipherTextBlocks, binStr[i:lengthOfBinString])
		} else {
			cipherTextBlocks = append(cipherTextBlocks, binStr[i:i + blockSize])
		}
	}

	// Transpose the blocks: make a block that is the first byte of every block,
	// and a block that is the second byte of every block, and so on
	var transposedBlocks = make(map[int][]string)

	/* 
	Looping through cipherTextBlocks in the outer loop rather than the inner loop as in:
	```
	for currentByte := 0; currentByte < keySize; currentByte++ {
	 	for _, block := range cipherTextBlocks {
	```
	prevents the quite pesky `slice bounds out of range` error I kept running into when
	len(last block in cipherTextBlocks) is less than blockSize (since the loop continued on
	as if the last block were of the same length as the previous ones even when it wasn't).

	Using len(block) / BYTESIZE instead in the inner loop below ensures the inner loop
	only runs for the exact length of the block, without expecting a specific number
	of bytes (blockSize) and ignoring keySize altogether. 
	*/
	for _, block := range cipherTextBlocks {
		numOfBytesInBlock := len(block) / BYTESIZE

		for currentByte := 0; currentByte < numOfBytesInBlock; currentByte++ {
			if currentByte == 0 {
				// block[0:8*1]
				transposedBlocks[currentByte] = append(transposedBlocks[currentByte], block[currentByte:BYTESIZE])
			} else {
				// block[8*1:8*2], block[8*2:8*3], block[8*3:8*4], ..., block[8*4:8*n]
				transposedBlocks[currentByte] = append(transposedBlocks[currentByte], block[BYTESIZE*currentByte:BYTESIZE*(currentByte+1)])
			}
		}	
	}

	fmt.Println(transposedBlocks)

	return 0
}