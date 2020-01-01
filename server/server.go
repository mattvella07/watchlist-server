package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mattvella07/watchlist-server/server/api"
)

const port = ":8080"

var (
	allowedHeaders = []string{"X-Requested-With", "Content-Type", "Authorization"}
	allowedMethods = []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}
	allowedOrigins = []string{"*"}
)

func setupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/watchlist", api.GetWatchlist).Methods("GET")

	return router
}

// Start sets up the http server and listens for incoming requests
func Start() {
	router := setupRoutes()

	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders(allowedHeaders), handlers.AllowedMethods(allowedMethods), handlers.AllowedOrigins(allowedOrigins))(router)))
}
