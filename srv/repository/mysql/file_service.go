package mysql

import (
	"app"
	"fmt"
)

type FileRepository struct {
	H *handle
}

func (f *FileRepository) Get(id int64) (*app.File, error) {

	row := f.H.Conn.QueryRow("SELECT `id`, `uuid`, `mime`, `size`, `parent`, `hash`, `key`, `name`, `attr`, `f_ctime`, `f_mtime`, `ctime`, `mtime` FROM `zd_files` WHERE `id` = ? LIMIT 1", id)
	file := app.File{}
	row.Scan(&file.Id, &file.Uuid, &file.Mime, &file.Size, &file.Parent, &file.Hash, &file.Key, &file.Name, &file.Attr, &file.FileCtime, &file.FileMtime, &file.Ctime, &file.Mtime)
	return &file, nil
}

func (f *FileRepository) GetByUUID(uuid string) (*app.File, error) {

	row := f.H.Conn.QueryRow("SELECT `id`, `uuid`, `mime`, `size`, `parent`, `hash`, `key`, `name`, `attr`, `f_ctime`, `f_mtime`, `ctime`, `mtime` FROM `zd_files` WHERE `uuid` = ? LIMIT 1", uuid)
	file := app.File{}
	row.Scan(&file.Id, &file.Uuid, &file.Mime, &file.Size, &file.Parent, &file.Hash, &file.Key, &file.Name, &file.Attr, &file.FileCtime, &file.FileMtime, &file.Ctime, &file.Mtime)
	return &file, nil
}

func (f *FileRepository) Delete(id int64) error {

	stmt, err := f.H.Conn.Prepare("DELETE FROM `zd_files` WHERE id = ?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (f *FileRepository) Create(file *app.File) (int64, error) {

	stmt, err := f.H.Conn.Prepare("INSERT INTO `zd_files` (`uuid`, `mime`, `size`, `parent`, `hash`, `key`, `name`, `attr`, `f_ctime`, `f_mtime`, `ctime`, `mtime`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(file.Uuid, file.Mime, file.Size, file.Parent, file.Hash, file.Key, file.Name, file.Attr, file.FileCtime, file.FileMtime, file.Ctime, file.Mtime)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (f *FileRepository) Modify(file *app.File) error {

	return fmt.Errorf("TODO implement")
}

func (f *FileRepository) List(parent int64, offset int, size int) ([]*app.File, error) {

	rows, err := f.H.Conn.Query("SELECT `id`, `uuid`, `mime`, `size`, `parent`, `hash`, `key`, `name`, `attr`, `f_ctime`, `f_mtime`, `ctime`, `mtime` FROM `zd_files` WHERE `parent` = ? LIMIT ?,?", parent, offset, size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*app.File

	// scan
	for rows.Next() {
		file := &app.File{}
		err := rows.Scan(&file.Id, &file.Uuid, &file.Mime, &file.Size, &file.Parent, &file.Hash, &file.Key, &file.Name, &file.Attr, &file.FileCtime, &file.FileMtime, &file.Ctime, &file.Mtime)
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
