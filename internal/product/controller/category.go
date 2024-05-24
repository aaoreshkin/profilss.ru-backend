package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/oreshkindev/profilss.ru-backend/common"
	"github.com/oreshkindev/profilss.ru-backend/internal/product/entity"
)

type CategoryController struct {
	usecase entity.CategoryUsecase
}

func NewCategoryController(usecase entity.CategoryUsecase) *CategoryController {
	return &CategoryController{
		usecase: usecase,
	}
}

func (controller *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	entity := &entity.Category{}

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

func (controller *CategoryController) Find(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Find()
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *CategoryController) First(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	result, err := controller.usecase.First(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	entity := &entity.Category{}

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

func (controller *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	err := controller.usecase.Delete(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, nil)
}
