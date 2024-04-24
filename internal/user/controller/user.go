package controller

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/oreshkindev/profilss.ru-backend/common"
	"github.com/oreshkindev/profilss.ru-backend/internal/user/entity"
)

type UserController struct {
	usecase entity.UserUsecase
}

func NewUserController(usecase entity.UserUsecase) *UserController {
	return &UserController{
		usecase: usecase,
	}
}

func (controller *UserController) Post(w http.ResponseWriter, r *http.Request) {
	entity := &entity.User{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Post(entity)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result.NewResponse())
}

func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Get()
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	for i := range result {
		result[i] = *result[i].NewResponse()
	}

	render.JSON(w, r, result)
}
