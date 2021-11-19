package api

import (
	"drive"
	"drive/app/msg"
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {

	drive.ResponseJson(w, drive.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
	})
}
