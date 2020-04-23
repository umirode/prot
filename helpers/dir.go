package helpers

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
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

func FindFilesInDirByExt(dir string, ext string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func JoinPathAndFileName(file string, paths ...string) string {
	var p string
	for _, i := range paths {
		p = path.Join(p, path.Clean(i))
	}

	return path.Join(p, file)
}
