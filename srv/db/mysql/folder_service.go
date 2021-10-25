package mysql

import (
	"drive"
	"fmt"
)

type FolderService struct {
	H *handle
}

func NewFolderService() (*FolderService, error) {
	h, err := getHandle()
	if err != nil {
		return nil, err
	}
	return &FolderService{H: h}, nil
}

func (fs *FolderService) Get(id int64) (*drive.Folder, error) {
	row := fs.H.Conn.QueryRow("SELECT `id`, `uuid`, `name`, `parent`, `attr` FROM `zd_files` WHERE `id` = ? LIMIT 1", id)
	f := &drive.Folder{}
	row.Scan(&f.Id, &f.Uuid, &f.Name, &f.Parent, &f.Attr)
	if f.Id == 0 {
		return nil, fmt.Errorf("no data")
	}
	return f, nil
}

func (fs *FolderService) GetByAlias(id string) (*drive.Folder, error) {
	row := fs.H.Conn.QueryRow("SELECT `id`, `uuid`, `name`, `parent`, `attr` FROM `zd_files` WHERE `uuid` = ? LIMIT 1", id)
	f := &drive.Folder{}
	row.Scan(&f.Id, &f.Uuid, &f.Name, &f.Parent, &f.Attr)
	if f.Id == 0 {
		return nil, fmt.Errorf("no data")
	}
	return f, nil
}

func (fs *FolderService) GetFolderPath(id string) ([]*drive.Folder, error) {

	folder, err := fs.GetByAlias(id)
	if err != nil {
		return nil, err
	}

	var result []*drive.Folder
	result = append(result, folder)
	pid := folder.Parent

	n := 0
	for {
		// loop safe check
		n++
		if n > 1000 {
			return nil, fmt.Errorf("loop error")
		}
		// get parent
		folder, err = fs.Get(pid)
		if err != nil || pid == 0 {
			break
		}
		result = append(result, folder)
		pid = folder.Parent
	}

	// reverse
	i := 0
	j := len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}

	return result, nil
}
