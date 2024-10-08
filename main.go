package main

import (
	"fmt"
)

// Riot API key and region
const apiKey = "RGAPI-de86bf29-d19c-42c5-b2e0-a6c2d05b8595"

func main() {

	account := getUserInput()

	urlACCOUNT := fmt.Sprintf("https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s?api_key=%s", account, apiKey)
	accountInfo := ACCOUNT(urlACCOUNT)
	puuid := accountInfo.Puuid
	fmt.Println(puuid)

	urlMATCH := fmt.Sprintf("https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=20&api_key=%s", puuid, apiKey)
	matches := MATCH(urlMATCH)

	printMatches(matches)
}

func getUserInput() string {
	summonerName := "Duke"
	summonerTag := "NA9"

	return fmt.Sprintf("%s/%s", summonerName, summonerTag)
}

func printMatches(matches []string) {
	for _, match := range matches {
		fmt.Println(match)
	}
}
