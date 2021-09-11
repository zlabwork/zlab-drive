package mysql

import (
	"drive"
	"fmt"
)

type FileService struct {
	h *handle
}

func (f *FileService) ListFiles(parent int64, offset, number int) ([]*drive.File, error) {

	rows, err := f.h.Conn.Query("SELECT `id`, `uid`, `uuid`, `name`, `mime`, `size`, `hash`, `parent`, `path`, `attr`, `f_ctime`, `ctime`, `mtime` FROM `zd_files` WHERE `parent` = ? LIMIT 100", parent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*drive.File

	// scan
	for rows.Next() {
		file := &drive.File{}
		err := rows.Scan(&file.Id, &file.UserId, &file.Uuid, &file.Name, &file.MimeType, &file.Size, &file.Hash, &file.Parent, &file.Path, &file.Attr, &file.FileCtime, &file.Ctime, &file.Mtime)
		if err != nil {
			return nil, err
		}
		result = append(result, file)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (f *FileService) GetFile(id int64) (*drive.File, error) {
	row := f.h.Conn.QueryRow("SELECT `id`, `uid`, `uuid`, `name`, `mime`, `size`, `hash`, `parent`, `path`, `attr`, `f_ctime`, `ctime`, `mtime` FROM `zd_files` WHERE `id` = ? LIMIT 1", id)
	file := drive.File{}
	row.Scan(&file.Id, &file.UserId, &file.Uuid, &file.Name, &file.MimeType, &file.Size, &file.Hash, &file.Parent, &file.Path, &file.Attr, &file.FileCtime, &file.Ctime, &file.Mtime)
	if file.Id == 0 {
		return nil, fmt.Errorf("no data")
	}
	return &file, nil
}

func (f *FileService) CreateFile(file *drive.File) error {
	stmt, err := f.h.Conn.Prepare("INSERT INTO `zd_files` (`uid`, `uuid`, `name`, `mime`, `size`, `hash`, `parent`, `path`, `attr`, `f_ctime`, `ctime`, `mtime`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(file.UserId, file.Uuid, file.Name, file.MimeType, file.Size, file.Hash, file.Parent, file.Path, file.Attr, file.FileCtime, file.Ctime, file.Mtime)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileService) DeleteFile(id int64) error {
	stmt, err := f.h.Conn.Prepare("DELETE FROM `zd_files` WHERE id = ?")
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(id); err != nil {
		return err
	}
	return nil
}
