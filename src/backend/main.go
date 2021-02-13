package main

import (
	"net/http"

	"github.com/TomChiu247/LeagueInsights/src/backend/router"
	"github.com/rs/zerolog/log"
)

func main() {

	router := router.NewRouter()

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Error().Err(err).Msg("Failed to start http server")
	}
}
