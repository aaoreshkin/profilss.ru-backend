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
	err := r.ParseMultipartForm(32 << 20) // 32MB is the max memory use, adjust if needed
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	body, header, err := r.FormFile("file")
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}
	defer body.Close()

	if err := controller.usecase.Create(header.Filename, body); err != nil {
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
