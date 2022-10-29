package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"encoding/base64"
	"log"
)

// Read image file
func getImageFromFilePath(filePath string) (image.Image, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    image, _, err := image.Decode(f)
    return image, err
}

// Read text file
func readTextFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Println("Unable to open file: ", err)
    }
	
    return string(content)
}

// Decodes Base64 to string
func decodeBase64(str string) string {
	decodedStr, e := base64.StdEncoding.DecodeString(str)
    if e != nil {
        fmt.Println(e)
    }
    
	return string(decodedStr)
}

