package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/oreshkindev/profilss.ru-backend/internal"
)

type Router struct {
	*chi.Mux
}

func NewRouter(manager *internal.Manager) (*Router, error) {

	r := chi.NewRouter()

	// Set CORS options once, not per request
	corsOpts := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}
	r.Use(cors.Handler(corsOpts))

	/*
		SetContentType is a middleware function that sets the Content-Type header
		of the response to the specified content type.
	*/

	// Use SetContentType once, not per request
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Use the logger middleware for logging requests
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/users", userRouter(manager))
	})

	return &Router{r}, nil
}

func userRouter(manager *internal.Manager) http.Handler {
	r := chi.NewRouter()

	r.Get("/", manager.User.UserController.Get)
	// r.Route("/{userID}", func(r chi.Router) {
	// 	r.Get("/", getUser)
	// })

	return r
}
