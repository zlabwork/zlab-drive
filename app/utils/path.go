package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultDirName = "zLabDrive/"
)

func HomeDir() string {
	if os.Getenv("DIR_HOME") == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Panicln(err)
		}
		return home + "/" + defaultDirName
	}
	return strings.TrimRight(os.Getenv("DIR_HOME"), "/") + "/"
}

func WorkDir(name string) string {
	dir := HomeDir() + strings.Trim(name, "/") + "/"
	if _, err := os.Stat(dir); err != nil {
		os.MkdirAll(dir, 0700)
	}
	return dir
}

func TempDir() string {
	dir := WorkDir("temp")
	return dir
}

func NewLogName(name string, version string) string {
	// 1610706262_2021_01_15_0.8.2-xx_data.txt
	now := time.Now()
	return strconv.FormatInt(now.Unix(), 10) + now.Format("_2006_01_02_") + version + "-" + name
}
