package tools

import (
	"log"
	"os"
)

func CreateDirAndFormatPath(outputPath string) string {
	if outputPath == "" {
		return ""
	}

	if outputPath[len(outputPath)-1:] != "/" {
		outputPath += "/"
	}

	_ = os.RemoveAll(outputPath)

	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	return outputPath
}
