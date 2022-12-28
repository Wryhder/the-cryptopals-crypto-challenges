package utilities

import (
	"os"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"encoding/base64"
	"log"
	"crypto/rand"
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

// Chunk text into n-byte blocks, where n can be key size or block size
// (assumes there are no dangling bits, such as in already-padded text)
func ChunkifyText(text []byte, size int)  [][]byte {
	var textBlocks [][]byte

	for i := 0; i < len(text); i+=size {
		textBlocks = append(textBlocks, text[i:i + size])
	}

	return textBlocks
}

// Generate random bytes of specified size/length;
// used to generate random AES keys or IVs
func GenerateRandomBytes(size int) []byte {
	randBytes := make([]byte, size)
	_, err := rand.Read(randBytes)
	if err != nil {
		fmt.Println("error: ", err)
		return nil
	}
	return randBytes
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
