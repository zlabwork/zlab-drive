package srv

import (
	"drive"
	"drive/srv/fs"
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
	return fs.Repo.Get(key)
}

func (fs *FileService) List(key string, offset int, limit int) ([]*drive.File, error) {
	return fs.Repo.List(key, offset, limit)
}

func (fs *FileService) Create(file *drive.File) error {
	return fs.Repo.Create(file)
}

func (fs *FileService) Delete(key string) error {
	return fs.Repo.Delete(key)
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
