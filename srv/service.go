package srv

import (
	"drive"
	"drive/srv/fs"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

type FileRepository interface {
	Get(key string) (*drive.File, error)
	List(key string, offset int, limit int) ([]*drive.File, error)
	Create(file *drive.File) error
	Delete(key string) error
	Modify(key string, newFile *drive.File) error
	Bytes(file *drive.File) ([]byte, error)
}

type FileService struct {
	Repo FileRepository
}

func (fs *FileService) Get(key string) (*drive.File, error) {

	id, err := base64.RawURLEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return fs.Repo.Get(string(id))
}

func (fs *FileService) List(key string, offset int, limit int) ([]*drive.File, error) {

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

func (fs *FileService) Create(file *drive.File) error {
	return fs.Repo.Create(file)
}

func (fs *FileService) Delete(key string) error {

	id, err := base64.RawURLEncoding.DecodeString(key)
	if err != nil {
		return err
	}

	return fs.Repo.Delete(string(id))
}

func (fs *FileService) Modify(key string, newFile *drive.File) error {
	return fs.Repo.Modify(key, newFile)
}

func (fs *FileService) Bytes(file *drive.File) ([]byte, error) {
	return fs.Repo.Bytes(file)
}

func NewFileService() (*FileService, error) {

	appDrive, err := strconv.Atoi(os.Getenv("APP_DRIVE"))
	if err != nil {
		return nil, err
	}

	switch appDrive {

	case drive.LocalDrive:
		return &FileService{Repo: fs.NewLocalDrive()}, nil

	case drive.S3Drive:
		return &FileService{Repo: fs.NewS3Drive()}, nil

	case drive.OssDrive:
		return &FileService{Repo: fs.NewOssDrive()}, nil

	}

	return nil, fmt.Errorf("error drive")

}
