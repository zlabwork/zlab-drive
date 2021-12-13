package fs

import "drive"

type S3Drive struct {
	Name string
}

func NewS3Drive() *S3Drive {
	return &S3Drive{
		Name: "Aws S3 Drive",
	}
}

func (s3 *S3Drive) Get(key string) (*drive.File, error) {
	return nil, nil
}

func (s3 *S3Drive) List(key string, offset int, limit int) ([]*drive.File, error) {
	return nil, nil
}

func (s3 *S3Drive) Create(file *drive.File) error {
	return nil
}

func (s3 *S3Drive) Delete(key string) error {
	return nil
}

func (s3 *S3Drive) Modify(key string, newFile *drive.File) error {
	return nil
}

func (s3 *S3Drive) Bytes(file *drive.File) ([]byte, error) {
	return nil, nil
}
