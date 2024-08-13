package routers

import (
	"github.com/fengye2419/content-moderation-service/internal/api"
	"github.com/fengye2419/content-moderation-service/internal/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// InitializeRouter 初始化路由
func InitializeRouter() *chi.Mux {
	// 初始化 xorm 引擎
	model.InitializeEngine("config/app.ini")

	// Initialize the router
	r := chi.NewRouter()

	// Use logging middleware
	r.Use(middleware.Logger)

	// Swagger UI endpoint
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/sensitiveWords", func(r chi.Router) {
			r.Post("/", api.AddSensitiveWordHandler)
			r.Delete("/{id}", api.DeleteSensitiveWordHandler)

			// Define the /validate endpoint
			r.Post("/validate", api.ValidateHandler)
		})
	})

	return r
}
