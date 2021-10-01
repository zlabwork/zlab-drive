package app

import (
	"drive/app/utils"
	"drive/srv/db/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func PreviewHandler(w http.ResponseWriter, r *http.Request) {
	userId := "123456" // TODO :: modify userId
	id := mux.Vars(r)["id"]
	suf := "_100x100"
	temp := utils.WorkDir("temp/"+userId+"/"+id[0:1]) + id + suf

	// temp is not exist
	if _, err := os.Stat(temp); err != nil {
		if os.IsNotExist(err) {
			fs, err := mysql.NewFileService()
			if err != nil {
				return
			}
			file, err := fs.FileAlias(id)
			if err != nil {
				return
			}
			path := utils.WorkDir(userId+"/data") + file.Path + file.Name
			bs, err := os.ReadFile(path)
			if err != nil {
				return
			}
			if os.WriteFile(temp, bs, 0755) != nil {
				log.Println("error write")
				return
			}
			// TODO :: resize
			w.WriteHeader(http.StatusOK)
			w.Write(bs)
			return
		}
	}
	bs, err := os.ReadFile(temp)
	if err != nil {
		return
	}

	// TODO :: cache params
	tf := time.Now().AddDate(0, 0, 7).Format(http.TimeFormat)
	// w.Header().Set("Last-Modified", "")
	w.Header().Set("Cache-Control", "private, max-age=10800, pre-check=10800")
	w.Header().Set("Content-type", "image/jpeg")
	w.Header().Set("ETag", id)
	w.Header().Set("Expires", tf)
	w.Header().Set("Date", tf)
	w.Write(bs)
}
