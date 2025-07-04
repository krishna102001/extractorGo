package logic

import (
	"encoding/base64"
	"log"
	"os"
)

func DecodeBase64(st string) {
	file_byte, err := base64.StdEncoding.DecodeString(st) // decoding the base64 string  into a file byte
	if err != nil {
		log.Printf("Error Occurred while decoding the string %v", err)
	}

	//creating a file
	file, err := os.Create("sample.pdf")
	if err != nil {
		log.Printf("Failed to create a file %v", err)
	}

	len_byte, err := file.Write(file_byte) // write the byte in file
	if err != nil {
		log.Println("Failed to write the byte content in file ", err)
	}

	file.Close()
	log.Println("Total length of byte written are ", len_byte)
}
