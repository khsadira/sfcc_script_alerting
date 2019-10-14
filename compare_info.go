package main

import (
	"encoding/json"
	"log"
)

func compareInfo(data []byte, dataBfr []byte) (string, bool) {
	var jsonData Data
	var jsonDataBfr Data

	if data == nil || dataBfr == nil {
		return "nothing turned in data or dataBfr", false
	}

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Printf("Sfcc Data: %s", err)
		return "", false
	}
	err = json.Unmarshal(dataBfr, &jsonDataBfr)
	if err != nil {
		log.Printf("Local Data %s:", err)
		return "", false
	}

	sData := jsonData.Hits[0].Data.LastModified
	sDataBfr := jsonDataBfr.Hits[0].Data.LastModified

	if sData == sDataBfr {
		return "", true
	}
	return "last modified change", false
}