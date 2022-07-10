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

	
	// Set 1 Challenge 2
	
	buffer1 := "1c0111001f010100061a024b53535009181c"
	buffer2 := "686974207468652062756c6c277320657965"

	// expectedXORCombination := "746865206b696420646f6e277420706c6179"
	actualXORCombination := fixedXOR(hexToByte(buffer1), hexToByte(buffer2))
	
	fmt.Printf("%x\n", string(actualXORCombination))
	


}