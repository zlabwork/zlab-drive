package adaptor

import (
	"drive"
	"drive/app/utils"
	"encoding/base64"
	"io/ioutil"
	"os"
	"strconv"
	"time"
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

	dirName := utils.WorkDir("data")
	f, err := os.Stat(dirName + key)
	if err != nil {
		return nil, err
	}

	return &drive.File{
		MimeType:  "",
		Hash:      "",
		Key:       base64.RawURLEncoding.EncodeToString([]byte(key)),
		Name:      f.Name(),
		Size:      f.Size(),
		FileMtime: f.ModTime().Unix(), // TODO: file time
	}, nil
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

func (loc *LocalDrive) Modify(key string, newFile *drive.File) error {
	return nil
}

func (loc *LocalDrive) Bytes(file *drive.File) ([]byte, error) {
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
