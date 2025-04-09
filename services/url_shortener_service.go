package services

type UrlShortenerService interface {
	CreateShortUrl(string) (string, error)
	GetLongUrl(string) error
}
