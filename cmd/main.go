package main

import (
	"awesomeProject/config"
	"awesomeProject/internal/handler/cron"
	"awesomeProject/internal/handler/http"
	"awesomeProject/internal/repository/postgres"
	"awesomeProject/internal/usecase/price"
	"awesomeProject/pkg/steam_data/lunex"
	"context"
	httpServer "net/http"
)

var (
	// Repositories
	postgreSQL *postgres.Postgres

	// Services
	lunexHTTPClient *lunex.HttpClient

	// Usecases
	getPriceUseCase   *price.GetPricesUseCase
	fetchPriceUseCase *price.FetchPricesUseCase

	// Handlers
	priceHandler   *http.PriceHandler
	fetchPriceTask *cron.FetchPriceTask
)

func main() {
	cfg := config.MustLoad()

	ctx := context.Background()

	if err := createRepositories(ctx, cfg); err != nil {
		panic(err)
	}
	defer postgreSQL.Close()

	createServices()
	createUseCases()
	createHandlers()

	if err := run(cfg); err != nil {
		panic(err)
	}

}

func run(cfg *config.Config) error {
	if _, err := cron.RegisterSchedule(fetchPriceTask); err != nil {
		return err
	}

	router := http.NewRouter(priceHandler)

	server := &httpServer.Server{
		Addr:         ":" + cfg.HttpServer.Port,
		Handler:      router,
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func createRepositories(ctx context.Context, cfg *config.Config) error {
	var err error
	postgreSQL, err = postgres.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}

	return nil
}

func createServices() {
	lunexHTTPClient = lunex.NewClient()
}

func createUseCases() {
	getPriceUseCase = price.NewGetPricesUseCase(postgreSQL)
	fetchPriceUseCase = price.NewFetchPricesUseCase(lunexHTTPClient, postgreSQL)
}

func createHandlers() {
	// HTTP handlers
	priceHandler = http.NewPriceHandler(getPriceUseCase, fetchPriceUseCase)

	// Cron commands
	fetchPriceTask = cron.NewFetchPriceTask(fetchPriceUseCase)
}
