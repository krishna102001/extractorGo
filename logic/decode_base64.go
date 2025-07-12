package logic

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func DecodeBase64(st string) (string, error) {
	file_byte, err := base64.StdEncoding.DecodeString(st) // decoding the base64 string  into a file byte
	if err != nil {
		log.Printf("Error Occurred while decoding the string %v", err)
		return "", errors.New("failed to decode")
	}

	var out_pdf_file string = "out_pdf_file"
	_, err = os.Stat(out_pdf_file)
	if os.IsNotExist(err) {
		if err := os.Mkdir(out_pdf_file, 0755); err != nil {
			log.Printf("Failed to make the directory %v", err)
			return "", errors.New("failed to create directory")
		}
	}

	filePath := filepath.Join(out_pdf_file, fmt.Sprintf("sample_pdf_%v.pdf", time.Now()))

	//creating a file
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create a file %v", err)
		return "", errors.New("failed to create a file")
	}

	len_byte, err := file.Write(file_byte) // write the byte in file
	if err != nil {
		log.Println("Failed to write the byte content in file ", err)
		return "", errors.New("failed to write the byte content")
	}

	file.Close()
	log.Println("Total length of byte written are ", len_byte)
	return filePath, nil
}
