package mysql

import (
	"database/sql"
	"drive"
)

type FileService struct {
	H *handle
}

func (f *FileService) FileAlias(id string) (*drive.File, error) {
	row := f.H.Conn.QueryRow("SELECT `uuid`, `name`, `mime`, `size`, `hash`, `key`, `attr`, `f_ctime`, `ctime`, `mtime` FROM `zd_files` WHERE `uuid` = ? LIMIT 1", id)
	file := drive.File{}
	row.Scan(&file.Uuid, &file.Name, &file.MimeType, &file.Size, &file.Hash, &file.Key, &file.Attr, &file.FileCtime, &file.Ctime, &file.Mtime)
	return &file, nil
}

func (f *FileService) File(id int64) (*drive.File, error) {
	row := f.H.Conn.QueryRow("SELECT `uuid`, `name`, `mime`, `size`, `hash`, `key`, `attr`, `f_ctime`, `ctime`, `mtime` FROM `zd_files` WHERE `id` = ? LIMIT 1", id)
	file := drive.File{}
	row.Scan(&file.Uuid, &file.Name, &file.MimeType, &file.Size, &file.Hash, &file.Key, &file.Attr, &file.FileCtime, &file.Ctime, &file.Mtime)
	return &file, nil
}

func (f *FileService) Files(parent int64) ([]*drive.File, error) {

	rows, err := f.H.Conn.Query("SELECT `uuid`, `name`, `mime`, `size`, `hash`, `key`, `attr`, `f_ctime`, `ctime`, `mtime` FROM `zd_files` WHERE `parent` = ? ", parent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*drive.File

	// scan
	for rows.Next() {
		file := &drive.File{}
		err := rows.Scan(&file.Uuid, &file.Name, &file.MimeType, &file.Size, &file.Hash, &file.Key, &file.Attr, &file.FileCtime, &file.Ctime, &file.Mtime)
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

func (f *FileService) CreateFile(file *drive.File) (sql.Result, error) {
	stmt, err := f.H.Conn.Prepare("INSERT INTO `zd_files` (`uuid`, `name`, `mime`, `size`, `hash`, `key`, `attr`, `f_ctime`, `f_mtime`, `ctime`, `mtime`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(file.Uuid, file.Name, file.MimeType, file.Size, file.Hash, file.Key, file.Attr, file.FileCtime, file.FileMtime, file.Ctime, file.Mtime)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FileService) DeleteFile(id int64) error {
	stmt, err := f.H.Conn.Prepare("DELETE FROM `zd_files` WHERE id = ?")
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(id); err != nil {
		return err
	}
	return nil
}

func (f *FileService) ListFiles(parent int64, offset, size int) ([]*drive.File, error) {

	rows, err := f.H.Conn.Query("SELECT `uuid`, `name`, `mime`, `size`, `hash`, `key`, `attr`, `f_ctime`, `ctime`, `mtime` FROM `zd_files` WHERE `parent` = ? LIMIT ?,?", parent, offset, size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*drive.File

	// scan
	for rows.Next() {
		file := &drive.File{}
		err := rows.Scan(&file.Uuid, &file.Name, &file.MimeType, &file.Size, &file.Hash, &file.Key, &file.Attr, &file.FileCtime, &file.Ctime, &file.Mtime)
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
