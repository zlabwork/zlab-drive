package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func Placeholder(w http.ResponseWriter, r *http.Request) {
	// size
	uris := mux.Vars(r)
	if !strings.Contains(uris["id"], "x") {
		uris["id"] = "200x200"
	}
	size := strings.Split(uris["id"], "x")
	width := size[0]
	height := size[1]

	// background
	background := "#" + r.FormValue("bg")
	if background == "#" {
		background = "#17223B"
	}

	// color
	color := "#" + r.FormValue("color")
	if color == "#" {
		color = "#FFFFFF"
	}

	// text
	text := r.FormValue("text")
	if text == "" {
		text = width + "x" + height
	}

	// font size
	wInt, _ := strconv.Atoi(width)
	fontSize := "24px"
	if wInt < 80 {
		fontSize = "16px"
	}

	textTag := `<text x='50%' y='50%' style='dominant-baseline:middle;text-anchor:middle;font-size:` + fontSize + `' fill='` + color + `'>` + text + `</text>`
	str := fmt.Sprintf(`<svg xmlns='http://www.w3.org/2000/svg' width='%s' height='%s'><rect x='0' y='0' width='%s' height='%s' fill='%s'/>%s</svg>`, width, height, width, height, background, textTag)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(str))
}
