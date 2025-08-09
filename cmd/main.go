package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"guitar-go/internal/app"
	"guitar-go/internal/config"
)

// @title Guitar Go API
// @version 1.0
// @description API for managing guitar chords

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configPath := flag.String("config", "./configs/config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if cfg.Server.Port == "" {
		log.Fatal("Server port must be configured")
	}
	if cfg.JWT.Secret == "" || cfg.JWT.Secret == "your-secret-key" {
		log.Fatal("JWT secret must be properly configured")
	}

	application := app.NewApp(cfg)
	if err := application.Init(); err != nil {
		log.Fatalf("Failed to init app: %v", err)
	}

	server := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: application.Router,
	}

	log.Printf("Server is running on port %s", cfg.Server.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
