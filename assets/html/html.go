// Code generated by go-bindata.
// sources:
// templates/html/marc-034.html
// DO NOT EDIT!

package html

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesHtmlMarc034Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x51\x4f\xdc\x30\x0c\x7e\xe6\x7e\x85\xc9\x33\xa1\x4c\xa0\x69\xd3\x9a\x4a\xe8\x06\xda\x90\xd0\xa6\x1d\x48\xdb\xa3\xdb\x1a\x1a\x48\x93\x90\xb8\xbd\x3b\x7e\xfd\xd4\xb4\x1c\x85\x31\x6d\x7d\x39\xdb\xf7\x7d\xfe\xe2\xaf\x71\xf3\xfd\xcf\xdf\x96\x57\xbf\xbe\x9f\x41\xc3\xad\x29\x16\xf9\xf0\x03\x06\xed\xad\x12\x64\x45\xb1\x00\xc8\x1b\xc2\xba\x58\xec\x01\xe4\xac\xd9\x50\x71\x79\xfa\x63\x09\x47\xc7\x27\xc0\x0e\x4a\xd7\xd9\x5a\xdb\x5b\x28\xdd\x06\xd8\x39\x93\x67\x23\x28\xe1\x5b\x62\x84\x86\xd9\x4b\x7a\xe8\x74\xaf\xc4\xd2\x59\x26\xcb\xf2\x6a\xeb\x49\x40\x35\x66\x4a\x30\x6d\x38\x1b\x94\x3f\x41\xd5\x60\x88\xc4\xea\xfa\xea\x5c\x7e\x10\xb3\x36\x16\x5b\x52\x22\xd0\x0d\x85\x40\x61\x46\x76\x41\xdf\x6a\x2b\xfe\xa2\xf8\x53\x5e\x9f\xca\xa5\x6b\x3d\xb2\x2e\xcd\x5c\xf4\xeb\x99\xfa\x28\x20\xfb\x43\x02\xbd\x37\x24\x5b\x57\x6a\x43\x72\x4d\xa5\x44\xef\x65\x85\x1e\x5f\xd2\xb7\x14\xff\x9b\x1d\x19\xb9\x8b\xb2\xc4\x20\x23\x6f\x5f\xb4\x29\x0d\x56\xf7\x6f\x35\xfa\x82\xb6\x6e\xc8\xd4\xe7\x41\x93\xad\xcd\x76\x6e\x57\xe8\xe8\x2d\x4a\xaf\x69\xed\x5d\xe0\x19\x74\xad\x6b\x6e\x54\x4d\xbd\xae\x48\xa6\xe4\x00\xb4\xd5\xac\xd1\xc8\x58\xa1\x21\xf5\xee\x00\x5a\xdc\xe8\xb6\x6b\x67\x05\x6d\x5f\x16\xba\x48\x21\x65\x83\x09\xca\xba\xa4\x9e\xe4\xf7\xa5\x84\x35\x41\x20\xb3\x05\x67\x81\x1b\x82\x16\xfd\x23\xd9\xbb\x78\x78\x99\x82\x8b\xd5\x30\x8a\xa1\x30\xdc\x97\xf8\xd0\xe9\xc0\xa0\x47\x64\x20\x43\x3d\x5a\x86\x8b\x55\xb6\x5c\xad\xa0\xa1\x40\x20\x65\xb1\xd8\x7b\x7e\x20\x3d\x49\xcb\x68\x7b\x3f\x70\x94\x48\x2e\xc6\x86\x88\x05\xf0\xd6\xd3\x74\x85\xaa\x18\x05\x34\x81\x6e\x94\x18\xe2\xac\xc5\x50\xc9\xa3\xe3\x93\xc3\xf4\xc7\x64\x58\xac\x82\xf6\x3c\xa7\xdd\x61\x8f\x63\x55\x40\x0c\x95\x12\xb3\xca\x73\x8f\xbb\x28\x8a\x3c\x1b\xab\x69\x2f\xb2\x71\x31\x00\xf2\xd2\xd5\xdb\x62\xb1\x18\x0e\x9a\xd7\xba\x07\x5d\x2b\x31\xf0\x04\x54\x06\x63\x54\x22\xb8\x75\xda\x25\x80\xfc\xc6\x85\x76\x0c\x01\x72\x6d\x7d\x37\x3f\x8a\x98\xde\xe4\x93\xa8\xd8\xb5\x1a\x33\x6f\xb0\xa2\xc6\x99\x9a\x82\x12\x67\x96\x29\x00\xc2\x6e\x1f\x7b\x34\x1d\x25\x13\xc5\x18\x2b\x21\x20\xea\x47\x52\xe2\xfd\x49\x9a\x7f\x92\x2d\x3b\x66\x67\x53\xef\xd8\x95\xad\xde\xb9\x38\x65\x45\xe5\x6c\x4f\x81\x5f\x6f\x78\x9e\x8d\xcc\x69\x94\xec\x79\x96\x3c\xab\x75\xff\xda\x82\x40\xb1\x33\x1c\x9f\x26\x7f\x2a\x97\xa5\xdb\x0c\xcb\x33\xf7\x66\xe2\xbf\x72\xd0\xcf\x41\xe0\xd1\x92\x79\x1b\x1a\x70\xfd\x2f\xe8\xee\x80\x79\x36\xbe\xae\x3c\x1b\x3f\x79\xbf\x03\x00\x00\xff\xff\xee\x7a\x67\x57\x03\x05\x00\x00")

func templatesHtmlMarc034HtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlMarc034Html,
		"templates/html/marc-034.html",
	)
}

func templatesHtmlMarc034Html() (*asset, error) {
	bytes, err := templatesHtmlMarc034HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/marc-034.html", size: 1283, mode: os.FileMode(420), modTime: time.Unix(1515885013, 0)}
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
	"templates/html/marc-034.html": templatesHtmlMarc034Html,
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
		"html": &bintree{nil, map[string]*bintree{
			"marc-034.html": &bintree{templatesHtmlMarc034Html, map[string]*bintree{}},
		}},
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

