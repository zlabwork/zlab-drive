package drive

type File struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"uid"`
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	MimeType  string `json:"mime"`
	Size      int    `json:"size"`
	Hash      string `json:"hash"`
	Parent    int64  `json:"parent"`
	Path      string `json:"path"`
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

type FileService interface {
	FileAlias(id string) (*File, error)
	File(id int64) (*File, error)
	Files(parent int64) ([]*File, error)
	CreateFile(u *File) error
	DeleteFile(id int64) error
}
