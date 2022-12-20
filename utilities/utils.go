package utilities

import (
	"os"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"encoding/base64"
	"log"
)

// Read image file
func GetImageFromFilePath(filePath string) (image.Image, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    image, _, err := image.Decode(f)
    return image, err
}

// Read text file
func ReadTextFile(filePath string) string {
	content, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("Unable to open file: ", err)
    }
	
    return string(content)
}

// Decodes Base64 to string
func DecodeBase64(str string) string {
	decodedStr, e := base64.StdEncoding.DecodeString(str)
    if e != nil {
        fmt.Println(e)
    }
    
	return string(decodedStr)
}

func AppendToFile(filePath, content string) {		
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	len, err := file.WriteString(content)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())
}
