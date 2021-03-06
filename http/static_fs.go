// Code generated by go-bindata.
// sources:
// static/javascript/marc-034.js
// static/css/marc-034.css
// DO NOT EDIT!

package http

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _staticJavascriptMarc034Js = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\xdd\x6e\xdb\x38\x13\xbd\x96\x9e\x62\xc0\x9b\x4a\x2d\x2d\xbb\x4d\xf3\xe1\x43\x0c\x61\xb1\x2d\x52\xb4\xbb\x49\x0a\x24\x59\x74\x81\x34\x08\x68\x69\xec\xb0\xa5\x48\x96\xa4\xea\xb8\x81\xdf\x7d\x41\xea\xc7\x72\xe3\x18\xc9\x8d\x6c\x69\xce\x9c\x19\xce\x1c\x0e\x39\x7e\x09\x15\x33\xc5\xcd\xe4\xe0\x6d\xf6\xcd\xc2\xcb\x71\x1c\x2f\xb9\x2c\xd5\x32\x63\x65\x79\xfc\x13\xa5\x3b\xe1\xd6\xa1\x44\x93\x10\xa1\x58\x49\x28\xcc\x6b\x59\x38\xae\x24\xf8\xf7\x04\x3d\x26\xbd\x8f\x63\x00\x80\x9f\xcc\x40\xc5\xf4\x34\xbc\x84\x07\x9f\x43\x72\x92\x9d\x32\xfd\x0b\x65\x0a\xf7\x71\x1c\x79\x0c\xd3\xfc\xe6\x3b\xae\x20\x87\x52\x15\x75\x85\xd2\x65\x33\x55\xae\xb2\x05\xba\x3f\x9d\x33\x7c\x56\x3b\x4c\x48\xc9\x1c\x1b\x55\xc1\x75\xc4\x34\x1f\x7d\xc7\x15\x49\xa7\x71\xd4\xf1\x65\x4c\xf3\xbf\x03\x4b\xcb\x37\x8d\xa3\x86\xbf\x62\xfa\x46\x69\x67\x21\x87\x7b\x70\x4c\x2e\x0c\xab\x3e\x6b\x9f\xb4\x3d\x82\xfb\x38\x8a\x22\x5b\xa0\xc4\x23\xe8\xa9\xde\x31\x8b\x15\xd3\x17\x6e\x25\xd0\x66\xe7\x38\xe7\x42\x50\x0f\x1c\x8f\xa1\xc5\x92\x71\xcb\x34\x36\xc1\x3c\xb2\x1e\x9c\xfd\xe2\x9a\xc4\xd1\x7a\x1d\x82\x57\x4c\x43\xbe\x61\xad\x98\x4e\x5e\x54\x4c\xbf\xa0\x7d\x4a\x69\x9f\xa4\x5d\x42\x0e\x57\x30\x3a\x3c\xa4\x30\x7a\xfd\xff\x09\x5c\x4f\x1b\x83\xc4\x60\xf0\xdf\x87\x9f\x67\xaa\x96\xa5\x0d\x26\xbb\xa4\x1e\x75\xdd\xc5\xcc\xe6\xdc\xbd\x0b\xe6\xa4\x41\xa5\x4d\x0b\xd6\x4d\x5b\x50\x58\xec\x6b\x8f\x62\x58\xf6\x05\xba\x63\x81\xfe\xef\xbb\xd5\xa7\x32\x21\x15\xd3\xa1\xc6\x28\x32\x2e\x25\x9a\x4b\xbc\x73\x90\x03\x39\x65\xda\x02\x33\x08\x25\xb7\x6c\x26\xb0\x84\x19\x16\xac\xb6\x08\x4d\x83\xbc\x78\xb8\x05\xa9\x1c\x68\x83\xd6\x53\x93\x2e\x87\x5e\x0c\xa1\x35\xfb\xa3\x9b\x62\x34\x39\x78\x4b\xda\xfc\xab\xec\x27\x13\xb5\x2f\x07\x21\xd3\x6d\x1e\xbb\x8f\xc7\xd6\xb3\x8a\xbb\x8e\x25\x3c\x6c\xa6\x64\x21\x78\xf1\x1d\xf2\x5e\xc1\x49\xda\x15\xe5\xc9\x59\x45\x1e\xd9\xa6\xd5\xbf\x39\xc3\xab\x24\x9d\xc6\x71\xe4\xe5\x5e\x41\xee\xf3\x4d\xef\xe3\xc8\x07\x36\xe8\x6a\x23\x61\xce\x84\xf5\x1e\xeb\xae\xfd\x28\x0b\xc8\xfd\x53\x95\xf8\xcf\xf9\xa7\xf7\xaa\xd2\x4a\xa2\x74\x49\x95\xb6\x0d\xaf\x8d\x6f\x95\x50\x05\xf3\xc9\x66\xda\x28\xa7\x0a\x25\xe0\x15\x90\xf1\x98\xc0\xab\x8d\xe9\x56\x59\x17\x3e\xcf\x66\xea\xee\x8f\xc9\xc1\xdb\xdc\x9b\x51\x16\xbd\xd8\x0c\x5b\xee\x5b\xa2\x61\xcb\xb0\x3a\xc3\x96\x4d\xdf\x3f\x5e\x9e\x9e\xb4\x65\x6f\x19\x3c\x35\xee\xad\x7a\x83\x08\x3c\xcd\xdf\x07\x54\x6d\x81\x98\xee\x8a\xf3\xe8\x3e\xe8\x8c\x3b\xf6\x42\x67\xda\xbd\x1f\x42\xc3\xc3\xf3\x91\x7d\xb1\xe9\x80\x92\x37\xb6\x2e\x0a\xb4\x76\xa8\x09\x63\x75\x90\x45\x9f\x9e\x33\x90\xc3\x5f\x17\x9f\xcf\x32\xeb\x0c\x97\x0b\x3e\x5f\x79\x10\x05\x59\x0b\x41\xe1\x4d\x3a\xc8\x49\x1b\x1c\x56\xa8\x30\xc8\x1c\xb6\x45\x4a\x88\x36\x48\x3a\xb4\x36\x98\x31\xad\x51\x96\xef\x6f\xb9\x28\x93\xdf\x7c\xfc\xbe\x3b\x53\x25\x26\xb5\x11\xe9\x73\x7d\xc8\x57\xf9\x55\x92\x34\x9d\x46\x51\x14\x3d\xcf\xd5\x3a\xd3\x87\xdb\x2c\xeb\x89\xea\x09\x72\x67\xcb\xad\x40\xda\xe0\x43\x3e\x2f\x0f\xc8\xc1\x58\x7d\x15\x54\x43\x86\x7d\xad\xb8\xf4\x46\xff\xfd\x6a\x72\x9d\x39\xf5\x81\xdf\x61\x99\xfc\x2f\xdd\xc6\xac\x3a\xcc\xeb\xc7\x30\xec\xae\xe7\x79\xf3\x38\xa6\xe7\x39\xd8\x81\xd9\x20\x0b\xa5\x4c\x50\x9b\x3f\x3e\xc8\x05\x85\x2f\x14\xce\x28\x1c\x93\x23\xb8\x0a\xe9\xd0\x90\x38\x0d\x94\xb4\x09\x7e\xed\x4f\x10\xf2\x85\xc2\x05\x85\x63\x0a\x67\x1d\xf6\x8e\x76\x1e\xec\xae\xf5\x68\xb1\x67\x01\xe8\xd9\x1b\x6c\xcf\x45\x07\x31\x02\xb6\x19\xad\xbf\x25\x59\x8b\x3d\xea\xab\x05\xd9\x5e\xd6\x5c\x19\x48\xbc\x9b\x60\x33\x14\xc0\x65\xbb\xc8\x34\xac\x31\x8e\xa2\x41\xab\x1a\xcb\x55\x40\xfa\x66\x45\xed\xd6\xb8\x69\xed\xfe\x27\xfb\xa6\xb8\x4c\x08\x0d\x61\x3a\x7f\x3f\xde\xf6\x24\xe5\xcd\x0d\xde\xff\x7b\xaa\x44\x43\xd4\x74\x18\x46\xf0\x3d\x41\x04\x6f\x42\x08\xfe\xa4\x00\x4d\x39\x5e\x01\x01\x92\xee\xf2\xf3\x99\x76\xb1\x6b\xb1\x65\x12\xbc\x2b\xf1\x83\xde\x3c\x6b\x7e\x7a\xa7\x76\x86\x0e\xe9\x6b\x11\x4e\x99\xee\x62\xd5\x0e\xd2\xe1\x6d\x62\xa8\x92\xae\x4d\xed\x0c\x1d\x0a\x73\x50\xb9\xdd\x63\xb4\x19\x90\xdd\x15\x4a\xb3\xb2\xe4\x72\x71\x04\x57\x87\x13\x0a\x87\x93\x6b\x58\xb7\x0c\xbb\x86\x2c\x85\xee\xa2\xd3\x77\x87\xad\xd0\x84\x8b\xd1\x02\x95\x1f\xa5\x61\xca\x86\xd2\x7a\x8b\xbf\x6b\x5e\xaa\xb0\x9c\x4d\xf5\xd6\xbd\x56\xc3\x04\xc2\x1f\x90\x83\xc4\x25\xfc\x7b\x7a\xf2\xd1\x39\x7d\x8e\x3f\x6a\xb4\x2e\x69\xee\x53\x06\x7f\x64\x4a\xfa\x0b\xe9\x83\xe3\xdd\x73\x38\xb3\x82\xae\x50\xfe\x56\xd9\x0d\x74\xcd\x8c\xc5\xc4\xdd\x72\x9b\x19\xb4\x5a\x49\x1b\x54\xb0\xb3\x87\x05\x73\xc5\x2d\x24\x18\x2a\x5e\x28\x69\x95\xc0\x4c\xa8\x45\x42\x8e\xcf\xcf\x3f\x9f\x13\x0a\x8d\x28\x7e\x3b\xef\x1f\xf0\x6c\x4e\x9d\xc4\xa7\x12\x4e\x24\xdf\xd4\x2d\xca\x0f\xc7\x97\xef\x3f\x12\xea\xaf\x00\xa1\xe3\x61\x7d\x1a\x65\x42\x16\xe8\x9a\xef\x14\x9c\xa9\x43\x48\x6f\xb4\x28\xcb\xbe\x16\xc3\x04\x9a\xf8\xeb\x74\x1a\xff\x17\x00\x00\xff\xff\x0d\x8c\xfb\x87\xeb\x0b\x00\x00")

func staticJavascriptMarc034JsBytes() ([]byte, error) {
	return bindataRead(
		_staticJavascriptMarc034Js,
		"static/javascript/marc-034.js",
	)
}

func staticJavascriptMarc034Js() (*asset, error) {
	bytes, err := staticJavascriptMarc034JsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/javascript/marc-034.js", size: 3051, mode: os.FileMode(420), modTime: time.Unix(1515890719, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticCssMarc034Css = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xdb\xae\xdb\x2a\x14\x7c\x36\x5f\xb1\x8e\xa2\x23\xed\x46\x72\xe2\xec\x24\xdd\x15\x51\xa5\xfe\x47\xd5\x07\x0c\x38\x41\x05\x16\x02\xdc\x38\x8d\xfc\xef\x15\xbe\xc5\xb9\xb4\xbc\x31\xcc\xba\x30\x33\xeb\x25\x14\xdb\xdd\x8a\x87\x00\xcb\x35\x21\x25\x8a\x0b\x5c\x09\xa4\x53\xa1\x8d\x79\xc5\x8c\xd2\x17\x0a\x81\xd9\x90\x07\xe9\x55\x75\x20\x2d\x21\x1c\x85\x84\x2b\xf4\x44\x8e\x1a\x3d\x85\xd2\xe3\xd9\x76\xaf\xca\xba\x3a\x7e\x8f\x17\x27\xbf\x46\xd9\xc4\x1f\x70\x25\x59\x06\x8e\x09\xa1\xec\x91\xc2\x6a\x2b\x4d\xc7\x5b\x2f\x49\x59\xc7\x88\xb6\xe7\x86\xba\x34\xaa\x63\xc3\x70\x4a\xf4\x42\x7a\x0a\x01\xb5\x12\x10\x4f\xca\x1e\x48\x76\xeb\xb3\xeb\xfb\xa4\xc5\x57\x1e\xcf\x63\xa1\x61\xfe\xa8\x6c\xae\x65\x15\xe9\xf6\xff\xc3\x1d\xe8\xd5\xf1\xf4\x8c\x46\x74\x74\xf3\x80\x95\x18\x23\x1a\x0a\x09\x6f\x09\x59\x39\x66\xa5\x1e\x47\x9c\x95\x88\x27\x0a\xbb\xe2\x56\xd4\xe4\xcf\xa0\xb2\xf9\x49\x76\x13\x61\x57\x14\xae\x99\x91\x5f\xe1\x95\x46\x16\x69\xda\x7b\x00\x86\xff\xcf\xbf\xdf\x12\xe2\xbc\x1c\xf7\x98\xb4\xd8\x24\x29\x6e\xb6\x05\xf5\x5b\xd2\x60\x98\xd6\x87\x17\x66\x1a\xb4\x18\x1c\xe3\xb2\x6b\xb7\x28\x4b\x6c\x64\x80\x2b\x99\xdf\xea\xf4\xd7\xac\xd7\x82\x42\xb7\xe3\x4d\xf9\xee\x7a\x4f\xd6\x2a\xf1\xb5\x0a\x31\x0f\xf1\xa2\x65\x9e\x2c\xa5\x16\xad\x3c\x8c\x6d\x46\x49\x57\xdb\xfd\x10\x80\x85\x61\x6e\x9c\x6b\x98\xe7\x93\xf5\xc3\xdc\xc9\xa7\x07\x53\x0a\xd7\xc0\x7f\xca\x38\xf4\x91\xd9\xd8\xb7\xf2\x6c\x0a\x00\xfe\x92\xbe\xd2\x78\xa6\x10\xb8\xc7\x49\x83\x92\xf1\x9f\x47\x8f\xb5\x15\x79\x1f\xd9\x05\xe7\x7c\x28\x96\xa1\xd6\x71\xd4\xe0\x9b\x91\x42\x31\x78\x4b\xfe\x0d\xa6\x6e\xdf\x0b\xd7\x7c\x02\x66\x05\xbc\xcd\xbc\xfe\xf8\xfc\x25\xc1\xd7\xbf\x65\x7e\x46\x7d\xff\xd8\x3f\x6f\x3d\xe4\xfc\x3e\x5c\x00\xd9\x6c\xf2\x73\x5d\xf6\xcf\xae\xd9\x3c\x75\x9b\xfd\xa3\x54\xd9\x3c\x7c\x2f\x9e\xfb\x0c\x26\xdf\x66\xf8\xb8\x57\x4b\x5a\xf2\x27\x00\x00\xff\xff\x70\xab\x67\xe1\x32\x04\x00\x00")

func staticCssMarc034CssBytes() ([]byte, error) {
	return bindataRead(
		_staticCssMarc034Css,
		"static/css/marc-034.css",
	)
}

func staticCssMarc034Css() (*asset, error) {
	bytes, err := staticCssMarc034CssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/css/marc-034.css", size: 1074, mode: os.FileMode(420), modTime: time.Unix(1515889949, 0)}
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
	"static/javascript/marc-034.js": staticJavascriptMarc034Js,
	"static/css/marc-034.css": staticCssMarc034Css,
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
	"static": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"marc-034.css": &bintree{staticCssMarc034Css, map[string]*bintree{}},
		}},
		"javascript": &bintree{nil, map[string]*bintree{
			"marc-034.js": &bintree{staticJavascriptMarc034Js, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
