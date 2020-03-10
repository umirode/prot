package tools

import (
	"os"
	"path"
)

func CreateDirRecursive(dir string) error {
	err := os.MkdirAll(path.Clean(dir), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func CleanDir(dir string) error {
	err := os.RemoveAll(path.Clean(dir))
	if err != nil {
		return err
	}

	return nil
}

func JoinPathAndFileName(file string, paths ...string) string {
	var p string
	for _, i := range paths {
		p = path.Join(p, path.Clean(i))
	}

	return path.Join(p, file)
}
