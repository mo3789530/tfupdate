package utils

import (
	"log"
	"os"
)

func JsonOpen(filepath string) []byte {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Panic(err)
	}

	return data
}
