package entity

import (
	"mime/multipart"
)

type (
	Doc struct {
	}

	DocUsecase interface {
		Create(string, multipart.File) error
		Delete(string) error
	}
)

// fields of struct that will be returned
func (response *Doc) NewResponse() *Doc {
	return &Doc{}
}
