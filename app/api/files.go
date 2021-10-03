package api

import (
	"drive"
	"drive/app/msg"
	"drive/srv/db/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fs, err := mysql.NewFileService()
	if err != nil {
		drive.ResponseJson(w, drive.JsonError{
			Code:    msg.ErrDB,
			Message: msg.Text(msg.ErrDB),
			Refer:   "https://zlab.dev",
		})
		return
	}
	defer fs.H.Close()

	// fetch folder
	var parent int64
	if len(vars["id"]) >= 32 {
		folder, err := fs.FileAlias(vars["id"])
		if err != nil {
			drive.ResponseJson(w, drive.JsonError{
				Code:    msg.ErrNoData,
				Message: msg.Text(msg.ErrNoData),
			})
			return
		}
		if folder.MimeType != "folder" {
			drive.ResponseJson(w, drive.JsonError{
				Code:    msg.Err,
				Message: msg.Text(msg.Err),
			})
			return
		}
		parent = folder.Id
	}

	// fetch files list
	data, err := fs.Files(parent)
	if err != nil {
		drive.ResponseJson(w, drive.JsonError{
			Code:    msg.Err,
			Message: msg.Text(msg.Err),
		})
		return
	}

	output := drive.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
		Data:    data,
	}
	drive.ResponseJson(w, output)
}
