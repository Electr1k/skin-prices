package main

import (
	"awesomeProject/config"
	"awesomeProject/internal/handler"
	"awesomeProject/internal/repository/postgres"
	"awesomeProject/internal/usecase/price"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	if err := run(context.Background(), cfg); err != nil {
		panic(err)
	}
}

func run(ctx context.Context, cfg *config.Config) error {
	repo, err := postgres.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to create repository: %w", err)
	}
	defer repo.Close()

	priceUseCase := price.NewGetPricesUseCase(repo)

	priceHandler := handler.NewPriceHandler(priceUseCase)

	router := handler.NewRouter(priceHandler)

	server := &http.Server{
		Addr:         cfg.HttpServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Printf("Server starting on %s\n", cfg.HttpServer.Address)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	}()

	<-done
	fmt.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}

	fmt.Println("Server stopped")
	return nil
}
