package app

import "time"

type File struct {
	Uuid      string    `json:"uuid"`
	MimeType  string    `json:"mime"`
	Size      int64     `json:"size"`
	Hash      string    `json:"hash"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
	Attr      string    `json:"attr"`
	FileCtime time.Time `json:"file_ctime"`
	FileMtime time.Time `json:"file_mtime"`
	Ctime     time.Time `json:"ctime"`
	Mtime     time.Time `json:"mtime"`
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
