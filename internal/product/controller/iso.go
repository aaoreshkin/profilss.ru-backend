package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/oreshkindev/profilss.ru-backend/common"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type IsoController struct {
	usecase entity.IsoUsecase
}

func NewIsoController(usecase entity.IsoUsecase) *IsoController {
	return &IsoController{
		usecase: usecase,
	}
}

func (controller *IsoController) Create(w http.ResponseWriter, r *http.Request) {
	entity := &entity.Iso{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Create(entity)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *IsoController) Find(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Find()
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *IsoController) First(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	result, err := controller.usecase.First(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *IsoController) Update(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	entity := &entity.Iso{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Update(entity, id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *IsoController) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	err := controller.usecase.Delete(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, nil)
}
