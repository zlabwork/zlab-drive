package app

import (
	"app/msg"
	"encoding/json"
	"net/http"
)

type JsonOK struct {
	Code    int
	Message string
	Data    interface{}
}

type JsonError struct {
	Code    int
	Message string
	Refer   string
}

func ResponseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	bs, _ := json.Marshal(data)
	w.Write(bs)
}

func Response(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	bs, _ := json.Marshal(JsonError{
		Code:    code,
		Message: msg.Text(code),
	})
	w.Write(bs)
}
