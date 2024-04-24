package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/oreshkindev/profilss.ru-backend/common"
	"github.com/oreshkindev/profilss.ru-backend/internal"
)

type Router struct {
	*chi.Mux
	manager *internal.Manager
}

func NewRouter(manager *internal.Manager) (*Router, error) {
	router := &Router{chi.NewRouter(), manager}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Use(middleware.Logger)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/user", router.UserHandler())
	})
	return router, nil
}

func (router *Router) UserHandler() chi.Router {
	router.Get("/", router.RBACMiddleware("W")(router.manager.User.UserController.Get))
	return router
}

func (router *Router) RBACMiddleware(requiredPermissionRule string) func(http.HandlerFunc) http.HandlerFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("empty token")))
				return
			}
			tokenString = tokenString[len("Bearer "):]

			// Parse token and check it
			parsedToken, err := common.ParseToken(tokenString)
			if err != nil {
				render.Render(w, r, common.ErrInvalidRequest(err))
				return
			}

			// Get permissionID from token
			permissionID := common.GetPermissionID(parsedToken)

			// Get permission rule from database by permissionID
			permissionRule, err := router.manager.User.PermissionController.Get(permissionID)
			if err != nil {
				render.Render(w, r, common.ErrInvalidRequest(err))
				return
			}

			// Check if permission rule is equal required
			if permissionRule != requiredPermissionRule {
				render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("permission denied")))
				return
			}

			handler(w, r)
		}
	}
}
