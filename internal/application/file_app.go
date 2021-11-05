package application

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

type File struct {
	repo port.FileRepo
}

func NewFileApp(repo port.FileRepo) *File {
	return &File{
		repo: repo,
	}
}

func handleImage() {

}
