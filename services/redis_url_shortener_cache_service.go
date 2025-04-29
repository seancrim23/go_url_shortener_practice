package services

//implement interface functions

type RedisUrlShortenerCacheService struct {
	Client *redis.Client
}

func NewRedisUrlShortenerCacheService( redisAddress string, redisUsername string, redisPassword string) (*RedisUrlShortenerCacheService) {
	return &RedisUrlShortenerCacheService{Client: redis.NewClient(&redis.Options{
		Addr: redisAddress,
		Username: redisUsername,
		Password: redisPassword,
		DB: 0,
	}}
}

func (r *RedisUrlShortenerCacheService) Get(shortUrl string) (string, error) {
	//look for shorturl in cache
	//pull longurl (and maybe all of the metadata is there too as a json)
	//or error
}

func (r *RedisUrlShortenerCacheService) Put(shortUrl string, longUrlJson string) (error) {
	//determine some criteria to where we'd do this (maybe the url is pulled x amount of times in a certain time period)
	//put the shorturl as the key to the longurl json value
	//better performing access than continually hitting db
	//or error
}