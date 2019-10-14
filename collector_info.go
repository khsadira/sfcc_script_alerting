package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getSfccInfo(host string, site string, token string) []byte {
	jsonFile, err := getInfoJson(host, site, token)
	if err != nil {
		log.Println(err)
		return nil
	}
	return jsonFile
}

func getInfoJson(host string, site string, token string) ([]byte, error) {
	client := &http.Client{}
	jsBody := fmt.Sprintf(`{"query":{"text_query":{"fields":["%s"],"search_phrase":"%s"}},"select": "(**)","count":1}`, "order_no", "100485506")
	jsonBody := []byte(jsBody)
	query := fmt.Sprintf("%s/s/%s/dw/shop/v19_8/order_search", host, site)
	req, err := http.NewRequest("POST", query, bytes.NewBuffer(jsonBody))
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}
	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	return buf, nil
}