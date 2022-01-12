package app

import "time"

type File struct {
	Id        int64
	Uuid      string
	Mime      string
	Size      int64
	Hash      string
	Parent    int64
	Key       string
	Name      string
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
