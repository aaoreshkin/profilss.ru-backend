package controller

import (
	"net/http"

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

func (controller *UserController) Get(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	println(result)

	// render.JSON(w, r, result)
}
