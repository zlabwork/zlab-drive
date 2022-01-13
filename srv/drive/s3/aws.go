package s3

import "app"

type Drive struct {
	namespace string
	Name      string
}

func NewDrive(namespace string) *Drive {
	return &Drive{
		namespace: namespace,
		Name:      "Aws S3 Drive",
	}
}

func (s3 *Drive) Get(key string) (*app.File, error) {
	return nil, nil
}

func (s3 *Drive) List(key string, offset int, limit int) ([]*app.File, error) {
	return nil, nil
}

func (s3 *Drive) Create(file *app.File) error {
	return nil
}

func (s3 *Drive) Delete(key string) error {
	return nil
}

func (s3 *Drive) Modify(key string, newFile *app.File) error {
	return nil
}

func (s3 *Drive) Bytes(file *app.File) ([]byte, error) {
	return nil, nil
}
