package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/olahol/melody"

	"github.com/oreshkindev/profilss.ru-backend/common"
	"github.com/oreshkindev/profilss.ru-backend/internal/chat/entity"
)

type ChatController struct {
	usecase entity.ChatUsecase
	m       *melody.Melody
}

func NewChatController(usecase entity.ChatUsecase) *ChatController {
	m := melody.New()

	return &ChatController{
		usecase: usecase,
		m:       m,
	}
}

func (controller *ChatController) Broadcast(w http.ResponseWriter, r *http.Request) {

	var (
		response entity.Chat
		err      error
	)

	controller.m.HandleRequest(w, r)

	controller.m.HandleMessage(func(s *melody.Session, message []byte) {

		if err = json.Unmarshal(message, &response); err != nil {
			log.Println("Ошибка десериализации:", err)
			return
		}

		if response.Message == "" {
			history, err := controller.usecase.First(response.SessionID)
			if err != nil {
				log.Println("Ошибка получения истории:", err)
				return
			}

			bytes, err := json.Marshal(history)
			if err != nil {
				log.Println("Ошибка конвертации в байты:", err)
			}

			controller.m.BroadcastFilter(bytes, func(q *melody.Session) bool {
				log.Println("q:", q.Request.URL.Path)
				log.Println("s:", s.Request.URL.Path)
				return q.Request.URL.Path == s.Request.URL.Path
			})
		} else {
			res, err := controller.usecase.Create(&response)
			if err != nil {
				log.Println("Ошибка сохранения истории:", err)
				return
			}

			bytes, err := json.Marshal(res)
			if err != nil {
				log.Println("Ошибка конвертации в байты:", err)
			}

			controller.m.BroadcastFilter(bytes, func(q *melody.Session) bool {
				log.Println("q:", q.Request.URL.Path)
				log.Println("s:", s.Request.URL.Path)
				return q.Request.URL.Path == s.Request.URL.Path
			})
		}
	})
}

func (controller *ChatController) Create(w http.ResponseWriter, r *http.Request) {
	entity := &entity.Chat{}

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

func (controller *ChatController) Find(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Find()
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *ChatController) First(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	result, err := controller.usecase.First(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *ChatController) Update(w http.ResponseWriter, r *http.Request) {

	entity := &entity.Chat{}

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

func (controller *ChatController) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	err := controller.usecase.Delete(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, nil)
}
