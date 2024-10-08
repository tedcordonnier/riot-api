package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type LeagueInfo struct {
	LeagueID     string `json:"leagueId"`
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	SummonerID   string `json:"summonerId"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

func LEAGUE(url string) LeagueInfo {
	var leagueInfo []LeagueInfo
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	err = json.Unmarshal(body, &leagueInfo)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	rankedInfo := leagueInfo[0]
	return rankedInfo
}
