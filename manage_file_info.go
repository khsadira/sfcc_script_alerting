package main

import (
	"io/ioutil"
	"log"
	"os"
)

func replaceFileInfo(filepath string, data []byte) {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Printf("failed opening file: %s", err)

	}
	defer file.Close()

	_, err = file.WriteAt(data, 0)

	if err != nil {
		log.Printf("failed writing to file: %s", err)
		return
	}

}

func getFileInfo(filepath string) []byte {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("File reading error %s", err)
		return nil
	}
	return data
}
