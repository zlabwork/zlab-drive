package api

import (
	"drive"
	"drive/app/msg"
	"drive/srv/db/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {

	// service
	fs, err := mysql.NewFolderService()
	if err != nil {
		drive.ResponseJson(w, drive.JsonError{
			Code:    msg.ErrDB,
			Message: msg.Text(msg.ErrDB),
			Refer:   "https://zlab.dev",
		})
		return
	}
	defer fs.H.Close()

	home := &drive.Folder{
		Id:   0,
		Name: "Home",
		Uuid: "0",
	}

	// check parameter
	vars := mux.Vars(r)
	if len(vars["id"]) < 32 {
		drive.ResponseJson(w, drive.JsonOK{
			Code:    msg.OK,
			Message: msg.Text(msg.OK),
			Data:    []*drive.Folder{home},
		})
		return
	}

	// find data
	folders, err := fs.GetFolderPath(vars["id"])
	if err != nil {
		drive.ResponseJson(w, drive.JsonError{
			Code:    msg.ErrNoData,
			Message: msg.Text(msg.ErrNoData),
		})
		return
	}
	folders = append([]*drive.Folder{home}, folders...)

	drive.ResponseJson(w, drive.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
		Data:    folders,
	})
}
