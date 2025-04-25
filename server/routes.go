package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (u *UrlShortenerServer) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/url", u.loadUrlRoutes)
	router.Route("/user", u.loadUserRoutes)

	u.router = router
}

func (u *UrlShortenerServer) loadUrlRoutes(router chi.Router) {
	//build out url related routes
}

func (u *UrlShortenerServer) loadUserRoutes(router chi.Router) {
	//build out user related routes
}
