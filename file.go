package drive

type File struct {
	Id        int64
	UserId    int64
	Uuid      string
	Name      string
	MimeType  string
	Size      int
	Hash      string
	Parent    int64
	Path      string
	Attr      string
	FileCtime int64
	Ctime     int64
	Mtime     int64
}

type FileService interface {
	File(id int64) (*File, error)
	Files() ([]*File, error)
	CreateFile(u *File) error
	DeleteFile(id int64) error
}
