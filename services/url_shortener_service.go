package services

import "context"

type UrlShortenerService interface {
	CreateShortUrl(context.Context, string) (string, error)
	GetLongUrl(context.Context, string) (string, error)
}
