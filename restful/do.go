package restful

import (
	"drive"
	"drive/msg"
	"drive/srv"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func DoHandler(w http.ResponseWriter, r *http.Request) {

	// 1. param
	action := r.PostFormValue("action")
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

	// 2. adaptor
	fs, err := srv.NewFileService()
	if err != nil {
		drive.ResponseJson(w, drive.JsonError{
			Code:    msg.Err,
			Message: msg.Text(msg.Err),
		})
		return
	}

	// 3. action
	switch action {
	case "delete":
		if fs.Delete(key) != nil {
			drive.ResponseJson(w, drive.JsonError{
				Code:    msg.ErrProcess,
				Message: msg.Text(msg.ErrProcess),
			})
			return
		}
	case "move":
		fmt.Println("move")
	default:
		drive.ResponseJson(w, drive.JsonError{
			Code:    msg.ErrParameter,
			Message: msg.Text(msg.ErrParameter),
		})
		return
	}

	drive.ResponseJson(w, drive.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
	})
}
