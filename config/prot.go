// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/prot.yaml
// templates/schema.json
package config

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesProtYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xf2\x49\xcc\x4b\xb7\x52\x48\xcf\xe7\xe5\xf2\xcd\x4f\x29\xcd\x49\x2d\xb6\xe2\xe5\x52\x56\x50\x08\x49\x2d\x2e\x81\xb0\x14\x14\x42\x8b\x72\x60\x4c\xc7\xd2\x92\x8c\x90\xca\x82\x54\x2b\x85\x80\xd2\xa4\x9c\xcc\x64\xef\xd4\xca\x62\x24\x29\xe7\xfc\xbc\xb4\xcc\x74\x98\x62\x05\x85\x80\xd4\x5c\xb7\xcc\x9c\x54\x84\x80\x67\x7a\x5e\x7e\x51\xaa\x47\x7e\x71\x89\x77\x6a\xa5\x95\x42\x49\x51\x69\x2a\x20\x00\x00\xff\xff\x6b\xd0\x3f\x42\x81\x00\x00\x00")

func templatesProtYamlBytes() ([]byte, error) {
	return bindataRead(
		_templatesProtYaml,
		"templates/prot.yaml",
	)
}

func templatesProtYaml() (*asset, error) {
	bytes, err := templatesProtYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/prot.yaml", size: 129, mode: os.FileMode(438), modTime: time.Unix(1583402758, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x96\xcd\x6a\xdc\x30\x10\xc7\xef\x81\xbc\xc3\xa0\xf6\x98\xd8\x9b\x42\x28\xec\x2d\xb4\x87\xb6\x87\xb4\x34\x81\x1e\x96\x3d\x68\xed\x59\x5b\xa9\x2d\xb9\xd2\x98\x26\x2c\x7e\xf7\x22\x6b\x6d\xcb\x5f\x31\x0d\x61\xc9\x5e\x22\x8d\xf4\x1f\x49\xbf\x19\xcf\xe4\x70\x7e\x06\xc0\xde\x9b\x28\xc5\x9c\xb3\x35\xb0\x94\xa8\x58\x87\xe1\x83\x51\xf2\xd2\x59\x03\xa5\x93\x30\xd6\x7c\x4f\x97\xab\x8f\xe1\x71\xe7\x85\xd3\x89\xd8\xd3\xe0\x23\xcf\x8b\x0c\x83\x48\xe5\xa1\x56\x8a\x02\xeb\xe4\xb8\x93\x9e\x0a\xb4\x5b\xd5\xee\x01\x23\x6a\x8c\x82\xb2\xda\x7a\x9f\x22\xfc\x54\x8a\xe0\xce\xf7\x1e\xa3\x89\xb4\x28\x48\x28\xd9\x6c\xb2\x6e\xc1\x5d\x01\x84\x01\x4a\xb1\x99\x51\xca\x09\x22\x95\x17\x5a\x18\x74\x2b\x28\x49\x68\x84\x6f\x77\xdf\x6f\x21\x56\x51\x99\xa3\xa4\xa0\xf5\xbd\xe7\x65\x46\x6c\x0d\x87\xca\x59\x34\xfe\x29\x85\x46\xfb\xa0\x8d\x35\x00\xb0\x28\xc5\xe8\x37\xc6\x4e\x62\x45\x22\x47\x69\x84\x92\xa6\x35\x89\x6e\x55\xf2\x1c\xdb\x49\xa1\x45\xd4\xcd\x88\x27\x86\xd9\xf1\xd6\x9d\x55\x68\x55\xa0\x26\x81\xc6\x5e\x60\x70\x5a\x6b\xe9\x00\xbf\x0b\x3b\x45\x38\xb8\x95\x07\x77\xa7\x54\x86\x5c\xfa\x2b\x3e\xe1\x4f\x4e\xd8\x83\xec\xde\xd5\x07\x7d\x23\x01\x1f\x8b\x8c\x4b\x6e\x2d\xc0\x77\xaa\xa4\x1a\x68\x51\xea\x42\x19\x04\xb5\x07\x4a\x85\x01\x21\x0d\x71\x19\x61\xd0\xf3\xd5\x80\xdd\xf3\xcc\x60\xb7\x70\xcc\x0e\xd3\xf1\xb5\xbf\x7a\x53\x33\xdd\xba\x41\x35\xc1\x7b\x09\xca\x38\x34\x33\x49\x37\xc6\xf2\xb9\x95\x9e\x8a\xcc\xa1\x5a\xc2\x72\xe8\x86\x00\x2c\x45\x91\xa4\x56\x78\xb5\x0a\x56\x17\xbd\xa5\xbf\x22\xa6\x94\xad\xe1\x3a\x58\x75\xf6\xaa\x05\xda\x9d\x33\x4e\x6f\x4f\xef\xf9\x6c\x0e\x9b\x70\x31\x95\xb5\xbd\x4b\xf4\x6f\xfd\x7c\x98\x7c\xeb\xe8\x12\x5e\xec\x84\x24\x4c\x50\x0f\x57\xfd\x00\xfe\xb2\xf2\x51\xec\x5e\x3b\x7e\x83\x18\x0e\xe2\x30\x1d\x46\xfb\xbb\xf6\xa7\x5b\x2f\x48\x13\xd0\x5f\x8e\xf0\xe8\xe0\xc5\x0c\xbf\xd4\xfa\x37\x0b\xf1\x6a\x35\x47\xb1\x19\x56\x83\xda\x21\x96\x0b\xa9\x98\xaa\xa1\x23\x56\x7d\x4e\x5f\x4f\x56\x3e\x57\x4b\x35\xe2\x6a\xae\x6c\xd6\x8d\x68\xe9\xf1\x7e\xb7\xf2\x9e\x6f\x48\x0b\x99\xcc\xbd\xfe\x96\xe7\x78\xaa\xf7\x33\xb6\x04\x80\xdd\x40\xa2\x11\x25\xc4\x4a\x69\x36\x47\xc3\x75\xe2\x25\x1c\xbd\x7e\xed\xf1\x90\x65\xbe\x9b\xcf\x86\x1f\x56\xf5\x76\x12\xe2\x43\x70\x3d\x47\xa1\xfe\x0f\x64\x09\x42\xbd\x69\xcc\x80\x6b\xcd\x9f\xe6\x10\xdc\xf3\xe4\x64\x7d\x73\xb3\x5d\x42\xb0\xe9\xf7\x4d\xe5\xe7\x78\x6d\xaa\x13\x86\x75\xa6\xed\x44\x9f\x13\x84\xf9\xb0\xc5\xcd\xe1\x0a\xdd\x66\xbf\x96\xcf\x7d\x4a\xa3\x52\x62\x95\x13\x15\xf7\x75\xeb\xed\xf4\x17\xf5\x4c\xb1\xfd\x2f\x6a\x4d\xd5\xb5\x7f\xaa\xf3\xb3\xea\x5f\x00\x00\x00\xff\xff\xa8\x4e\xa6\xa4\x4f\x0c\x00\x00")

func templatesSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_templatesSchemaJson,
		"templates/schema.json",
	)
}

func templatesSchemaJson() (*asset, error) {
	bytes, err := templatesSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/schema.json", size: 3151, mode: os.FileMode(438), modTime: time.Unix(1583482211, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/prot.yaml":   templatesProtYaml,
	"templates/schema.json": templatesSchemaJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"prot.yaml":   &bintree{templatesProtYaml, map[string]*bintree{}},
		"schema.json": &bintree{templatesSchemaJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
