package drive

type File struct {
	Uuid      string `json:"uuid"`
	MimeType  string `json:"mime"`
	Size      int64  `json:"size"`
	Hash      string `json:"hash"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	Attr      string `json:"attr"`
	FileCtime int64  `json:"file_ctime"`
	FileMtime int64  `json:"file_mtime"`
	Ctime     int64  `json:"ctime"`
	Mtime     int64  `json:"mtime"`
}

type Attr struct {
	Favorite bool
	Color    string
	Width    int
	Height   int
	Duration int
}

const (
	ImageSizeDefault = iota
	ImageSizeSmall
	ImageSizeMiddle
	ImageSizeLarge
)
