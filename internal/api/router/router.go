package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yusufhalabi/quran-search/internal/api/handlers"
)

func Setup() *mux.Router {
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/api/search", handlers.Search).Methods(http.MethodGet)
	r.HandleFunc("/api/suggestions", handlers.GetSuggestions).Methods(http.MethodGet)

	return r
}
