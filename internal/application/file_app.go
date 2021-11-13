package application

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

type FileApp struct {
	repo port.FileRepo
}

func NewFileApp(repo port.FileRepo) *FileApp {
	return &FileApp{
		repo: repo,
	}
}

func (file FileApp) HandleImage() {

}
