package main

import (
	"fmt"
)

func main()  {
	/*
	Set 1 Challenge 1

	hexVal := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedBase64Equivalent := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	
	fmt.Println(hexToBase64(hexVal))
	*/

	
	// // Set 1 Challenge 2
	
	// buffer1 := "1c0111001f010100061a024b53535009181c"
	// buffer2 := "686974207468652062756c6c277320657965"

	// // expectedXORCombination := "746865206b696420646f6e277420706c6179"
	// actualXORCombination := fixedXOR(hexToByte(buffer1), hexToByte(buffer2))
	
	// fmt.Printf("%x\n", string(actualXORCombination))
	
	// // Set 1 Challenge 3
	// problem := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// fmt.Println(SingleByteXORCipher(hexToByte(problem)))

	// // Set 1 Challenge 4
	// fmt.Println(DetectSingleCharXOR("./data/s1c4_60charstrings.txt"))
	
// 	// Set 1 Challenge 5
// 	var openingStanza string = `Burning 'em, if you ain't quick and nimble
// I go crazy when I hear a cymbal`

	// // 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272
	// // a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f

	// fmt.Println(RepeatingKeyXOR(openingStanza, "ICE"))
	// fmt.Println(RepeatingKeyXOR(string(hexToByte()), "ICE"))

	// // Set 1 Challenge 6
	// fmt.Println(BreakRepeatingKeyXOR((readTextFile(sampleFile)), 2, 40))

	// // Set 1 Challenge 7
	// AESinECBFileContent := readTextFile("./data/s1c7_encodedAESinECBmodesample.txt")
	// fmt.Println(DecryptAES128_ECB(AESinECBFileContent, "YELLOW SUBMARINE"))

	// // Set 1 Challenge 8
	// fmt.Println(DetectAES128_ECB("./data/s1c8_encodedciphertextstrings.txt"))

	// Set 2 Challenge 9
	fmt.Println(PKCS7padding("YELLOW SUBMARINE", 20))
	// fmt.Println(PKCS7padding("YELLOW SUBMARINE", 16))
	// fmt.Println(PKCS7padding("YELLOW S", 8))
	// fmt.Println(PKCS7padding("YELL", 8))
	// fmt.Println(PKCS7padding("YELL0", 8))
}