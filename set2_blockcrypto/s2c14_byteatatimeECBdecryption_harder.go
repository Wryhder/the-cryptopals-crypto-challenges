/* Set 2 Challenge 14 - Byte-at-a-time ECB decryption (Harder) */

package set2_blockcrypto

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"bytes"
	"strings"

	set1 "wryhder/cryptopals-crypto-challenges/set1_basics"
	utils "wryhder/cryptopals-crypto-challenges/utilities"
)

/* 
Tentative approach (a few things different in final solution):

- Detect cipher blocksize
- Detect total length of prepended bytes + unknown string (based on the point at which
	length of ciphertext changes; subtract length of own plaintext from old ciphertext length)
- Call oracle multiple times, comparing output blocks from each call to detect where blocks
    start to differ (to mark off the prepended bytes). Note this block.
- Call the oracle with increasingly longer plaintext until the ciphertext of the next block
    after the noted block (in previous step) matches the ciphertext of a block with input
	plaintext such as "AAAAAAAAAAAAAAAA" (exact length based on blocksize).
- Calculate the length of the prepended bytes based on the previous step (total length of all
	blocks up until the matched block with own plaintext, minus known length of own plaintext)
- Decrypt the unknown string using the same method from ByteatatimeECBdecryption_Simple()  
*/

// Generates a random count of random bytes
func generateRandCountBytes() (randBytes []byte, err error) {
	randCount := generateRandInt(1, 256)
	randBytes = make([]byte, randCount)
    _, err = rand.Read(randBytes)

    return
}

var BYTES_TO_PREPEND, _ = generateRandCountBytes()

// Encrypts buffers under ECB mode using a consistent but unknown key
// on each function run; prepends random bytes
func AES128_ECB_SingleKey_WithRandBytes(plaintext string) string {
	toAppend := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg" +
		"aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq" +
		"dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg" +
		"YnkK"

	plaintext = string(BYTES_TO_PREPEND) + plaintext + utils.DecodeBase64(toAppend)
	paddedPlaintext := string(PKCS7padding([]byte(plaintext), aes.BlockSize))

	ciphertext := set1.EncryptAES128_ECB(paddedPlaintext, string(ENCRYPTION_KEY))
	return ciphertext
}

// Detects the block size of a cipher
func detectCipherBlockSize_Harder() int {
	var ciphertextLengths []int
	var blockSize int

	for n := 0; n < 64; n++ {
		plaintext := strings.Repeat("A", n)
		ciphertext := AES128_ECB_SingleKey_WithRandBytes(plaintext)
		decoded := utils.DecodeBase64(ciphertext)

		ciphertextLengths = append(ciphertextLengths, (len(decoded)))

		if len(ciphertextLengths) > 1 &&
			ciphertextLengths[n] != ciphertextLengths[n-1] {
			// current length minus previous length
			blockSize = ciphertextLengths[n] - ciphertextLengths[n-1]
			break
		}
	}
	return blockSize
}

// Detects the total length of both unknown strings, prepended bytes and target string
// (useless, but I thought I'd need it)
func calcTotalLengthOfUnknownStrings() int {
	var ciphertextLengths []int
	var totalLengthOfUnknownStrings int
	var lengthOfPlaintext int

	for n := 0; n < 64; n++ {
		plaintext := strings.Repeat("A", n)
		ciphertext := AES128_ECB_SingleKey_WithRandBytes(plaintext)
		decoded := utils.DecodeBase64(ciphertext)

		ciphertextLengths = append(ciphertextLengths, (len(decoded)))

		if len(ciphertextLengths) > 1 &&
			ciphertextLengths[n] != ciphertextLengths[n-1] {
			
			previousLength := ciphertextLengths[n-1]
			lengthOfPlaintext = n
			totalLengthOfUnknownStrings = previousLength - lengthOfPlaintext
			break
		}
	}

	return totalLengthOfUnknownStrings
}

// Compares output blocks from multiple oracle calls to detect where blocks
// start to differ (to mark off the prepended bytes and target string)
func markOffPrependedRandBytes() int {
	blockSize := detectCipherBlockSize_Harder()
	var ciphertexts [][]byte
	// Block where prepended bytes end or target string starts (in cases where the length
	// of the prepended bytes is a multiple of the blocksize)
	var markOffPoint int

	// Generate ciphertexts from multiple calls to oracle
	for n := 0; n < 2; n++ {
		plaintext := strings.Repeat("A", n)
		ciphertext := AES128_ECB_SingleKey_WithRandBytes(plaintext)
		decoded, _ := base64.StdEncoding.DecodeString(ciphertext)

		ciphertexts = append(ciphertexts, decoded)
	}

	// Compare generated ciphertexts to find block where prepended bytes end
	ciphertext1 := utils.ChunkifyText(ciphertexts[0], blockSize)
	ciphertext2 := utils.ChunkifyText(ciphertexts[1], blockSize)
	for i := 0; i < len(ciphertext1); i++ {
		if !bytes.Equal(ciphertext1[i], ciphertext2[i]) {
			markOffPoint = i + 1    // add 1 because of zero indexing
			break
		}
	}
	return markOffPoint
}

// Generate sample start block ciphertext; next block after result of
// markOffPrependedRandBytes()
func generateSampleStartBlockCiphertext() []byte {
	endOfPrependedBytes := markOffPrependedRandBytes()

	// Assumes that as long as we know where the prepended random bytes end in
	// the ciphertext, we can use a plaintext at least thrice the length of the
	// blocksize to generate an entire block containing only our plaintext. The
	// generated ciphertext gives us something to match against.  
	plaintext := strings.Repeat("A", blockSize * 3)
	ciphertext := AES128_ECB_SingleKey_WithRandBytes(plaintext)

	decoded, _ := base64.StdEncoding.DecodeString(ciphertext)
	targetBlock := utils.ChunkifyText(decoded, blockSize)[endOfPrependedBytes + 1]
	sampleStartBlockCiphertext := targetBlock

	return sampleStartBlockCiphertext
}

// Find the optimal length of plaintext to give oracle to create a starting block
// for decryption
func findSweetSpot() int {
	blockSize := detectCipherBlockSize_Harder()
	endOfPrependedBytes := markOffPrependedRandBytes()
	var optimalLen int

	for n := 0; n < 64; n++ {
		plaintext := strings.Repeat("A", n)
		ciphertext := AES128_ECB_SingleKey_WithRandBytes(plaintext)

		decoded, _ := base64.StdEncoding.DecodeString(ciphertext)
		targetBlock := utils.ChunkifyText(decoded, blockSize)[endOfPrependedBytes + 1]
		sampleStartBlockCiphertext := generateSampleStartBlockCiphertext()
		if bytes.Equal(targetBlock, sampleStartBlockCiphertext) {
			optimalLen = n
			break
		}
	}

	return optimalLen
}

// Generates a dictionary of every possible last byte by feeding different
// strings to the oracle
func generateCiphertextsToMatchAgainst_Harder(n int, decrypted []string) map[string][]byte {
	var ciphertextsToMatchAgainst = make(map[string][]byte)

	for key := 0; key <= 255; key++ {
		plaintext := strings.Repeat("A", n) + strings.Join(decrypted, "") +
			string(uint8(key))
		ciphertext := AES128_ECB_SingleKey_WithRandBytes(plaintext)
		// not using utility function because we want to work with raw bytes
		// directly to avoid encoding-related errors
		decoded, _ := base64.StdEncoding.DecodeString(ciphertext)
		ciphertextsToMatchAgainst[plaintext] = decoded
	}

	return ciphertextsToMatchAgainst
}

// Matches actual output of cipher against generated ciphertexts in a bid to decrypt
// each byte of an unknown string appended to a plaintext by the cipher 
func matchOutputToGeneratedCiphertexts_Harder(n, start, end int,
	ciphertextsToMatchAgainst map[string][]byte) (matched string) {

	plaintext := strings.Repeat("A", n)
	ciphertext := AES128_ECB_SingleKey_WithRandBytes(plaintext)

	// This piece of code is pretty much useless, except that I need the
	// endOfCipherTextReached check to get the loop in ByteatatimeECBdecryption_Harder
	// to run at all.
	lengthOfCiphertext = len(ciphertext)
	if end >= lengthOfCiphertext {
		endOfCipherTextReached = true
	}

	decoded, _ := base64.StdEncoding.DecodeString(ciphertext)
	for key, value := range ciphertextsToMatchAgainst {
		if bytes.Equal(value[start:end], decoded[start:end]) {
			matched = key
		}
	}

	return
}

// Decrypts each byte of an unknown string appended to a plaintext by a cipher (AES-ECB)
func ByteatatimeECBdecryption_Harder() string {
	var decryptedBytes []string
	blockSize := detectCipherBlockSize_Harder()
	optimalLength := findSweetSpot()
	// markOffPrependedRandBytes() works block-wise whereas we need to traverse individual
	// characters of the ciphertext (not entirely sure why I need to add 1, but it works)
	startingPoint := (markOffPrependedRandBytes() + 1) * blockSize
	
	_, err := detectEncryptionMode(blockSize)
	if err == nil {
		out:
		for start, end := startingPoint, startingPoint + blockSize; !endOfCipherTextReached;
			start, end = start + blockSize, end + blockSize {
			for n := optimalLength - 1; n >= optimalLength - blockSize; n-- {
				ciphertextsToMatchAgainst := 
					generateCiphertextsToMatchAgainst_Harder(n, decryptedBytes)
				matched := matchOutputToGeneratedCiphertexts_Harder(n, start, end,
					ciphertextsToMatchAgainst)
				// It's assumed we've hit the padding
				if (len(matched) == 0) {
					break out
				}

				discoveredByte := matched[len(matched)-1:]
				decryptedBytes = append(decryptedBytes, discoveredByte)
			}
		}
	}

	// decryptedBytes will include one padding character which we need to chop off
	return strings.Join(decryptedBytes[:len(decryptedBytes) - 1], "")
}

