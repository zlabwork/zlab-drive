package restful

import (
	"app"
	"app/msg"
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {

	app.ResponseJson(w, app.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
	})
}
