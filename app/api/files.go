package api

import (
	"drive"
	"drive/app/msg"
	"drive/srv/adaptor"
	"encoding/base64"
	"github.com/gorilla/mux"
	"net/http"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {

	// decode
	vars := mux.Vars(r)
	s, err := base64.RawURLEncoding.DecodeString(vars["id"])
	if err != nil {
		drive.ResponseJson(w, drive.JsonError{
			Code:    msg.ErrEncode,
			Message: msg.Text(msg.ErrEncode),
		})
		return
	}
	key := string(s)

	// fs & fetch
	fs, err := adaptor.NewAdaptor()
	if err != nil {
		drive.ResponseJson(w, drive.JsonError{
			Code:    msg.Err,
			Message: msg.Text(msg.Err),
		})
		return
	}
	data, err := fs.List(key, 0, 20) // TODO: offset & limit
	if err != nil {
		drive.ResponseJson(w, drive.JsonError{
			Code:    msg.ErrNoData,
			Message: msg.Text(msg.ErrNoData),
		})
		return
	}

	// output
	drive.ResponseJson(w, drive.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
		Data:    data,
	})

}
