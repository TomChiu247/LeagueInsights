package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TomChiu247/LeagueInsights/src/backend/util"
)

// PlayerResponse handles player response
type PlayerResponse struct {
	Name string `json:"name"`
}


//func enableCorss(w *http.ResponseWriter) {
//	(*w).Header().Set("Access-Control-Allow-Origin", "*")
//}

// PlayerHandler handles player
func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")

	name := r.URL.Query().Get("name")

	apiKey := "RGAPI-1ec59ed7-03e5-4eb2-b614-2b304e5461c9"

	resp, err := http.Get(fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", name, apiKey))

	if err != nil {
		log.Fatalln(err)
	}

	var playerData map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&playerData)
	if err != nil {
		log.Fatalln(err)
	}

	accountID := playerData["accountId"]
	respMatch, errMatch := http.Get(fmt.Sprintf("https://na1.api.riotgames.com/lol/match/v4/matchlists/by-account/%s?api_key=%s&endIndex=10", accountID, apiKey))

	if errMatch != nil {
		log.Fatalln(errMatch)
	}

	var matchData map[string]interface{}

	err = json.NewDecoder(respMatch.Body).Decode(&matchData)
	if err != nil {
		log.Fatalln(err)
	}

	util.JSON(w, r, http.StatusOK, playerData)
	util.JSON(w, r, http.StatusOK, matchData)
}
