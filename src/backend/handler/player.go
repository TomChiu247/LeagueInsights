package handler

import (
	"fmt"
	"io/ioutil"
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

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)

	util.JSON(w, r, http.StatusOK, sb)
}
