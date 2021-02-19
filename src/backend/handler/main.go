package handler

import (
	"net/http"

	"github.com/TomChiu247/LeagueInsights/src/backend/util"
)

// MainResponse handles main response
type MainResponse struct {
	Data string `json:"data"`
}

//func enableCors(w *http.ResponseWriter) {
//	(*w).Header().Set("Access-Control-Allow-Origin", "*")
//}

// MainHandler handles main
func MainHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")

	data := r.URL.Query().Get("data")

	resp := MainResponse{
		Data: data,
	}

	util.JSON(w, r, http.StatusOK, resp)
}
