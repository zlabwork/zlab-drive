package web

import (
	"html/template"
	"io"
)

func templateRender(w io.Writer, tmpl string, data interface{}) {
	t, _ := template.ParseFiles("../template/"+tmpl, "../template/parts/header.html", "../template/parts/footer.html", "../template/parts/sidebar.html")
	t.Execute(w, data)
}

func renderDrive(w io.Writer, version, tmpl string, data interface{}) {
	t, _ := template.ParseFiles(
		"../template/"+version+"/"+tmpl,
		"../template/"+version+"/parts/header.html",
		"../template/"+version+"/parts/footer.html",
		"../template/"+version+"/parts/sidebar.html",
		"../template/"+version+"/parts/setting.html",
	)
	t.Execute(w, data)
}
