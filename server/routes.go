package server

import (
	"go_url_shortener/handler"
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

	u.Router = router
}

func (u *UrlShortenerServer) loadUrlRoutes(router chi.Router) {
	urlHandler := &handler.Url{
		Service: u.Service,
	}

	router.Get("/", urlHandler.GetURLForm)
	router.Post("/", urlHandler.CreateShortUrl)
	router.Get("/{id}", urlHandler.GetLongUrl)
}

func (u *UrlShortenerServer) loadUserRoutes(router chi.Router) {
	//build out user related routes
}
