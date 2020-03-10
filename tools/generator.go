package tools

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

type generatorLangArgs struct {
	ProtocArg       string
	ProtocOutputExt string
}

func generatorGetSupportLanguagesMap() map[string]generatorLangArgs {
	return map[string]generatorLangArgs{
		"go": {
			ProtocArg:       "--go_out=plugins=grpc:.",
			ProtocOutputExt: ".pb.go",
		},
	}
}

func (*Generator) findFilesInDirByExt(dir string, ext string) []string {
	var files []string

	filepath.Walk(dir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, path)
			}
		}
		return nil
	})

	return files
}

func (g *Generator) GenerateProto(modulePath string, lang string) ([]string, error) {
	supportLanguagesMap := generatorGetSupportLanguagesMap()

	supportLang, ok := supportLanguagesMap[lang]
	if !ok {
		return nil, errors.New("lang " + lang + " not exists")
	}

	cmd := exec.Command("protoc", modulePath+"*.proto", supportLang.ProtocArg)

	//TODO: Debug
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	//TODO: Debug

	err := cmd.Run()
	if err != nil {

		//TODO: Debug
		fmt.Println("Result: " + out.String())
		fmt.Println("Result: " + stderr.String())
		//TODO: Debug

		return nil, err
	}

	return g.findFilesInDirByExt(modulePath, supportLang.ProtocOutputExt), nil
}
