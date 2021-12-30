package restful

import (
	"app"
	"app/msg"
	"app/srv"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func DoHandler(w http.ResponseWriter, r *http.Request) {

	// 1. param
	action := r.PostFormValue("action")
	vars := mux.Vars(r)
	key := strings.TrimSpace(vars["id"])

	// 2. adaptor
	fs, err := srv.NewFileService()
	if err != nil {
		app.ResponseJson(w, app.JsonError{
			Code:    msg.Err,
			Message: msg.Text(msg.Err),
		})
		return
	}

	// 3. action
	switch action {
	case "delete":
		if fs.Delete(key) != nil {
			app.ResponseJson(w, app.JsonError{
				Code:    msg.ErrProcess,
				Message: msg.Text(msg.ErrProcess),
			})
			return
		}
	case "move":
		fmt.Println("move")
	default:
		app.ResponseJson(w, app.JsonError{
			Code:    msg.ErrParameter,
			Message: msg.Text(msg.ErrParameter),
		})
		return
	}

	app.ResponseJson(w, app.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
	})
}
