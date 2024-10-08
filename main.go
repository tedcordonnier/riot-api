package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type RankedInfo struct {
	Account string
	Tier    string
	Rank    string
	Wins    int
	Losses  int
}

var tmpl = template.Must(template.ParseFiles("index.html"))

// Riot API key and region
const apiKey = "RGAPI-a4a03f5d-d354-4ac3-878b-8a9b23fcfe5d"

func main() {

	account := getUserInput()

	urlACCOUNT := fmt.Sprintf("https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s?api_key=%s", account, apiKey)
	accountInfo := ACCOUNT(urlACCOUNT)
	puuid := accountInfo.Puuid
	fmt.Println(puuid)

	urlSUMMONER := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/%s?api_key=%s", puuid, apiKey)
	summonerInfo := SUMMONER(urlSUMMONER)
	summonerID := summonerInfo.ID
	fmt.Println(summonerID)

	urlLEAGUE := fmt.Sprintf("https://na1.api.riotgames.com/lol/league/v4/entries/by-summoner/%s?api_key=%s", summonerID, apiKey)
	leagueInfo := LEAGUE(urlLEAGUE)
	account, tier, rank, wins, losses := account, leagueInfo.Tier, leagueInfo.Rank, leagueInfo.Wins, leagueInfo.Losses
	fmt.Println(tier, rank, wins, losses)

	rankedInfo := RankedInfo{
		Account: account,
		Tier:    tier,
		Rank:    rank,
		Wins:    wins,
		Losses:  losses,
	}

	http.HandleFunc("/leagues", func(w http.ResponseWriter, r *http.Request) {
		leaguesHandler(w, r, rankedInfo)
	})

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	select {} // This will keep the program running indefinitely
}

func getUserInput() string {
	summonerName := "Duke"
	summonerTag := "NA9"

	return fmt.Sprintf("%s/%s", summonerName, summonerTag)
}

func leaguesHandler(w http.ResponseWriter, r *http.Request, rankedInfo RankedInfo) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, rankedInfo)
}
