package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	uri := mux.Vars(r)

	switch uri["version"] {

	case "v1":
		w.WriteHeader(http.StatusOK)
		bs, _ := os.ReadFile("../template/v1/react.html")
		w.Write(bs)

	case "v2":
		w.WriteHeader(http.StatusOK)
		type data struct {
			Title string
		}
		vars := &data{Title: "zlab drive"}
		renderDrive(w, "v2", "main/index.html", vars)

	}

}
