package srv

import "app"

type DbRepository interface {
	Get(id int64) (*app.File, error)
	GetByUUID(uuid string) (*app.File, error)
	Delete(id int64) error
	Create(file *app.File) (int64, error)
	Modify(file *app.File) error
	List(parent int64, offset int, size int) ([]*app.File, error)
}

type RepoService struct {
	Repo DbRepository
}

func (rs *RepoService) Get(id int64) (*app.File, error) {
	return rs.Repo.Get(id)
}

func (rs *RepoService) GetByUUID(uuid string) (*app.File, error) {
	return rs.Repo.GetByUUID(uuid)
}

func (rs *RepoService) Delete(id int64) error {
	return rs.Repo.Delete(id)
}

func (rs *RepoService) Create(file *app.File) (int64, error) {
	return rs.Repo.Create(file)
}

func (rs *RepoService) Modify(file *app.File) error {
	return rs.Repo.Modify(file)
}

func (rs *RepoService) List(parent int64, offset int, size int) ([]*app.File, error) {
	return rs.Repo.List(parent, offset, size)
}
