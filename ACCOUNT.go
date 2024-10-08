package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AccountInfo struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

func ACCOUNT(url string) AccountInfo {
	var accountInfo AccountInfo
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	err = json.Unmarshal(body, &accountInfo)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	return accountInfo
}
