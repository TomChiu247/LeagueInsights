package handler

import (
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

	resp := PlayerResponse{
		Name: name,
	}

	util.JSON(w, r, http.StatusOK, resp)
}
