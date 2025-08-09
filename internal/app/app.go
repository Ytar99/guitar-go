package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpswagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"

	_ "guitar-go/docs"
	"guitar-go/internal/config"
	"guitar-go/internal/db"
	"guitar-go/internal/handlers"
	"guitar-go/internal/middleware"
	"guitar-go/internal/models"
	"guitar-go/internal/repositories"
	"guitar-go/internal/services"
)

type App struct {
	Config *config.Config
	Router *mux.Router
	DB     *gorm.DB
}

func NewApp(cfg *config.Config) *App {
	return &App{
		Config: cfg,
		Router: mux.NewRouter(),
	}
}

func (a *App) Init() error {
	database, err := db.NewDatabase(a.Config)
	if err != nil {
		return err
	}
	if err := database.Connect(); err != nil {
		return err
	}

	// Verify database connection
	if err := database.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	a.DB = database.GetDB()

	if err := database.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	userRepo := repositories.NewUserRepository(a.DB)
	authService := services.NewAuthService(a.Config)
	userService := services.NewUserService(userRepo)

	m := middleware.NewMiddleware(authService)
	appHandlers := handlers.NewHandlers(authService, userService)

	a.SetupRoutes(appHandlers, m)
	a.Router.Use(loggingMiddleware)

	return nil
}

func (a *App) SetupRoutes(h *handlers.Handlers, m *middleware.Middleware) {
	a.Router.PathPrefix("/swagger/").Handler(httpswagger.WrapHandler)
	a.Router.HandleFunc("/login", h.Login).Methods("POST")
	a.Router.HandleFunc("/register", h.CreateUser).Methods("POST")

	api := a.Router.PathPrefix("/api").Subrouter()
	api.Use(m.Auth)

	admin := api.PathPrefix("/admin").Subrouter()
	admin.Use(m.RequireRole("admin"))
	admin.HandleFunc("/users", h.GetUsers).Methods("GET")
	admin.HandleFunc("/users", h.CreateUser).Methods("POST")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
