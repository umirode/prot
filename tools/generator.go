package tools

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
)

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}
}

type generatorLangArgs struct {
	ProtocOutArg    string
	ProtocOutputExt []string
}

func generatorGetSupportLanguagesMap() map[string]generatorLangArgs {
	return map[string]generatorLangArgs{
		"go": {
			ProtocOutArg: "--go_out=plugins=grpc:",
			ProtocOutputExt: []string{
				".pb.go",
			},
		},
		"dart": {
			ProtocOutArg: "--dart_out=grpc:",
			ProtocOutputExt: []string{
				".pb.dart",
				".pbenum.dart",
				".pbgrpc.dart",
				".pbjson.dart",
			},
		},
	}
}

func (g *Generator) GenerateProto(moduleDir string, lang string) ([]string, error) {
	supportLanguagesMap := generatorGetSupportLanguagesMap()

	supportLang, ok := supportLanguagesMap[lang]
	if !ok {
		return nil, errors.New("lang " + lang + " not exists")
	}

	protoFiles, err := FindFilesInDirByExt(moduleDir, ".proto")
	if err != nil {
		return nil, err
	}

	if len(protoFiles) == 0 {
		return nil, errors.New(".proto files not found in " + moduleDir + " directory")
	}

	cmd := exec.Command("protoc", supportLang.ProtocOutArg+moduleDir, "-I"+moduleDir)
	cmd.Args = append(cmd.Args, protoFiles...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	logrus.Debug("protoc generation command: " + cmd.String())

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("%v %v", err, stderr.String())
	}

	var generatedFiles []string

	for _, ext := range supportLang.ProtocOutputExt {
		generatedFilesByExt, err := FindFilesInDirByExt(moduleDir, ext)
		if err != nil {
			return nil, fmt.Errorf("finding generated proto files error: %v", err)
		}

		generatedFiles = append(generatedFiles, generatedFilesByExt...)
	}

	return generatedFiles, nil
}
