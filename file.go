package github.com/tolgaerdonmez/goutils

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// CheckFile checks if file exists returns "dir"|"file"|""
func CheckFile(filename string) string {
	info, err := os.Stat(filename)
	if err != nil {
		return ""
	}
	if info.IsDir() {
		return "dir"
	}
	return "file"
}

// ReadDir returns the filenames in directory
func ReadDir(dir string) []string {
	filenames := []string{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return filenames
	}

	for _, file := range files {
		if file.IsDir() || !file.IsDir() && (strings.Contains(file.Name(), ".jpg") || strings.Contains(file.Name(), ".jpeg")) {
			filenames = append(filenames, path.Join(dir, file.Name()))
		}
	}

	return filenames
}

// ReadDirRecursive reads the directory recursively by given depth
func ReadDirRecursive(dir string) []string {
	filenames := []string{}
	filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(info.Name(), ".jpg") {
				filenames = append(filenames, path)
			}
			return nil
		})
	return filenames
}

// MkdirIfNot creates the dir if not exists
func MkdirIfNot(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
		return MkdirIfNot(dir)
	}
	return nil
}
