package services

// TODO think about this some more but probably wont need more than just get or put
// think of more specific functions, what are we getting and putting
type UrlShortenerCacheService interface {
	Get(string) (string, error)
	Put(string, string) error
}
