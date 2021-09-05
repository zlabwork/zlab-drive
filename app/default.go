package app

import "net/http"

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	type data struct {
		Title string
	}
	vars := &data{Title: "zlab drive"}
	templateRender(w, "index.html", vars)
}
