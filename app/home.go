package app

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	type data struct {
		Title string
	}
	vars := &data{Title: "zlab drive"}
	renderDrive(w, "main/index.html", vars)
}
