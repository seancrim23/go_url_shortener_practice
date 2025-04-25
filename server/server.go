package server

import (
	"go_url_shortener/services"
	"net/http"
)

type UrlShortenerServer struct {
	service      services.UrlShortenerService
	cacheService services.UrlShortenerCacheService
	router       http.Handler
	config       Config
}

// refactor this to have the config load everything
func NewUrlShortenerServer(config Config) (*UrlShortenerServer, error) {
	h := new(UrlShortenerServer)
	h := &UrlShortenerServer{}

	h.service = service
	h.cacheService = cacheService

	h.loadRoutes()
	//add routing figure out what the new good router is

	return h, nil
}
