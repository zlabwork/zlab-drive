package srv

import (
	"app"
	"app/srv/drive/local"
	"app/srv/drive/oss"
	"app/srv/drive/s3"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

type FileRepository interface {
	Get(key string) (*app.File, error)
	List(key string, offset int, limit int) ([]*app.File, error)
	Create(file *app.File) error
	Delete(key string) error
	Modify(key string, newFile *app.File) error
	Bytes(file *app.File) ([]byte, error)
}

type FileService struct {
	Repo FileRepository
}

func (fs *FileService) Get(key string) (*app.File, error) {

	id, err := base64.RawURLEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return fs.Repo.Get(string(id))
}

func (fs *FileService) List(key string, offset int, limit int) ([]*app.File, error) {

	id := ""
	if len(key) > 2 {
		s, err := base64.RawURLEncoding.DecodeString(key)
		if err != nil {
			return nil, err
		}
		id = string(s)
	}

	return fs.Repo.List(id, offset, limit)
}

func (fs *FileService) Create(file *app.File) error {
	return fs.Repo.Create(file)
}

func (fs *FileService) Delete(key string) error {

	id, err := base64.RawURLEncoding.DecodeString(key)
	if err != nil {
		return err
	}

	return fs.Repo.Delete(string(id))
}

func (fs *FileService) Modify(key string, newFile *app.File) error {
	return fs.Repo.Modify(key, newFile)
}

func (fs *FileService) Bytes(file *app.File) ([]byte, error) {
	return fs.Repo.Bytes(file)
}

func NewFileService() (*FileService, error) {

	appDrive, err := strconv.Atoi(os.Getenv("APP_DRIVE"))
	if err != nil {
		return nil, err
	}

	switch appDrive {

	case app.LocalDrive:
		return &FileService{Repo: local.NewDrive()}, nil

	case app.S3Drive:
		return &FileService{Repo: s3.NewDrive()}, nil

	case app.OssDrive:
		return &FileService{Repo: oss.NewDrive()}, nil

	}

	return nil, fmt.Errorf("error drive")

}
