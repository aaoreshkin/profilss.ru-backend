package usecase

import (
	"io"
	"mime/multipart"
	"os"
)

type DocUsecase struct {
}

var (
	path = os.Getenv("REMOTE_PATH")
)

func NewDocUsecase() *DocUsecase {
	return &DocUsecase{}
}

func (usecase *DocUsecase) Create(fileName string, fileBody multipart.File) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path+"/in/", 0755)
		if err != nil {
			return err
		}
	}

	r, err := os.Create(path + "/in/" + fileName)
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
