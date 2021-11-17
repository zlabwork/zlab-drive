package adaptor

import (
	"drive"
	"fmt"
	"os"
	"strconv"
)

// NewAdaptor TODO :: modify return interface
func NewAdaptor() (interface{}, error) {

	appDrive, err := strconv.Atoi(os.Getenv("APP_DRIVE"))
	if err != nil {
		return nil, err
	}

	switch appDrive {

	case drive.LocalDrive:
		return NewLocalDrive(), nil

	case drive.S3Drive:
		return NewS3Drive(), nil

	case drive.OssDrive:
		return NewOssDrive(), nil

	}

	return nil, fmt.Errorf("error drive")

}
