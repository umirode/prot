package tools

import (
	"os"
)

func CreateDirAndFormatPath(outputPath string, cleanDir bool) (string, error) {
	if outputPath == "" {
		return "", nil
	}

	if outputPath[len(outputPath)-1:] != "/" {
		outputPath += "/"
	}

	if cleanDir {
		err := os.RemoveAll(outputPath)
		if err != nil {
			return "", err
		}
	}

	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
