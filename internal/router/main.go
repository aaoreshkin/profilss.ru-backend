package router

import (
	"fmt"
	"net/http"
	"strings"

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
		r.Mount("/bid", router.BidHandler())
		r.Mount("/post", router.PostHandler())
		r.Mount("/service", router.ServiceHandler())
		r.Mount("/product", router.ProductHandler())
		r.Mount("/user", router.UserHandler())
		r.Mount("/doc", router.DocHandler())
		r.Mount("/hr", router.HrHandler())
	})
	return router, nil
}

func (router *Router) DocHandler() chi.Router {
	r := chi.NewRouter()

	controller := router.manager.Doc.DocController

	r.With(router.RBACMiddleware([]Rule{Superuser, Manager})).Post("/", controller.Create)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Delete("/{id}", controller.Delete)

	return r
}

func (router *Router) BidHandler() chi.Router {
	r := chi.NewRouter()

	controller := router.manager.Bid.BidController

	r.Post("/", controller.Create)
	r.With(router.RBACMiddleware([]Rule{Superuser, Manager})).Get("/", controller.Find)
	r.With(router.RBACMiddleware([]Rule{Superuser, Manager})).Get("/{id}", controller.First)
	r.With(router.RBACMiddleware([]Rule{Superuser, Manager})).Put("/{id}", controller.Update)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Delete("/{id}", controller.Delete)

	return r
}

func (router *Router) HrHandler() chi.Router {
	r := chi.NewRouter()

	controller := router.manager.Hr.HrController

	r.Post("/", controller.Create)
	r.With(router.RBACMiddleware([]Rule{Superuser, Manager})).Get("/", controller.Find)
	r.With(router.RBACMiddleware([]Rule{Superuser, Manager})).Get("/{id}", controller.First)
	r.With(router.RBACMiddleware([]Rule{Superuser, Manager})).Put("/{id}", controller.Update)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Delete("/{id}", controller.Delete)

	return r
}

func (router *Router) PostHandler() chi.Router {
	r := chi.NewRouter()

	controller := router.manager.Post.PostController

	r.With(router.RBACMiddleware([]Rule{Superuser})).Post("/", controller.Create)
	r.Get("/", controller.Find)
	r.Get("/{id}", controller.First)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Put("/{id}", controller.Update)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Delete("/{id}", controller.Delete)

	return r
}

func (router *Router) ServiceHandler() chi.Router {
	r := chi.NewRouter()

	controller := router.manager.Service.ServiceController

	r.With(router.RBACMiddleware([]Rule{Superuser})).Post("/", controller.Create)
	r.Get("/", controller.Find)
	r.Get("/{id}", controller.First)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Put("/{id}", controller.Update)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Delete("/{id}", controller.Delete)

	return r
}

func (router *Router) ProductHandler() chi.Router {
	r := chi.NewRouter()

	controller := router.manager.Product.ProductController

	r.With(router.RBACMiddleware([]Rule{Superuser})).Post("/", controller.Create)
	r.Get("/", controller.Find)
	r.Get("/{id}", controller.First)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Put("/{id}", controller.Update)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Delete("/{id}", controller.Delete)

	r.Mount("/characteristic", router.CharacteristicHandler())
	r.Mount("/category", router.CategoryHandler())
	r.Mount("/measure", router.MeasureHandler())

	return r
}

func (router *Router) CharacteristicHandler() chi.Router {
	r := chi.NewRouter()

	controller := router.manager.Product.CharacteristicController

	r.With(router.RBACMiddleware([]Rule{Superuser})).Post("/", controller.Create)
	r.Get("/", controller.Find)
	r.Get("/{id}", controller.First)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Delete("/{id}", controller.Delete)

	return r
}

func (router *Router) CategoryHandler() chi.Router {
	r := chi.NewRouter()

	controller := router.manager.Product.CategoryController

	r.With(router.RBACMiddleware([]Rule{Superuser})).Post("/", controller.Create)
	r.Get("/", controller.Find)
	r.Get("/{id}", controller.First)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Delete("/{id}", controller.Delete)

	return r
}

func (router *Router) MeasureHandler() chi.Router {
	r := chi.NewRouter()

	controller := router.manager.Product.MeasureController

	r.With(router.RBACMiddleware([]Rule{Superuser})).Post("/", controller.Create)
	r.Get("/", controller.Find)
	r.Get("/{id}", controller.First)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Delete("/{id}", controller.Delete)

	return r
}

func (router *Router) UserHandler() chi.Router {
	r := chi.NewRouter()

	controller := router.manager.User.UserController

	r.With(router.RBACMiddleware([]Rule{Superuser})).Post("/", controller.Create)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Get("/", controller.Find)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Get("/{email}", controller.First)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Put("/{id}", controller.Update)
	r.With(router.RBACMiddleware([]Rule{Superuser})).Delete("/{id}", controller.Delete)
	r.Post("/account", controller.Verify)

	return r
}

func (router *Router) RBACMiddleware(requiredRule []Rule) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("empty token")))
				return
			}
			if !strings.HasPrefix(tokenString, "Bearer ") {
				render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("missing or invalid token prefix: %s", tokenString)))
				return
			}
			tokenString = tokenString[len("Bearer "):]

			// Parse the token string into a jwt.Token struct
			parsedToken, err := common.ParseToken(tokenString)
			if err != nil {
				render.Render(w, r, common.ErrUnauthorized(fmt.Errorf("Token is invalid or expired: %s", err.Error())))
				return
			}

			// Get permissionID from token
			permissionID := common.GetPermissionID(parsedToken)

			// Get permission rule from database by permissionID
			permission, err := router.manager.User.PermissionController.First(permissionID)
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
