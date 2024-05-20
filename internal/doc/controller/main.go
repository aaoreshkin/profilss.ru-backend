package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/oreshkindev/profilss.ru-backend/common"
	"github.com/oreshkindev/profilss.ru-backend/internal/doc/entity"
)

type DocController struct {
	usecase entity.DocUsecase
}

func NewDocController(usecase entity.DocUsecase) *DocController {
	return &DocController{
		usecase: usecase,
	}
}

func (controller *DocController) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	fileBody, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer fileBody.Close()

	if err := controller.usecase.Create(header.Filename, fileBody); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, nil)
}

func (controller *DocController) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	err := controller.usecase.Delete(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, nil)
}
