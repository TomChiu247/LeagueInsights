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

// PlayerHandler handles player
func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	apiKey := "RGAPI-5a3f34a8-9f7d-4a27-8d3c-5cc3fea22667"

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
