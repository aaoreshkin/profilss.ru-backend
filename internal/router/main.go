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

type (
	Router struct {
		*chi.Mux
		manager *internal.Manager
	}

	Rule string
)

const (
	Superuser Rule = "Superuser"
	Manager   Rule = "Manager"
)

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
	router.With(router.RBACMiddleware([]Rule{Superuser, Manager})).Get("/", router.manager.User.UserController.Get)

	return router
}

func (router *Router) RBACMiddleware(requiredRule []Rule) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("empty token")))
				return
			}
			tokenString = tokenString[len("Bearer "):]

			// Parse the token string into a jwt.Token struct
			parsedToken, err := common.ParseToken(tokenString)
			if err != nil {
				render.Render(w, r, common.ErrInvalidRequest(err))
				return
			}

			// Get permissionID from token
			permissionID := common.GetPermissionID(parsedToken)

			// Get permission rule from database by permissionID
			permission, err := router.manager.User.PermissionController.Get(permissionID)
			if err != nil {
				render.Render(w, r, common.ErrInvalidRequest(err))
				return
			}

			// Check if permission rule has required permission
			for _, rule := range requiredRule {
				if permission.Rule == string(rule) {
					next.ServeHTTP(w, r)
					return
				}

				render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("permission rule %s is not allowed", permission.Rule)))
			}
		})
	}
}
