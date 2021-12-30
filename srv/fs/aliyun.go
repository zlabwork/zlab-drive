package fs

import "app"

type OssDrive struct {
	Name string
}

func NewOssDrive() *OssDrive {
	return &OssDrive{
		Name: "AliYun OSS Drive",
	}
}

func (oss *OssDrive) Get(key string) (*app.File, error) {
	return nil, nil
}

func (oss *OssDrive) List(key string, offset int, limit int) ([]*app.File, error) {
	return nil, nil
}

func (oss *OssDrive) Create(file *app.File) error {
	return nil
}

func (oss *OssDrive) Delete(key string) error {
	return nil
}

func (oss *OssDrive) Modify(key string, newFile *app.File) error {
	return nil
}

func (oss *OssDrive) Bytes(file *app.File) ([]byte, error) {
	return nil, nil
}
