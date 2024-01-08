// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1565447861_initial_schema.down.sql (28B)
// 1565447861_initial_schema.up.sql (140B)
// doc.go (377B)

package migrations

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __1565447861_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x71\xf5\x71\x0d\x71\x55\x08\x71\x74\xf2\x71\x55\xc8\x2d\x4b\x29\x8e\xcf\x4d\x2d\x2e\x4e\x4c\x4f\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\xc2\x3a\x18\x9b\x1c\x00\x00\x00")

func _1565447861_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1565447861_initial_schemaDownSql,
		"1565447861_initial_schema.down.sql",
	)
}

func _1565447861_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1565447861_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1565447861_initial_schema.down.sql", size: 28, mode: os.FileMode(0644), modTime: time.Unix(1701858009, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x92, 0x55, 0x8d, 0x3, 0x68, 0x1a, 0x9c, 0xd7, 0xc7, 0xb4, 0x5a, 0xb1, 0x27, 0x47, 0xf4, 0xc6, 0x8d, 0x85, 0xbb, 0xae, 0xb6, 0x69, 0xc5, 0xbc, 0x21, 0xba, 0xc0, 0xc6, 0x2a, 0xc8, 0xb2, 0xf7}}
	return a, nil
}

var __1565447861_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xc8\x2d\x4b\x29\x8e\xcf\x4d\x2d\x2e\x4e\x4c\x4f\x2d\x56\xd0\xe0\x52\x50\x50\x50\xc8\x4c\x51\x70\xf2\xf1\x77\x52\x08\x08\xf2\xf4\x75\x0c\x8a\x54\xf0\x76\x8d\xd4\x01\x4b\xa4\x17\xe5\x97\x16\xc4\xc3\xa4\xfd\xfc\x43\x14\xfc\x42\x7d\x7c\x20\x72\x25\x99\xb9\xa9\xc5\x25\x89\xb9\x05\x0a\x9e\x7e\x21\xae\xee\xae\x41\x68\xf2\x49\xf9\x29\x95\xa8\xfa\xb8\x34\xad\xb9\x00\x01\x00\x00\xff\xff\xae\x9b\x57\x51\x8c\x00\x00\x00")

func _1565447861_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1565447861_initial_schemaUpSql,
		"1565447861_initial_schema.up.sql",
	)
}

func _1565447861_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1565447861_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1565447861_initial_schema.up.sql", size: 140, mode: os.FileMode(0644), modTime: time.Unix(1701858009, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x21, 0x13, 0xbf, 0x64, 0x18, 0xf7, 0xe2, 0xd8, 0xb5, 0x7d, 0x8, 0xf1, 0x66, 0xb9, 0xb3, 0x49, 0x68, 0xe2, 0xa2, 0xea, 0x90, 0x11, 0x70, 0x9c, 0x15, 0x28, 0x3f, 0x3f, 0x90, 0x3c, 0x76, 0xf}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\xbb\x6e\xc3\x30\x0c\x45\x77\x7f\xc5\x45\x96\x2c\xb5\xb4\x74\xea\xd6\xb1\x7b\x7f\x80\x91\x68\x89\x88\x1e\xae\x48\xe7\xf1\xf7\x85\xd3\x02\xcd\xd6\xf5\x00\xe7\xf0\xd2\x7b\x7c\x66\x51\x2c\x52\x18\xa2\x68\x1c\x58\x95\xc6\x1d\x27\x0e\xb4\x29\xe3\x90\xc4\xf2\x76\x72\xa1\x57\xaf\x46\xb6\xe9\x2c\xd5\x57\x49\x83\x8c\xfd\xe5\xf5\x30\x79\x8f\x40\xed\x68\xc8\xd4\x62\xe1\x47\x4b\xa1\x46\xc3\xa4\x25\x5c\xc5\x32\x08\xeb\xe0\x45\x6e\x0e\xef\x86\xc2\xa4\x06\xcb\x64\x47\x85\x65\x46\x20\xe5\x3d\xb3\xf4\x81\xd4\xe7\x93\xb4\x48\x46\x6e\x47\x1f\xcb\x13\xd9\x17\x06\x2a\x85\x23\x96\xd1\xeb\xc3\x55\xaa\x8c\x28\x83\x83\xf5\x71\x7f\x01\xa9\xb2\xa1\x51\x65\xdd\xfd\x4c\x17\x46\xeb\xbf\xe7\x41\x2d\xfe\xff\x11\xae\x7d\x9c\x15\xa4\xe0\xdb\xca\xc1\x38\xba\x69\x5a\x29\x9c\x29\x31\xf4\xab\x88\xf1\x34\x79\x9f\xfa\x5b\xe2\xc6\xbb\xf5\xbc\x71\x5e\xcf\x09\x3f\x35\xe9\x4d\x31\x77\x38\xe7\xff\x80\x4b\x1d\x6e\xfa\x0e\x00\x00\xff\xff\x9d\x60\x3d\x88\x79\x01\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 377, mode: os.FileMode(0644), modTime: time.Unix(1693317681, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xef, 0xaf, 0xdf, 0xcf, 0x65, 0xae, 0x19, 0xfc, 0x9d, 0x29, 0xc1, 0x91, 0xaf, 0xb5, 0xd5, 0xb1, 0x56, 0xf3, 0xee, 0xa8, 0xba, 0x13, 0x65, 0xdb, 0xab, 0xcf, 0x4e, 0xac, 0x92, 0xe9, 0x60, 0xf1}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"1565447861_initial_schema.down.sql": _1565447861_initial_schemaDownSql,
	"1565447861_initial_schema.up.sql":   _1565447861_initial_schemaUpSql,
	"doc.go":                             docGo,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"1565447861_initial_schema.down.sql": {_1565447861_initial_schemaDownSql, map[string]*bintree{}},
	"1565447861_initial_schema.up.sql":   {_1565447861_initial_schemaUpSql, map[string]*bintree{}},
	"doc.go":                             {docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
