package adaptor

import (
	"drive"
	"drive/app/utils"
	"encoding/base64"
	"io/ioutil"
)

type LocalDrive struct {
	Name string
}

func NewLocalDrive() *LocalDrive {
	return &LocalDrive{
		Name: "Local File Drive",
	}
}

func (loc *LocalDrive) Get(key string) (*drive.File, error) {
	return nil, nil
}

func (loc *LocalDrive) List(key string, offset int, limit int) ([]*drive.File, error) {

	if key == "/" {
		key = ""
	}

	var fs []*drive.File
	dirName := utils.WorkDir("data")
	files, _ := ioutil.ReadDir(dirName + key)
	for _, f := range files {
		mime := ""
		if f.IsDir() {
			mime = "folder"
		}
		obj := &drive.File{
			MimeType:  mime,
			Hash:      "",
			Key:       base64.RawURLEncoding.EncodeToString([]byte(key + "/" + f.Name())),
			Name:      f.Name(),
			Size:      f.Size(),
			FileMtime: f.ModTime().Unix(), // TODO: https://github.com/djherbis/times
		}
		fs = append(fs, obj)
	}
	return fs, nil
}

func (loc *LocalDrive) Create(file *drive.File) error {
	return nil
}

func (loc *LocalDrive) Delete(key string) error {
	return nil
}

func (loc *LocalDrive) Modify(key string, newFile *drive.File) error {
	return nil
}
