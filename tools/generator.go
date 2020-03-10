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
	ProtocOutArg    string
	ProtocOutputExt string
}

func generatorGetSupportLanguagesMap() map[string]generatorLangArgs {
	return map[string]generatorLangArgs{
		"go": {
			ProtocOutArg:    "--go_out=plugins=grpc:",
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

func (g *Generator) GenerateProto(moduleDir string, lang string) ([]string, error) {
	supportLanguagesMap := generatorGetSupportLanguagesMap()

	supportLang, ok := supportLanguagesMap[lang]
	if !ok {
		return nil, errors.New("lang " + lang + " not exists")
	}

	cmd := exec.Command("protoc", supportLang.ProtocOutArg+moduleDir, "-I"+moduleDir, JoinPathAndFileName("*.proto", moduleDir))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	fmt.Println(cmd.String())

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("%v %v", err, stderr.String())
	}

	return g.findFilesInDirByExt(moduleDir, supportLang.ProtocOutputExt), nil
}
