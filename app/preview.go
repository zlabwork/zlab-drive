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

const noPicture = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-exclamation-triangle" viewBox="0 0 16 16">
  <path d="M7.938 2.016A.13.13 0 0 1 8.002 2a.13.13 0 0 1 .063.016.146.146 0 0 1 .054.057l6.857 11.667c.036.06.035.124.002.183a.163.163 0 0 1-.054.06.116.116 0 0 1-.066.017H1.146a.115.115 0 0 1-.066-.017.163.163 0 0 1-.054-.06.176.176 0 0 1 .002-.183L7.884 2.073a.147.147 0 0 1 .054-.057zm1.044-.45a1.13 1.13 0 0 0-1.96 0L.165 13.233c-.457.778.091 1.767.98 1.767h13.713c.889 0 1.438-.99.98-1.767L8.982 1.566z"/>
  <path d="M7.002 12a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 5.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 5.995z"/>
</svg>`

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
				w.Header().Set("Content-Type", "image/svg+xml")
				w.Write([]byte(noPicture))
				return
			}
			path := utils.WorkDir(userId+"/data") + file.Path + file.Name
			bs, err := os.ReadFile(path)
			if err != nil {
				w.Header().Set("Content-Type", "image/svg+xml")
				w.Write([]byte(noPicture))
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
