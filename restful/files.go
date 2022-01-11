package restful

import (
	"app"
	"app/msg"
	"app/srv"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {

	// decode
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		app.Response(w, msg.ErrParameter)
		return
	}

	// fs & fetch
	fs, err := srv.NewFileService()
	if err != nil {
		app.ResponseJson(w, app.JsonError{
			Code:    msg.Err,
			Message: msg.Text(msg.Err),
		})
		return
	}
	data, err := fs.List(id, 0, 20) // TODO: offset & limit
	if err != nil {
		app.ResponseJson(w, app.JsonError{
			Code:    msg.ErrNoData,
			Message: msg.Text(msg.ErrNoData),
		})
		return
	}

	// output
	app.ResponseJson(w, app.JsonOK{
		Code:    msg.OK,
		Message: msg.Text(msg.OK),
		Data:    data,
	})

}
