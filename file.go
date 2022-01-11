package app

import "time"

type File struct {
	Uuid      string
	MimeType  string
	Size      int64
	Hash      string
	Name      string
	Key       string
	Attr      string
	FileCtime time.Time
	FileMtime time.Time
	Ctime     time.Time
	Mtime     time.Time
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
