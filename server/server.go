package server

import (
	"go_url_shortener/services"
	"net/http"
)

type UrlShortenerServer struct {
	service      services.UrlShortenerService
	cacheService services.UrlShortenerCacheService
	http.Handler
}

func NewUrlShortenerServer(service services.UrlShortenerService, cacheService services.UrlShortenerCacheService) (*UrlShortenerServer, error) {
	h := new(UrlShortenerServer)

	h.service = service
	h.cacheService = cacheService

	//add routing figure out what the new good router is

	return h, nil
}
