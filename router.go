package main

import (
	"net/http"

	"github.com/denis-onder/archibald/handlers"
	"github.com/denis-onder/archibald/middleware"
)

func initializeRoutes(mux *http.ServeMux) {
	mux.Handle("/welcome", middleware.Logger(handlers.Welcome()))
}
