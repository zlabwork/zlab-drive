package srv

import (
	"app"
	"app/srv/drive/local"
	"app/srv/drive/oss"
	"app/srv/drive/s3"
	"context"
	"fmt"
	"os"
	"strconv"
)

type DriveRepository interface {
	Get(key string) (*app.File, error)
	List(key string, offset int, limit int) ([]*app.File, error)
	Create(file *app.File) error
	Delete(key string) error
	Modify(key string, newFile *app.File) error
	Bytes(file *app.File) ([]byte, error)
}

type DriveService struct {
	Repo DriveRepository
}

func (fs *DriveService) Get(key string) (*app.File, error) {

	return fs.Repo.Get(key)
}

func (fs *DriveService) List(key string, offset int, limit int) ([]*app.File, error) {

	return fs.Repo.List(key, offset, limit)
}

func (fs *DriveService) Create(file *app.File) error {
	return fs.Repo.Create(file)
}

func (fs *DriveService) Delete(key string) error {

	return fs.Repo.Delete(key)
}

func (fs *DriveService) Modify(key string, newFile *app.File) error {
	return fs.Repo.Modify(key, newFile)
}

func (fs *DriveService) Bytes(file *app.File) ([]byte, error) {
	return fs.Repo.Bytes(file)
}

func NewDriveService(ctx context.Context) (*DriveService, error) {

	uid := ctx.Value(app.UserIdKey).(int64)
	namespace := strconv.FormatInt(uid, 10)

	appDrive, err := strconv.Atoi(os.Getenv("APP_DRIVE"))
	if err != nil {
		return nil, err
	}

	switch appDrive {

	case app.LocalDrive:
		return &DriveService{Repo: local.NewDrive(namespace)}, nil

	case app.S3Drive:
		return &DriveService{Repo: s3.NewDrive(namespace)}, nil

	case app.OssDrive:
		return &DriveService{Repo: oss.NewDrive(namespace)}, nil

	}

	return nil, fmt.Errorf("error drive")

}
