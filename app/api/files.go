package api

import (
	"drive"
	"drive/app/msg"
	"drive/srv/db/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	uris := mux.Vars(r)
	vars := r.URL.Query()

	size := 50
	parent, _ := strconv.ParseInt(uris["id"], 10, 64)
	offset, _ := strconv.Atoi(vars.Get("offset"))

	fs, err := mysql.NewFileService()
	if err != nil {
		output := drive.JsonError{
			Code:    msg.ErrDB,
			Message: msg.Text(msg.ErrDB),
			Refer:   "https://zlab.dev",
		}
		drive.ResponseJson(w, output)
		return
	}
	data, err := fs.ListFiles(parent, offset, size)
	if err != nil {
		output := drive.JsonError{
			Code:    msg.Err,
			Message: msg.Text(msg.Err),
			Refer:   "https://zlab.dev",
		}
		drive.ResponseJson(w, output)
		return
	}

	output := drive.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
		Data:    data,
	}
	drive.ResponseJson(w, output)
}
