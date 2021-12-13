package web

import (
	"net/http"
	"os"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	bs, _ := os.ReadFile("../template/main/react.html")
	w.Write(bs)
}
