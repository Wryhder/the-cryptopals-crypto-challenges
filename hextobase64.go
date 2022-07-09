/* Set 1 Challenge 1 - Convert hex to base64 */

package main

import (
	"encoding/hex"
	"encoding/base64"
    "fmt"
)

func main() {
	hexVal := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	// expectedBase64Equivalent := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	fmt.Println(hexToBase64(hexVal))
}

func hexToBase64(hexVal string) string {
	byteEquivalent := hexToByte(hexVal)
	base64Equivalent := byteToBase64(byteEquivalent)

	return base64Equivalent
}

func hexToByte(hexVal string) []byte {
	byteEquivalent, err := hex.DecodeString(hexVal)
	if err != nil {
		panic(err)
	}

	return byteEquivalent
}

func byteToBase64(byteVal []byte) string {
	base64Equivalent := base64.StdEncoding.EncodeToString(byteVal)

	return base64Equivalent
}