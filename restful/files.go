package restful

import (
	"drive"
	"drive/msg"
	"drive/srv"
	"encoding/base64"
	"github.com/gorilla/mux"
	"net/http"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {

	// decode
	vars := mux.Vars(r)
	key := ""
	if len(vars["id"]) >= 2 {
		s, err := base64.RawURLEncoding.DecodeString(vars["id"])
		if err != nil {
			drive.ResponseJson(w, drive.JsonError{
				Code:    msg.ErrEncode,
				Message: msg.Text(msg.ErrEncode),
			})
			return
		}
		key = string(s)
	}

	// fs & fetch
	fs, err := srv.NewFileService()
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
