package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"back-train/config"
	"back-train/internal/delivery/http/handler"
	"back-train/internal/delivery/http/router"
	"back-train/internal/repository"
	"back-train/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// Load Konfigurasi
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Koneksi Databaseç
	dbPool, err := pgxpool.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbPool.Close()

	// Inisialisasi Fiber
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Inisialisasi Layers (Dependency Injection)
	// Repository
	userRepo := repository.NewUserRepository(dbPool)
	alumniRepo := repository.NewAlumniRepository(dbPool)
	mahasiswaRepo := repository.NewMahasiswaRepository(dbPool)
	pekerjaanRepo := repository.NewPekerjaanRepository(dbPool)

	// Usecase (Service)
	authUsecase := usecase.NewAuthUsecase(userRepo, cfg.JWTSecretKey, cfg.JWTExpirationHours)
	alumniUsecase := usecase.NewAlumniUsecase(alumniRepo)
	mahasiswaUsecase := usecase.NewMahasiswaUsecase(mahasiswaRepo)
	pekerjaanUsecase := usecase.NewPekerjaanUsecase(pekerjaanRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authUsecase)
	alumniHandler := handler.NewAlumniHandler(alumniUsecase)
	mahasiswaHandler := handler.NewMahasiswaHandler(mahasiswaUsecase)
	pekerjaanHandler := handler.NewPekerjaanHandler(pekerjaanUsecase)

	// Setup Router
	router.SetupRoutes(app, authHandler, alumniHandler, mahasiswaHandler, pekerjaanHandler, cfg)

	// Start Server
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server is running on port %s", cfg.ServerPort)
	err = app.Listen(serverAddr)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
		os.Exit(1)
	}
}
