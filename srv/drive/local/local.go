package local

import (
	"app"
	"app/utils"
	"encoding/base64"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Drive struct {
	Name string
}

func NewDrive() *Drive {
	return &Drive{
		Name: "Local File Drive",
	}
}

func (loc *Drive) Get(key string) (*app.File, error) {

	dirName := utils.WorkDir("data")
	f, err := os.Stat(dirName + key)
	if err != nil {
		return nil, err
	}

	return &app.File{
		MimeType:  "",
		Hash:      "",
		Key:       base64.RawURLEncoding.EncodeToString([]byte(key)),
		Name:      f.Name(),
		Size:      f.Size(),
		FileMtime: f.ModTime(), // TODO: file time
	}, nil
}

func (loc *Drive) List(key string, offset int, limit int) ([]*app.File, error) {

	if key == "/" {
		key = ""
	}

	var fs []*app.File
	dirName := utils.WorkDir("data")
	files, _ := ioutil.ReadDir(dirName + key)
	for _, f := range files {
		mime := ""
		if f.IsDir() {
			mime = "folder"
		}
		obj := &app.File{
			MimeType:  mime,
			Hash:      "",
			Key:       base64.RawURLEncoding.EncodeToString([]byte(key + "/" + f.Name())),
			Name:      f.Name(),
			Size:      f.Size(),
			FileMtime: f.ModTime(), // TODO: https://github.com/djherbis/times
		}
		fs = append(fs, obj)
	}
	return fs, nil
}

func (loc *Drive) Create(file *app.File) error {
	return nil
}

func (loc *Drive) Delete(key string) error {

	// TODO: clear cache
	dirName := utils.WorkDir("data")
	trash := utils.WorkDir("trash")
	path1 := dirName + key
	path2 := trash + key + strconv.FormatInt(time.Now().UnixNano(), 10)

	//err := os.RemoveAll(dirName + key)
	err := os.Rename(path1, path2)
	if err != nil {
		return err
	}
	return nil
}

func (loc *Drive) Modify(key string, newFile *app.File) error {
	return nil
}

func (loc *Drive) Bytes(file *app.File) ([]byte, error) {
	k, err := base64.RawURLEncoding.DecodeString(file.Key)
	if err != nil {
		return nil, err
	}
	dirName := utils.WorkDir("data")
	bs, err := os.ReadFile(dirName + string(k))
	if err != nil {
		return nil, err
	}
	return bs, nil
}
