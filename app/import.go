package app

import (
	"crypto/md5"
	"drive"
	"drive/app/utils"
	"drive/srv/db/mysql"
	"encoding/hex"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func ImportHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	name := vars.Get("name")
	filename := utils.WorkDir("123456/data") + name
	listFile(0, filename)
}

func listFile(parentId int64, filename string) error {
	l := len(utils.WorkDir("123456/data"))
	path := filename[l:]

	if _, err := os.Stat(filename); err != nil {
		return err
	}
	ls, err := os.ReadDir(filename)
	if err != nil {
		return err
	}

	// todo mysql conn
	fs, err := mysql.NewFileService()
	if err != nil {
		return err
	}
	defer fs.H.Close()

	for _, item := range ls {
		// 忽略隐藏文件
		if strings.HasPrefix(item.Name(), ".") {
			continue
		}
		ts := time.Now().Unix()
		if item.IsDir() {
			file := &drive.File{
				UserId:    123456, // todo
				Uuid:      uuid.New().String(),
				Name:      item.Name(),
				MimeType:  "folder",
				Size:      0,
				Hash:      "",
				Parent:    parentId,
				Path:      path + "/",
				Attr:      "",
				FileCtime: ts,
				FileMtime: ts,
				Ctime:     ts,
				Mtime:     ts,
			}
			res, err := fs.CreateFile(file)
			if err != nil {
				log.Println(err)
			}
			id, _ := res.LastInsertId()
			listFile(id, filename+"/"+item.Name())
		} else {
			info, err := item.Info()
			h := md5.New()
			bs, err := os.ReadFile(filename + "/" + item.Name())
			if err != nil {
				log.Println(err)
				return err
			}
			h.Write(bs)
			if err != nil {
				log.Println(err)
			}

			file := &drive.File{
				UserId:    123456, // todo
				Uuid:      uuid.New().String(),
				Name:      item.Name(),
				MimeType:  "",
				Size:      int(info.Size()),
				Hash:      hex.EncodeToString(h.Sum(nil)),
				Parent:    parentId,
				Path:      path + "/",
				Attr:      "",
				FileCtime: info.ModTime().Unix(),
				FileMtime: info.ModTime().Unix(),
				Ctime:     ts,
				Mtime:     ts,
			}
			fs.CreateFile(file)
		}
	}
	return nil
}
