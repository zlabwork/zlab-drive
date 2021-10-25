package drive

type Folder struct {
	Id     int64  `json:"id"`
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	Parent int64  `json:"parent"`
	Attr   string `json:"attr"`
}

type FolderService interface {
	Get(id int64) (*Folder, error)
	GetByAlias(id string) (*Folder, error)
	GetFolderPath(id string) ([]*Folder, error)
}
