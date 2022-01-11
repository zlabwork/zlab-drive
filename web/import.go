package web

import (
	"app"
	"app/msg"
	"app/srv/repository/mysql"
	"app/utils"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ImportHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	name := vars.Get("name")
	uid := r.Context().Value(app.UserIdKey).(int64)

	// checking
	if uid == 0 || name == "" {
		app.Response(w, msg.ErrParameter)
		return
	}

	base := utils.WorkDir(strconv.FormatInt(uid, 10) + string(os.PathSeparator) + "data")
	err := listFile(0, base, name)
	if err != nil {
		log.Println(err)
		app.Response(w, msg.ErrProcess)
	}
}

func listFile(parentId int64, base, pathName string) error {

	filename := base + string(os.PathSeparator) + pathName
	if _, err := os.Stat(filename); err != nil {
		return err
	}
	ls, err := os.ReadDir(filename)
	if err != nil {
		return err
	}

	// mysql conn
	fs, err := mysql.NewFileRepository()
	if err != nil {
		return err
	}
	defer fs.H.Close()

	for _, item := range ls {
		// 忽略隐藏文件
		if strings.HasPrefix(item.Name(), ".") {
			continue
		}
		tn := time.Now()

		// if folder
		var hash, mimeType string
		if item.IsDir() {
			mimeType = "dir"
		} else {
			mimeType = "image/jpeg" // TODO: 识别文件类型
			bs, err := os.ReadFile(filename + "/" + item.Name())
			if err != nil {
				return err
			}
			h := md5.New()
			h.Write(bs)
			hash = hex.EncodeToString(h.Sum(nil))
		}
		info, err := item.Info()
		if err != nil {
			return err
		}

		// FIXME: UserId & ParentId
		newPath := pathName + "/" + item.Name()
		file := &app.File{
			// UserId:    123456,
			// Parent:    parentId,
			Uuid:      uuid.New().String(),
			Name:      item.Name(),
			MimeType:  mimeType,
			Size:      info.Size(),
			Hash:      hash,
			Key:       base64.RawURLEncoding.EncodeToString([]byte(newPath)),
			Attr:      "",
			FileCtime: info.ModTime(),
			FileMtime: info.ModTime(),
			Ctime:     tn,
			Mtime:     tn,
		}

		id, err := fs.Create(file)
		if err != nil {
			return err
		}
		if item.IsDir() {
			listFile(id, base, newPath)
		}
	}
	return nil
}
