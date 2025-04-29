package server

import (
	"context"
	"fmt"
	"go_url_shortener/services"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
)

type UrlShortenerServer struct {
	service      services.UrlShortenerService
	cacheService services.UrlShortenerCacheService
	router       http.Handler
	config       Config
}

// refactor this to have the config load everything
func NewUrlShortenerServer(config Config) *UrlShortenerServer {
	h := &UrlShortenerServer{
		service:      services.NewFirestoreUrlShortenerService(config),
		cacheService: services.NewRedisUrlShortenerCacheService(config),
		config:       config,
	}

	h.loadRoutes()

	return h, nil
}

func (u *UrlShortenerServer) StartUrlShortenerServer(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", u.config.ServerPort),
		Handler: u.router,
	}

	//ping redis to confirm connection
	err := u.cacheService.Client.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		if err := u.cacheService.Client.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	//make firestore connection
	app, err := firebase.NewApp(ctx, u.service.firebaseConfig)
	if err != nil {
		return fmt.Errorf("error making new firebase app: %w", err)
	}
	database, err := app.Firestore(ctx)
	if err != nil {
		return fmt.Println("error making firestore connection: %w", err)
	}
	//doing this here feels weird?
	u.service.database = database

	defer func() {
		if err := u.service.database.Close(); err != nil {
			fmt.Println("failed to close firestore", err)
		}
	}()

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
