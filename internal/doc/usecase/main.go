package usecase

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type DocUsecase struct {
}

var (
	destination = os.Getenv("REMOTE_PATH")
)

func NewDocUsecase() *DocUsecase {
	return &DocUsecase{}
}

func (usecase *DocUsecase) Create(fileName string, fileBody multipart.File) error {

	var (
		path string
	)

	if filepath.Ext(fileName) == ".png" || filepath.Ext(fileName) == ".jpg" || filepath.Ext(fileName) == ".jpeg" {
		path = "/in/"
	} else {
		path = "/media/"
	}

	if _, err := os.Stat(filepath.Join(destination, path)); os.IsNotExist(err) {
		err := os.Mkdir(filepath.Join(destination, path), os.ModePerm)
		if err != nil {
			return err
		}
	}

	r, err := os.Create(filepath.Join(destination, path, fileName))
	if err != nil {

		return err
	}
	defer r.Close()

	if _, err := io.Copy(r, fileBody); err != nil {
		return err
	}

	return nil
}

func (usecase *DocUsecase) Delete(id string) error {
	return nil
}
