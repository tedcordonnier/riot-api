package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type SummonerInfo struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

func SUMMONER(url string) SummonerInfo {
	var summonerInfo SummonerInfo
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	err = json.Unmarshal(body, &summonerInfo)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return summonerInfo
}
