package oss

import "app"

type Drive struct {
	namespace string
	Name      string
}

func NewDrive(namespace string) *Drive {
	return &Drive{
		namespace: namespace,
		Name:      "AliYun OSS Drive",
	}
}

func (oss *Drive) Get(key string) (*app.File, error) {
	return nil, nil
}

func (oss *Drive) List(key string, offset int, limit int) ([]*app.File, error) {
	return nil, nil
}

func (oss *Drive) Create(file *app.File) error {
	return nil
}

func (oss *Drive) Delete(key string) error {
	return nil
}

func (oss *Drive) Modify(key string, newFile *app.File) error {
	return nil
}

func (oss *Drive) Bytes(file *app.File) ([]byte, error) {
	return nil, nil
}
