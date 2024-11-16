package main

import (
	"context"
	"log"

	"github.com/kms-qwe/auth/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	authApp, err := app.NewApp(ctx)
	if err != nil {
		log.Panicf("failed to init app: %s\n", err.Error())
	}

	err = authApp.Run(ctx, cancel)
	if err != nil {
		log.Panicf("failed to run app: %s\n", err.Error())
	}
}
