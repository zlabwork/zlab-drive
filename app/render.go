package app

import (
	"html/template"
	"io"
)

func templateRender(w io.Writer, tmpl string, data interface{}) {
	t, _ := template.ParseFiles("../templates/"+tmpl, "../templates/parts/header.html", "../templates/parts/footer.html", "../templates/parts/sidebar.html")
	t.Execute(w, data)
}
