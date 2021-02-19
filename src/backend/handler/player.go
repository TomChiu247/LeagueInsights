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

	apiKey := "RGAPI-91fdb58d-b815-4f85-93de-0f2d5505f0ae"

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
