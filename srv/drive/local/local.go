package local

import (
	"app"
	"app/utils"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Drive struct {
	namespace string
	Name      string
}

func NewDrive(namespace string) *Drive {
	return &Drive{
		namespace: namespace,
		Name:      "Local File Drive",
	}
}

func (loc *Drive) Get(key string) (*app.File, error) {

	dirName := loc.getDirData()
	f, err := os.Stat(dirName + key)
	if err != nil {
		return nil, err
	}

	return &app.File{
		Mime:      "",
		Hash:      "",
		Key:       key,
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
	dirName := loc.getDirData()
	files, _ := ioutil.ReadDir(dirName + key)
	for _, f := range files {
		mime := ""
		if f.IsDir() {
			mime = "folder"
		}
		obj := &app.File{
			Mime:      mime,
			Hash:      "",
			Key:       key + "/" + f.Name(),
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
	dirName := loc.getDirData()
	trash := loc.getDirTrash()
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

	dirName := loc.getDirData()
	bs, err := os.ReadFile(dirName + file.Key)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (loc *Drive) getDirData() string {
	return utils.WorkDir(loc.namespace + "/data/")
}

func (loc *Drive) getDirTrash() string {
	return utils.WorkDir(loc.namespace + "/trash/")
}
