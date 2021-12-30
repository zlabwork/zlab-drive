package mysql

import (
	"app"
	"fmt"
)

func getHandle() (*handle, error) {
	c := app.Cfg.Db.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci", c.User, c.Pass, c.Host, c.Port, c.Name)
	return ConnectMySQL(dsn)
}

func NewFileService() (*FileService, error) {
	h, err := getHandle()
	if err != nil {
		return nil, err
	}
	return &FileService{H: h}, nil
}
