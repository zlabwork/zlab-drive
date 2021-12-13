package web

import (
	"bytes"
	"crypto/md5"
	"drive"
	"drive/msg"
	"drive/srv"
	"drive/utils"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/disintegration/gift"
	"github.com/gorilla/mux"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const noPicture = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-exclamation-triangle" viewBox="0 0 16 16">
  <path d="M7.938 2.016A.13.13 0 0 1 8.002 2a.13.13 0 0 1 .063.016.146.146 0 0 1 .054.057l6.857 11.667c.036.06.035.124.002.183a.163.163 0 0 1-.054.06.116.116 0 0 1-.066.017H1.146a.115.115 0 0 1-.066-.017.163.163 0 0 1-.054-.06.176.176 0 0 1 .002-.183L7.884 2.073a.147.147 0 0 1 .054-.057zm1.044-.45a1.13 1.13 0 0 0-1.96 0L.165 13.233c-.457.778.091 1.767.98 1.767h13.713c.889 0 1.438-.99.98-1.767L8.982 1.566z"/>
  <path d="M7.002 12a1 1 0 1 1 2 0 1 1 0 0 1-2 0zM7.1 5.995a.905.905 0 1 1 1.8 0l-.35 3.507a.552.552 0 0 1-1.1 0L7.1 5.995z"/>
</svg>`

// PreviewHandler TODO :: modify userId
func PreviewHandler(w http.ResponseWriter, r *http.Request) {

	// 1. key & hash
	id := mux.Vars(r)["id"]
	s, err := base64.RawURLEncoding.DecodeString(id)
	if err != nil {
		drive.ResponseJson(w, drive.JsonError{
			Code:    msg.ErrEncode,
			Message: msg.Text(msg.ErrEncode),
		})
		return
	}
	key := string(s)
	h := md5.New()
	h.Write(s)
	ha := h.Sum(nil)
	haStr := hex.EncodeToString(ha)

	// 2. width & height
	vars := r.URL.Query()
	name := strings.ToLower(vars.Get("name"))

	var width, height int
	var size string
	if name == "large" {
		width = drive.Cfg.Image.Large.Width
		height = drive.Cfg.Image.Large.Height
		size = "large"
	} else {
		width = drive.Cfg.Image.Thumb.Width
		height = drive.Cfg.Image.Thumb.Height
		size = "small"
	}

	// 3. cache name
	suf := fmt.Sprintf("_%dx%d", width, height)
	temp := utils.WorkDir("temp/"+os.Getenv("APP_DRIVE")+haStr[0:2]) + string(os.PathSeparator) + haStr + suf

	// 4. temp is not exist
	if _, err := os.Stat(temp); err != nil {
		if os.IsNotExist(err) {
			// fetch from adaptor
			fs, err := srv.NewFileService()
			if err != nil {
				return
			}
			// 1>. fetch
			file, err := fs.Get(key)
			if err != nil {
				w.Header().Set("Content-Type", "image/svg+xml")
				w.Write([]byte(noPicture))
				return
			}
			bs, err := fs.Bytes(file)
			if err != nil {
				w.Header().Set("Content-Type", "image/svg+xml")
				w.Write([]byte(noPicture))
				return
			}

			// 2>. create image.Image
			rd := bytes.NewReader(bs)
			if err != nil {
				w.Header().Set("Content-Type", "image/svg+xml")
				w.Write([]byte(noPicture))
				return
			}
			srcImg, _, err := image.Decode(rd)
			if err != nil {
				w.Header().Set("Content-Type", "image/svg+xml")
				w.Write([]byte(noPicture))
				return
			}

			// 3>. resize & save
			dst := resizeImage(srcImg, size)
			err = saveImage(temp, dst)
			if err != nil {
				log.Println("error when save image", err)
				return
			}
		}
	}
	bs, err := os.ReadFile(temp)
	if err != nil {
		return
	}

	// TODO :: cache params
	tf := time.Now().AddDate(0, 0, 7).Format(http.TimeFormat)
	// w.Header().Set("Last-Modified", "")
	w.Header().Set("Cache-Control", "private, max-age=10800, pre-check=10800")
	w.Header().Set("Content-type", "image/png")
	w.Header().Set("ETag", id)
	w.Header().Set("Expires", tf)
	w.Header().Set("Date", tf)
	w.Write(bs)
}

// @docs https://github.com/disintegration/gift
func resizeImage(src image.Image, s string) *image.RGBA {
	// 1. Create a new filter list and add some filters.
	var w, h int
	var g *gift.GIFT
	if s == "small" {
		w = drive.Cfg.Image.Thumb.Width
		h = drive.Cfg.Image.Thumb.Height
		g = gift.New(
			gift.ResizeToFill(w, h, gift.LanczosResampling, gift.CenterAnchor),
		)
	} else if s == "large" {
		w = drive.Cfg.Image.Large.Width
		h = drive.Cfg.Image.Large.Height
		g = gift.New(
			gift.ResizeToFit(w, h, gift.LanczosResampling),
		)
	}

	// 2. Create a new image of the corresponding size.
	// dst is a new target image, src is the original image.
	dst := image.NewRGBA(g.Bounds(src.Bounds()))

	// 3. Use the Draw func to apply the filters to src and store the result in dst.
	g.Draw(dst, src)
	return dst
}

func saveImage(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		return err
	}
	return nil
}