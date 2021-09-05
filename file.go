package drive

type File struct {
	Id        int64
	Name      string
	MimeType  string
	UserId    int64
	Size      int
	Hash      string
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
