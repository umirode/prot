// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package config generated by go-bindata.// sources:
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _templatesProtYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcc\xcb\xaa\xc2\x30\x10\x06\xe0\x7d\x9f\x62\x9e\xa0\xd9\xcf\xea\x1c\x04\x51\xaa\x52\x44\x70\x29\xbd\x8c\xc9\x40\xd3\x09\xc9\x64\x91\xb7\x17\xf1\x52\x44\x77\x73\xf9\xff\x6f\xd7\xcd\x16\xc1\x4a\xb5\x97\x31\x4f\x94\xb0\x02\x78\x8c\x87\xce\xd3\x7d\x03\x38\x52\x90\xc4\x2a\xb1\x20\x38\xd5\x90\xd0\x18\xcb\xea\x72\x5f\x0f\xe2\x4d\xf6\x1c\x65\x24\x13\xa2\x68\x6d\x59\x3f\x84\x33\xab\xfb\xcf\xea\xbe\x25\xcb\xfa\xb7\x28\xf8\x43\x01\x58\x9a\x00\xa7\x12\x08\xa1\xcd\xfd\xc4\x43\x43\x25\x3d\xcf\x2b\x99\xaf\x6c\x5f\x21\x80\x96\xfc\x9a\x27\x42\xe0\xf1\x12\x53\x57\x07\xf2\xef\xdf\xd6\xce\x12\x69\x23\x49\x1b\x2a\x08\x1a\x33\xdd\x02\x00\x00\xff\xff\x60\x44\x55\x55\x00\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/prot.yaml", size: 256, mode: os.FileMode(420), modTime: time.Unix(1585202484, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x5b\x8f\xd3\x3c\x10\x7d\xcf\xaf\xb0\xbc\x7d\xf8\xbe\xd5\x96\xec\x1b\xa2\x6f\x0b\x12\x02\x2d\x97\x0a\xc1\x0b\x55\x40\x6e\x3b\x6d\xbd\x4a\xe3\x62\x3b\x42\x65\xc9\x7f\x47\x76\x92\xd6\xf1\x25\x21\xbd\xa0\x45\xec\x3e\xe4\x32\x67\x66\x7c\xce\xd8\x33\x9b\xde\x47\x08\xe1\x81\x98\xad\x60\x4d\xf0\x08\xe1\x95\x94\x9b\x51\x1c\xdf\x09\x96\x0d\x4b\xeb\x13\xc6\x97\xf1\x9c\x93\x85\x1c\x5e\x3f\x8d\x4b\xdb\x05\xbe\x52\x71\x72\xbb\x01\x15\xc4\xa6\x77\x30\x93\xa5\x8d\xc3\xb7\x9c\x72\x98\xe3\x11\x9a\x44\x08\x21\x84\xdf\x90\x6c\xa9\x31\x84\xf0\x5b\x36\xcf\x53\x10\x38\x42\x28\xd1\xee\x1b\xce\x36\xc0\x25\x05\x81\x47\xe8\xde\x08\xa8\xdf\x14\x3d\xaa\xb2\xe1\x8b\x78\xef\x1c\x1b\x49\x0d\x22\x42\x72\x6a\xda\x21\xcb\xd7\x3b\x22\xda\xb2\x64\x3b\x14\x21\x3c\x27\x5c\xe2\xea\x35\xd1\xf7\xc2\x22\xda\x41\xa3\x76\x73\x98\x18\x25\x29\x57\x82\x05\xc9\x53\xa9\xf2\x15\x3b\x63\x95\x68\xfb\x8e\xac\x1b\x4b\x29\x88\x48\x09\x3c\x53\xa9\xfe\xfb\x32\x21\xc3\x1f\xc9\xcf\xc9\xcd\xf0\xf3\xf5\xf0\x59\xf2\xbf\x7e\xbd\x1c\xd4\xc4\x8d\x7c\x65\xd0\xd8\xad\xa9\x86\x83\x89\x4c\xaf\x76\xa1\xb1\x4e\xf1\x35\x29\xaf\x97\x03\xa3\x98\x61\xf1\x1a\x73\xce\xc5\x0e\xf9\x00\x1b\x26\xa8\x64\x7c\x8b\x0d\x20\x69\x44\x7b\x8e\x89\x2f\xde\xc6\xfa\x89\x89\x8d\x4c\x57\x76\x1e\xeb\x88\x35\xe0\xa2\xe9\x8d\x6f\x72\xb9\x3a\x96\x8b\xce\x11\x64\xe1\xa9\x70\x7b\x95\x35\xfa\x51\x45\xdb\x31\x08\xe1\x17\x2c\x5b\x50\x4b\x93\xb5\x03\xa8\x63\x17\xf6\xf9\x7d\xc8\x01\xda\xe3\x8a\x94\x4b\x37\xdc\xef\x0d\x1f\xa7\xf7\x1b\xe8\x38\x9f\xa6\x74\x76\x0b\x5b\xe1\x8d\x56\x1e\x44\x88\xef\x8c\xcf\x43\xf8\x73\x22\xe8\x4c\xef\x92\x07\x4f\x1c\x5b\x61\x59\x0a\xa7\xba\x24\x4d\xdf\x2f\xbc\x8c\xbd\x05\xa5\x8b\x40\xa9\xbb\x37\xaa\xf2\x6a\xd9\xae\xca\x63\xc6\x32\xa1\x66\x96\x59\xae\x80\xb7\xad\x2f\x6c\x75\x94\xeb\xa5\xe4\x0a\xb2\x63\x05\x55\x47\xa6\x55\x52\x6b\x87\x18\x7e\x63\x58\xbf\xa4\x29\x84\xe4\x7a\xda\xa3\x2f\xdd\xc6\x32\xed\x6e\x87\xf7\x4f\x5c\x2f\x10\x66\x8b\xba\xc6\x5b\xf3\xcf\xbb\x7d\xb6\x9e\x5d\xf3\x9c\x5b\x57\x47\x97\x9e\x5c\xdf\x27\x01\xfc\x7c\xa2\x74\xf6\x7e\x4a\x3a\xbc\xf7\xdf\x1d\x78\x49\xe5\xe1\xba\x5f\x2f\x33\xc6\xe1\x15\x13\xf2\x16\x7c\xff\x67\x4f\x54\x80\xe6\x32\xbf\x59\x89\x29\x63\x29\x90\xac\x55\x5c\x10\x0b\x21\x3d\xe6\x99\x3b\xeb\x5d\xe2\x0f\x60\x84\xd7\x9d\xf2\x8f\x0c\xf0\xce\x5e\xea\xac\xc8\xa9\x66\xfc\x83\x9b\x19\x87\x4f\xf7\xf3\x8f\xf5\xd3\xcf\xf3\xc7\xd6\x6f\xfd\x54\xf5\xeb\x08\x59\xff\x9e\xde\xcf\xc8\xba\xfd\xab\xe7\x8f\xf6\xbf\x66\x73\xd6\x19\xd0\xa9\xf7\x71\x0e\xf4\x44\x8e\x9a\x03\xf6\x8f\xe9\x28\x84\xee\x9f\xeb\xa7\xf2\xae\xae\x45\x54\x44\xbf\x02\x00\x00\xff\xff\x84\x0a\xb8\x9a\xa2\x13\x00\x00")

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

	info := bindataFileInfo{name: "templates/schema.json", size: 5026, mode: os.FileMode(420), modTime: time.Unix(1585338293, 0)}
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
