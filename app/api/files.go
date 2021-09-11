package api

import (
	"drive"
	"drive/srv/db/mysql"
	"net/http"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	fs, err := mysql.NewFileService()
	data, err := fs.ListFiles(0, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		output := drive.JsonError{
			Code:    http.StatusInternalServerError,
			Message: "error",
			Refer:   "https://zlab.dev",
		}
		drive.ResponseJson(w, output)
		return
	}

	output := drive.JsonOK{
		Code: http.StatusOK,
		Data: data,
	}
	w.WriteHeader(http.StatusOK)
	drive.ResponseJson(w, output)
}
