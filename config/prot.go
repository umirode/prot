// Code generated for package config by go-bindata DO NOT EDIT. (@generated)
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

var _templatesProtYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcc\xcd\x8a\xc2\x30\x10\x07\xf0\x7b\xa1\xef\x30\x4f\xd0\xdc\xe7\xb4\xcb\xc2\xb2\x4b\x55\x8a\x08\x1e\xa5\x1f\x63\x32\xd0\x74\x42\x32\x39\xe4\xed\x05\x45\xab\xd0\xdb\x7c\xfc\xff\xbf\x5d\xbf\x58\x04\x2b\x75\xb5\x97\x29\xcf\x94\xb0\xae\x00\x1e\xf3\xa1\xf7\x74\x5f\x01\x8e\x14\x24\xb1\x4a\x2c\x08\x4e\x35\x24\x34\xc6\xb2\xba\x3c\x34\xa3\x78\x93\x3d\x47\x99\xc8\x84\x28\xda\x58\xd6\x4f\xe3\xcc\xea\xbe\xb3\xba\x0d\xcb\xb2\x7e\xad\x0e\x6e\x39\x00\x6f\x5d\x80\x53\x09\x84\xd0\xe5\x61\xe6\xb1\xa5\x92\x9e\xf7\x1f\x59\xae\x6c\x5f\x31\x80\x8e\xfc\x2f\xcf\x84\xc0\xd3\x25\xa6\xbe\x09\xe4\xd7\xe7\xbf\x5d\x24\xd2\x9f\x24\x6d\xa9\x20\x68\xcc\x74\x0b\x00\x00\xff\xff\xdc\xdd\xe4\x01\x0a\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/prot.yaml", size: 266, mode: os.FileMode(438), modTime: time.Unix(1583773904, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4d\x8f\xd3\x30\x10\xbd\xaf\xb4\xff\xc1\x32\x3d\xc0\x6a\x83\xf7\x86\xe8\x6d\x85\x84\x40\xcb\x47\x85\xe0\x42\x15\x90\xdb\x4c\x52\xaf\xd2\x38\xd8\x8e\x50\x59\xf2\xdf\x91\x9d\x34\xb1\x13\x27\x69\xb7\x74\x73\x68\x63\xbf\x99\x37\xf3\x5e\xa6\x4e\x1f\x2e\x2f\x10\xc2\x33\xb9\xde\xc0\x96\xe2\x39\xc2\x1b\xa5\xf2\x39\x21\xf7\x92\x67\x41\xb5\xfb\x92\x8b\x84\x44\x82\xc6\x2a\xb8\x79\x45\xea\xc8\x6b\x93\xa7\x76\x39\xe8\x24\xbe\xba\x87\xb5\xaa\x37\x05\xfc\x2a\x98\x80\x08\xcf\xd1\x52\x6f\x20\x84\x3f\xd0\x2c\xa9\x50\x84\xf0\x47\x1e\x15\x29\x48\xac\x97\x61\x95\x92\x0b\x9e\x83\x50\x0c\x24\x9e\xa3\x07\x3b\xa9\x59\xea\x2e\x99\xe6\xc4\xcf\x48\x1b\x4e\x6c\x66\xab\x21\xa9\x04\x73\x00\xc8\x8a\x6d\xdb\x90\xd9\x4a\x38\xde\xaf\xc2\xea\xa6\xec\xb6\x38\x55\x7d\x1f\xd7\x6f\xc0\x76\xc4\x00\x11\xc4\xb4\x48\x95\xa6\x2c\xdb\xdd\x9a\x6b\xf7\x89\x6e\xdd\x72\x1a\xa3\x4a\x81\xc8\x34\xdb\xf3\x1f\x4b\x1a\xfc\x09\xff\x2e\x6f\x83\xef\x37\xc1\xeb\xf0\x85\x59\x5e\xcd\x1a\x01\x36\x65\x95\xb6\xf0\x58\x6a\xf0\x41\x2e\x27\x6c\x5c\x30\x31\x1c\x3f\xc3\xea\xf3\x6a\xd6\xea\x1c\x35\xc1\x80\xfd\xf1\x68\xa0\x2f\x90\x73\xc9\x14\x17\x3b\x6c\x23\xa1\x4b\xe0\x1b\x16\x1f\x45\x0f\x3c\x4e\x13\xb1\xa8\xae\x7b\x44\x9d\x41\x73\xf1\xb2\x13\x8f\x6f\x0b\xb5\x39\xb9\x1f\x43\x32\xdc\x89\xcf\xec\x09\xc3\x0d\xfc\x55\xe7\xf7\xb2\x10\xc2\x6f\x78\x16\xb3\xae\xb4\xee\xd3\x40\x53\x4f\xa4\xad\xe1\x85\x1e\xe1\x01\xa9\x1b\xf3\xf4\x3c\x72\x02\x38\x41\xfd\xd3\xc0\x81\x17\xc5\x2a\x65\xeb\x3b\xd8\xc9\x9e\x7a\xd4\x9e\x16\xf6\x55\x76\xb7\xba\x23\x80\x10\xa6\x69\xfa\x39\xf6\x57\xf5\x1b\xc3\xe2\x21\xcf\x0e\xf0\xbc\x0e\x1b\x73\xbe\x0e\x59\xf3\x4c\xea\x73\x69\x4a\xb6\x57\xe7\xc8\x76\xdf\x02\x53\x4e\x6d\x20\x3b\x5d\x57\x3d\x03\xe3\xca\xc6\x27\xdf\x0a\x5c\xc0\xf6\x2d\x4b\x61\x50\xb6\x6f\xea\x8f\xee\xd9\xa9\x34\x11\xf7\xf8\x9f\x05\xd9\x57\x18\x69\x18\x4d\x9e\x60\xee\xe5\x7f\x92\x5d\x4d\x0b\x2a\xe5\x6f\x2e\xa2\xf3\x6b\x6b\x2a\x3d\x9d\xc6\x6f\x12\xc4\x19\x85\x19\xfa\x23\xd5\x4c\x85\xb7\xff\x38\x70\xc2\xd4\x29\xe2\xdf\x27\x19\x17\xf0\x8e\x4b\x75\x07\xde\xf7\xea\x7f\x72\xc1\xad\x73\xa8\x1d\x2b\xce\x53\xa0\xd9\xb8\xc0\x61\x70\x10\x3a\xea\xac\x3b\xe0\x95\xd0\x79\x6d\x38\xb8\xb5\x68\x6e\xeb\x1b\xf3\x55\x5e\x5e\x94\xff\x02\x00\x00\xff\xff\x23\x2d\xd1\xb6\xab\x0b\x00\x00")

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

	info := bindataFileInfo{name: "templates/schema.json", size: 2987, mode: os.FileMode(438), modTime: time.Unix(1583773400, 0)}
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
