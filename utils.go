package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"image"
	_ "image/jpeg"
	_ "image/png"
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
        fmt.Println(err)
    }
	
    return string(content)
}
