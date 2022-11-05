/* Set 1 Challenge 1 - Convert hex to base64 */

package set1_basics

import (
	"encoding/hex"
	"encoding/base64"
)

func HexToBase64(hexVal string) string {
	byteEquivalent := HexToByte(hexVal)
	base64Equivalent := ByteToBase64(byteEquivalent)

	return base64Equivalent
}

func HexToByte(hexVal string) []byte {
	byteEquivalent, err := hex.DecodeString(hexVal)
	if err != nil {
		panic(err)
	}

	return byteEquivalent
}

func ByteToBase64(byteVal []byte) string {
	base64Equivalent := base64.StdEncoding.EncodeToString(byteVal)

	return base64Equivalent
}