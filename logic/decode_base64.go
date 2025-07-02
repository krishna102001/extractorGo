package logic

import (
	"encoding/base64"
	"log"
	"os"
)

func DecodeBase64(st string) {
	file_byte, err := base64.StdEncoding.DecodeString(st)
	if err != nil {
		log.Println("err filebyte", err)
	}

	file, _ := os.Create("sample.pdf")
	l, err := file.Write(file_byte)
	if err != nil {
		log.Println("write", err)
	}
	log.Println("len : ", l)
}
