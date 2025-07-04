package logic

import (
	"encoding/base64"
	"log"
	"os"
)

func EncodeBase64(pathFile string) {

	file_byte, err := os.ReadFile(pathFile) // read the file
	if err != nil {
		log.Printf("Failed to read the file %s and error are %v", pathFile, err)
	}

	encodePDF := base64.StdEncoding.EncodeToString(file_byte) // encode the file byte into base64

	file, err := os.Create("base.txt") //creating a  file
	if err != nil {
		log.Printf("Failed to create a file %v", err)
	}

	len_string, err := file.WriteString(encodePDF)
	if err != nil {
		log.Println("Failed to write the string content in file ", err)
	}
	defer file.Close()
	log.Println("Total length of string written are ", len_string)
}
