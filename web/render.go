package web

import (
	"html/template"
	"io"
)

func templateRender(w io.Writer, tmpl string, data interface{}) {
	t, _ := template.ParseFiles("../template/"+tmpl, "../template/parts/header.html", "../template/parts/footer.html", "../template/parts/sidebar.html")
	t.Execute(w, data)
}

func renderDrive(w io.Writer, tmpl string, data interface{}) {
	t, _ := template.ParseFiles("../template/"+tmpl, "../template/parts/drive_header.html", "../template/parts/drive_footer.html", "../template/parts/drive_sidebar.html", "../template/parts/drive_setting.html")
	t.Execute(w, data)
}
