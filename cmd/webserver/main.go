package main

import (
	"context"
	"fmt"
	"go_url_shortener/server"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("hello world")
	server := server.NewUrlShortenerServer(server.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := server.StartUrlShortenerServer(ctx)
	if err != nil {
		fmt.Println("failed to start url shortener server: ", err)
	}
}
