package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	"github.com/richmondwang/golang-wallet-api/cmd/utils"
	"github.com/richmondwang/golang-wallet-api/pkg"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	client, err := utils.GetDBClient()
	if err != nil {
		log.Fatalf("connection postgres failed: %v", err)
	}
	defer client.Close()

	server := &pkg.Server{
		Port: utils.GetEnvOrDefault("PORT", "8080"),
		DB:   client,
	}

	server.Serve(ctx)
}
