package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func MATCH(url string) []string {
	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the JSON array into a Go slice
	var matches []string
	err = json.Unmarshal(body, &matches)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return matches
}
