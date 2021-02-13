package router

import (
	"github.com/TomChiu247/LeagueInsights/src/backend/handler"
	"github.com/gorilla/mux"
)

// NewRouter returns a router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handler.MainHandler).Methods("GET")
	r.HandleFunc("/player", handler.PlayerHandler).Methods("GET")
	return r
}
