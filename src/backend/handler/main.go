package handler

import (
	"net/http"

	"github.com/TomChiu247/LeagueInsights/src/backend/util"
)

type MainResponse struct {
	Data string `json:"data"`
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")

	resp := MainResponse{
		Data: data,
	}
	util.JSON(w, r, http.StatusOK, resp)
}
