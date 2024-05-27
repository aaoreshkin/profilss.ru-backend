package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/oreshkindev/profilss.ru-backend/common"
	"github.com/oreshkindev/profilss.ru-backend/internal/support/entity"
)

type SupportController struct {
	usecase entity.SupportUsecase
}

func NewSupportController(usecase entity.SupportUsecase) *SupportController {

	return &SupportController{
		usecase: usecase,
	}
}

func (controller *SupportController) Broadcast(message []byte) ([]byte, error) {

	var (
		response entity.Support
		err      error
	)

	if err = json.Unmarshal(message, &response); err != nil {
		return nil, err
	}

	if response.Message == "" {
		var (
			history []entity.Support
			bytes   []byte
		)

		if history, err = controller.usecase.First(response.SessionID); err != nil {
			return nil, err
		}

		if bytes, err = json.Marshal(history); err != nil {
			return nil, err
		}

		return bytes, nil
	} else {
		var (
			result *entity.Support
			bytes  []byte
		)

		if result, err = controller.usecase.Create(&response); err != nil {
			return nil, err
		}

		if bytes, err = json.Marshal(result); err != nil {
			return nil, err
		}

		return bytes, nil
	}
}

func (controller *SupportController) Create(w http.ResponseWriter, r *http.Request) {
	entity := &entity.Support{}

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

func (controller *SupportController) Find(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Find()
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *SupportController) First(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	result, err := controller.usecase.First(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *SupportController) Update(w http.ResponseWriter, r *http.Request) {

	entity := &entity.Support{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Update(entity)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *SupportController) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	err := controller.usecase.Delete(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, nil)
}
