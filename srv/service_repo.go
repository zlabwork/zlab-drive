package srv

import (
	"app"
	"app/srv/repository/mysql"
)

type FileRepository interface {
	Get(id int64) (*app.File, error)
	GetByUUID(uuid string) (*app.File, error)
	Delete(id int64) error
	Create(file *app.File) (int64, error)
	Modify(file *app.File) error
	List(parent int64, offset int, size int) ([]*app.File, error)
}

type FileService struct {
	Repo FileRepository
}

func (rs *FileService) Get(id int64) (*app.File, error) {
	return rs.Repo.Get(id)
}

func (rs *FileService) GetByUUID(uuid string) (*app.File, error) {
	return rs.Repo.GetByUUID(uuid)
}

func (rs *FileService) Delete(id int64) error {
	return rs.Repo.Delete(id)
}

func (rs *FileService) Create(file *app.File) (int64, error) {
	return rs.Repo.Create(file)
}

func (rs *FileService) Modify(file *app.File) error {
	return rs.Repo.Modify(file)
}

func (rs *FileService) List(parent int64, offset int, size int) ([]*app.File, error) {
	return rs.Repo.List(parent, offset, size)
}

func NewFileService() (*FileService, error) {

	repo, err := mysql.NewFileRepository()
	if err != nil {
		return nil, err
	}

	return &FileService{Repo: repo}, nil
}
