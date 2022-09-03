/* Set 1 Challenge 4 - Detect single-character XOR */

package main

import (
	"bufio"
    "fmt"
    "os"
)

/* 
1. Read text file
2. Iterate through lines in files
3. Run strings through SingleByteXORCipher function
*/

// func DetectSingleCharXOR (filePath string) string {
func DetectSingleCharXOR (filePath string) map[string]string {
	f, _ := os.Open(filePath)
	var allResults = make(map[string]string)
	highestTextScore := 0.0

    // Create a new Scanner for the file
    scanner := bufio.NewScanner(f)

    // Loop over all lines in the file to find the string
	// that was encrypted by single-character XOR 
    for scanner.Scan() {
		line := scanner.Text()
		decryptedVal := SingleByteXORCipher(hexToByte(line))
		currentScore := textScorer(decryptedVal)

		fmt.Println(currentScore)

		allResults[fmt.Sprintf("%.5f", currentScore)] = decryptedVal
		
		if currentScore < highestTextScore {
			continue
		} else {
			highestTextScore = currentScore
		}
    }

	// return allResults[fmt.Sprintf("%.5f", highestTextScore)]
	return allResults
}