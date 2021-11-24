package adaptor

import "drive"

type OssDrive struct {
	Name string
}

func NewOssDrive() *OssDrive {
	return &OssDrive{
		Name: "AliYun OSS Drive",
	}
}

func (oss *OssDrive) Get(key string) (*drive.File, error) {
	return nil, nil
}

func (oss *OssDrive) List(key string, offset int, limit int) ([]*drive.File, error) {
	return nil, nil
}

func (oss *OssDrive) Create(file *drive.File) error {
	return nil
}

func (oss *OssDrive) Delete(key string) error {
	return nil
}

func (oss *OssDrive) Modify(key string, newFile *drive.File) error {
	return nil
}

func (oss *OssDrive) Bytes(file *drive.File) ([]byte, error) {
	return nil, nil
}
