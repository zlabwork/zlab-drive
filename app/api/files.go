package api

import (
	"drive"
	"drive/app/msg"
	"drive/srv/db/mysql"
	"net/http"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	fs, err := mysql.NewFileService()
	data, err := fs.ListFiles(0, 0, 0)
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
