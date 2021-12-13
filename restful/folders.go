package restful

import (
	"drive"
	"drive/msg"
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {

	drive.ResponseJson(w, drive.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
	})
}
