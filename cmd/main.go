package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/uxsnap/review_bot/internal/app"
	"github.com/uxsnap/review_bot/internal/config"
)

const defaultEnvPath = ".env"

func main() {
	if err := config.Init(defaultEnvPath); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	a, err := app.New()

	if err != nil {
		log.Println("\n === Cannot create app. === ")
		return
	}

	a.Run(ctx)
}
