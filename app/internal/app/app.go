package app

import (
	"context"
	"flag"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpcServer
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func NewApp(ctx context.Context) (*App, error) {

}

func (a *App) Run(ctx context.Context) error {

}
