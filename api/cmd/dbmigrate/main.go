package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"
	"github.com/richmondwang/golang-wallet-api/cmd/utils"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	client, err := utils.GetDBClient()
	if err != nil {
		log.Fatalf("connection postgres failed: %v", err)
	}
	defer client.Close()

	// migrate
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
