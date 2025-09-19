package router

import (
	"back-train/config"
	"back-train/internal/delivery/http/handler"
	"back-train/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	authHandler *handler.AuthHandler,
	alumniHandler *handler.AlumniHandler,
	mahasiswaHandler *handler.MahasiswaHandler,
	pekerjaanHandler *handler.PekerjaanHandler,
	cfg *config.Config,
) {
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Middleware
	authMiddleware := middleware.AuthMiddleware(cfg.JWTSecretKey)
	adminMiddleware := middleware.RoleMiddleware("admin")

	// Alumni routes
	alumni := api.Group("/alumni", authMiddleware)
	alumni.Get("/", alumniHandler.GetAllAlumni)
	alumni.Get("/:id", alumniHandler.GetAlumniByID)
	alumni.Post("/", adminMiddleware, alumniHandler.CreateAlumni)
	alumni.Put("/:id", adminMiddleware, alumniHandler.UpdateAlumni)
	alumni.Delete("/:id", adminMiddleware, alumniHandler.DeleteAlumni)

	// Mahasiswa routes
	mahasiswa := api.Group("/mahasiswa", authMiddleware)
	mahasiswa.Get("/", mahasiswaHandler.GetAllMahasiswa)
	mahasiswa.Get("/:id", mahasiswaHandler.GetMahasiswaByID)
	mahasiswa.Post("/", adminMiddleware, mahasiswaHandler.CreateMahasiswa)
	mahasiswa.Put("/:id", adminMiddleware, mahasiswaHandler.UpdateMahasiswa)
	mahasiswa.Delete("/:id", adminMiddleware, mahasiswaHandler.DeleteMahasiswa)

	// Pekerjaan routes
	pekerjaan := api.Group("/pekerjaan", authMiddleware)
	pekerjaan.Get("/", pekerjaanHandler.GetAllPekerjaan)
	pekerjaan.Get("/:id", pekerjaanHandler.GetPekerjaanByID)
	pekerjaan.Post("/", adminMiddleware, pekerjaanHandler.CreatePekerjaan)
	pekerjaan.Put("/:id", adminMiddleware, pekerjaanHandler.UpdatePekerjaan)
	pekerjaan.Delete("/:id", adminMiddleware, pekerjaanHandler.DeletePekerjaan)
}
