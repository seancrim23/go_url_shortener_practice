package server

import "os"

// build this out further with db info...
type Config struct {
	RedisAddress  string
	RedisUsername string
	RedisPass     string
	ServerPort    uint16
}

// build out the setters along with what i add to the struct
// just should be whatever app specific stuff
func LoadConfig() Config {
	cfg := Config{
		RedisAddress:  "tbd",
		RedisUsername: "default",
		ServerPort:    3000,
	}

	if redisPass, exists := os.LookupEnv("REDIS_PASS"); exists {
		cfg.RedisPass = redisPass
	}

	return cfg
}
