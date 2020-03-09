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

var _templatesSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4d\x8f\xd3\x30\x10\xbd\xaf\xb4\xff\xc1\xf2\xf6\x00\xab\x0d\xde\x1b\xa2\xb7\x15\x12\x02\x2d\x1f\x15\x82\x0b\x55\x40\x6e\x33\x49\xbd\x4a\xe3\x60\x3b\x42\x65\xc9\x7f\x47\x76\xd2\xc4\x4e\x9c\xa4\xdd\x52\x72\x68\xe3\x79\xe3\x37\xf3\x5e\xa6\x4e\x1f\x2f\x2f\x10\xc2\x33\xb9\xde\xc0\x96\xe2\x39\xc2\x1b\xa5\xf2\x39\x21\x0f\x92\x67\x41\x15\x7d\xc1\x45\x42\x22\x41\x63\x15\xdc\xbe\x24\x55\xec\x0a\xdf\x98\x8d\x6a\x97\x83\xde\xc5\x57\x0f\xb0\x56\x75\x50\xc0\xcf\x82\x09\x88\xf0\x1c\x2d\x75\x00\x21\xfc\x9e\x66\x49\x85\x22\x84\x3f\xf0\xa8\x48\x41\x62\xbd\x0c\xab\x2d\xb9\xe0\x39\x08\xc5\x40\xe2\x39\x7a\xb4\x37\x35\x4b\xdd\x26\xd3\x9c\xf8\x8a\xb4\xe9\xc4\x66\xb6\x1a\x92\x4a\x30\x07\x80\xac\xd8\xb6\x0d\x99\x50\xc2\xf1\x7e\x15\x56\x37\x65\xb7\xc5\xa9\xea\xfb\xbc\x7e\x03\xb6\x23\x06\x88\x20\xa6\x45\xaa\x34\x65\xd9\x46\x6b\xae\xdd\x47\xba\x75\xcb\x69\x8c\x2a\x05\x22\xd3\x6c\xcf\xbe\x2f\x69\xf0\x3b\xfc\xb3\xbc\x0b\xbe\xdd\x06\xaf\xc2\xe7\x66\x79\x3d\x6b\x04\xd8\x94\xd5\xb6\x85\xc7\x52\x83\x0f\x72\x39\x69\xe3\x82\x89\xe1\xf8\x11\x56\x9f\xd7\xb3\x56\xe7\xa8\x09\x06\xec\x8f\x47\x03\x7d\x86\x9c\x4b\xa6\xb8\xd8\x61\x1b\x09\x5d\x02\xdf\xb0\xf8\x28\x7a\xe0\x71\x9a\x88\x45\x75\xd3\x23\xea\x0c\x9a\x8b\x97\x9d\x7c\x7c\x57\xa8\xcd\xc9\xfd\x18\x92\xe1\x4e\x7c\x66\x4f\x18\x6e\xe0\x2f\x7a\x7f\x6f\x17\x42\xf8\x35\xcf\x62\xd6\x95\xd6\x7d\x1a\x68\xea\x89\xb4\x35\xbc\xd0\x13\x3c\x20\x75\x63\x9e\x9e\x47\x4e\x00\x27\xa9\x7f\x1a\x38\xf0\xa2\x58\xa5\x6c\x7d\x0f\x3b\xd9\x53\x8f\xda\xd3\xc2\xbe\xca\x6e\xa8\x3b\x02\x08\x61\x9a\xa6\x9f\x62\x7f\x55\xbf\x31\x2c\x1e\xf2\xec\x00\xcf\xeb\xb4\x31\xe7\xeb\x94\x35\xcf\xa4\x3e\x97\xa6\x64\x7b\x75\x8e\x84\xfb\x16\x98\x72\x6a\x03\xd9\xe9\xba\xea\x19\x18\x57\x36\x3e\xf9\x56\xe2\x02\xb6\x6f\x58\x0a\x83\xb2\x7d\x53\x7f\x74\xcf\x4e\xa5\x89\xbc\xa7\xff\x2c\xc8\xbe\xc2\x48\xc3\x68\xf2\x04\x73\x2f\xff\x93\xec\x6a\x5a\x50\x29\x7f\x71\x11\x9d\x5f\x5b\x53\xe9\xff\x69\xfc\x2a\x41\x9c\x51\x98\xa1\x3f\x52\xcd\x54\x7a\xfb\x8f\x03\x27\x4c\x9d\x22\xfe\x5d\x92\x71\x01\x6f\xb9\x54\xf7\xe0\x7d\xaf\xfe\x23\x17\xdc\x3a\x87\xda\xb1\xe2\x3c\x05\x9a\x8d\x0b\x1c\x06\x07\xa1\xa3\xce\xba\x03\x5e\x09\x9d\xd7\x86\x83\x5b\x8b\xe6\xb6\xbe\x31\x5f\xe5\xe5\x45\xf9\x37\x00\x00\xff\xff\x81\xb5\x63\xb2\xac\x0b\x00\x00")

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

	info := bindataFileInfo{name: "templates/schema.json", size: 2988, mode: os.FileMode(438), modTime: time.Unix(1583782568, 0)}
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
