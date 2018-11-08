// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// web/favicon.ico
// web/index.html
// web/script.js
// web/style.css
package main

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

var _webFaviconIco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4d\x68\x53\x4b\x18\xfd\xf2\xda\xf7\xfa\xde\xe3\xc1\x0b\xbc\x87\x1b\x85\x06\x5c\x28\xb8\xd3\xda\x9d\xd6\x9d\x3b\x05\x7f\x40\x17\x4a\x3a\xdf\x6d\x8c\xa5\x6e\x75\x55\x7a\x41\xc5\xae\x74\x11\x15\x42\xa1\x21\xbd\x73\x26\xa5\xb4\x4b\xbb\x71\x51\x41\xeb\x1f\x0a\xe2\x46\xd0\x8d\x52\x70\x53\x90\x56\x41\x5b\xac\xfd\x64\x7a\x13\x49\x43\x1a\x7a\x5b\xee\xb4\x8b\x1c\x38\x5c\x66\xe8\xe9\xc9\xcc\x99\xf9\xf2\xdd\x10\x25\xa8\x85\x92\x49\xfb\x4c\x51\xb6\x95\xa8\x93\x68\xd5\x78\xb8\x95\x28\x95\x0a\xc7\xa3\x49\xa2\x3f\xfe\x22\xda\x6b\xff\x86\x88\x52\x14\xce\x37\x42\x10\xa8\x3d\x00\x9f\x8e\x42\xab\xa9\xe8\x01\xaf\x1f\x60\x89\x46\xaf\xdf\x6a\x8d\xc9\xfc\x0f\xf4\x1c\x2d\x95\xd4\xb1\x28\xb4\x1a\xab\xd5\x9a\x55\xf9\x7f\xbe\x03\xbc\xaf\x51\x3e\x83\xd5\x96\xf5\xf3\x53\x53\x7e\x2b\xc0\x7d\x1b\xd4\xdb\x71\xce\x18\x7e\x15\x55\x1f\x04\x76\x2d\xfc\x62\x03\xfc\x0e\x78\x27\x1a\x06\x5b\x07\x85\x42\xfa\x4f\x80\xef\x02\x2a\x1d\x55\xab\xb5\xb7\x0b\xe0\x09\xad\x7b\x0e\x46\xd5\x06\x81\x77\x08\xe0\xb1\x91\x91\xec\x8e\xa8\x5a\xc0\xcb\x00\xea\x76\x3e\x9f\xf9\x3d\x8a\x6e\x72\xf2\x52\x9b\xd6\x7c\x67\x13\x6b\x1d\x77\xbd\x56\x63\xb8\x17\xe0\x41\x7b\x1e\xa3\x6a\x2d\x4a\xa5\x74\xe3\xcb\x5e\x03\x3f\x11\xb2\x1e\x64\x9e\x48\x66\x88\xe4\x09\x91\xdc\x22\x5a\xee\x22\x5a\x1a\x20\x5a\x10\xa2\x79\x21\xfa\x26\xcd\x3a\xd3\xac\x33\xf5\xd1\xac\x33\xeb\x43\xb3\xce\xac\xbf\xce\x58\xc3\x23\xeb\xa8\x33\xdb\x05\x36\x57\xc0\xbb\x1f\x2f\x79\xbc\x81\xff\x6c\xf4\x7a\x19\x99\xb3\xf5\xbc\x0b\x85\x74\x12\xe0\x45\x07\xfe\x8b\xb9\xdc\xd9\xff\x6a\xfd\xed\x9d\x00\xbc\x03\x71\x50\x6b\x4b\xee\x08\x02\xee\x28\x16\x55\x67\x2e\x77\x7e\x67\xb5\x77\x58\x6b\x6d\x2d\x89\x8f\x5a\x7b\x17\xb4\xe6\x6c\x10\xa8\x8b\xc5\xa2\xea\x1b\x1a\x3a\x77\xa6\x2a\xf7\x47\x0e\xf6\x7d\x15\x83\x40\x3d\xab\xef\xef\xbd\x04\x3c\x5f\x6b\x35\x1d\xb3\xff\xf3\x3a\xfe\x33\xf9\x7c\xe6\x6f\x3b\x67\x6b\x3a\xa0\xde\xb8\xf4\x37\x86\xa7\x6a\xee\xe3\x84\xe3\xfd\xb7\xf7\xef\x70\x78\x26\xbb\xf7\x03\xfc\x25\x2e\x7f\xad\xd5\xd3\x35\xce\xdf\x32\xc0\x1f\x01\xfe\x11\x67\xfe\x5a\xab\xc7\x5b\x79\xfe\x01\x9e\xae\xf8\x6b\xcd\x97\x01\x95\x8f\x9b\xc6\xf0\x10\xc0\x9f\xec\x1e\x07\x01\x5f\x5f\xeb\x7b\x20\x0e\x94\x4a\xd9\x7f\x6c\xef\x11\x9e\xa9\xee\x53\x2e\xbd\x47\x47\x33\xbb\x01\xef\x35\xc0\x6f\x81\x9e\x7d\x2e\xbd\xed\xfb\x40\x79\xcf\xef\x0d\x0f\xa7\x93\xae\x7c\x45\x28\x01\xf0\x95\xb0\x1f\xe7\x41\xdf\xf7\x7f\x73\xe5\x3d\x36\xd6\xfb\x2b\x6b\x63\xd4\x49\x57\xbe\x54\x93\xb5\x31\x5b\x97\xb5\xed\x67\x5c\xf9\x6e\x93\xac\x3f\x6f\xe4\xbd\x6f\x33\x68\x66\xdd\xcc\xda\x05\xb6\x2a\x6b\x5a\xe9\x0d\x54\xb6\xfc\x9b\x8a\xef\x32\xeb\x0a\xc2\x3e\xd0\x3b\xee\xda\xd7\x15\x44\x44\xde\x13\xb5\x59\x3e\x20\x6a\xb1\xf4\x89\x12\x15\x5e\x23\x4a\xdc\x2c\x3f\xab\xe7\x45\xe4\xa1\x88\xdc\x10\x91\x7f\x45\x96\xda\x45\x16\x06\x44\xe6\x44\xe4\x83\x88\x5c\x0d\xa7\x97\xdb\x45\x96\xba\x44\xe6\x06\xc2\xe9\x2a\x85\x74\xc9\x0a\x7e\x06\x00\x00\xff\xff\x65\xe4\xf4\x34\xae\x19\x00\x00")

func webFaviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_webFaviconIco,
		"web/favicon.ico",
	)
}

func webFaviconIco() (*asset, error) {
	bytes, err := webFaviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/favicon.ico", size: 6574, mode: os.FileMode(420), modTime: time.Unix(1538664377, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x5b\x77\xe3\xb6\xf1\x7f\xff\x7f\x8a\xf9\x23\xc7\xf5\xee\x03\xc1\x9b\xee\x2b\x2a\x75\x9c\xd5\x76\xdb\x24\xc7\xeb\xdd\x93\xa6\xa7\xa7\x0f\x10\x39\xa4\x60\x93\x00\x17\x00\x29\x6b\x93\x7c\xf7\x1e\x90\xb4\x2c\x39\x92\xef\x6d\x93\x07\x5b\xb8\x0c\x7e\x98\xf9\xcd\xe0\x36\x9c\xfe\x7f\x22\x63\xb3\x2e\x11\x96\xa6\xc8\x67\xff\x37\x6d\x7f\x00\x00\xa6\x4b\x64\x49\x5b\x6c\xaa\x86\x9b\x1c\x67\x1f\x4f\x7e\x78\x0b\x7f\xc7\x85\x8e\x99\x98\xba\x6d\xdb\x8d\x4c\x81\x86\x41\xbc\x64\x4a\xa3\x89\x48\x65\x52\x67\x44\xb6\xba\x73\x2e\x2e\x41\x61\x1e\x11\x6d\xd6\x39\xea\x25\xa2\x21\xb0\x54\x98\x76\x2d\x34\xd6\x7a\x7b\x80\x8e\x15\x2f\x0d\x68\x15\x47\x64\x69\x4c\xa9\x27\xae\x5b\x2a\x49\x53\x29\x0c\x5b\xa1\x96\x05\xd2\x58\x16\xae\xc2\x1c\x99\x46\xed\xd6\x7d\x1a\x52\xdf\xbd\xd0\x2e\xcb\x73\x7a\xa1\x09\x70\x61\x30\x53\xdc\xac\x23\xa2\x97\x2c\x1c\xf5\x1c\x3c\xf9\xf1\x92\xbf\x3d\xf7\xd2\xef\xdc\xf5\x47\xfe\xf9\xe3\x30\xf9\xa9\x1a\x7d\xfa\xae\x94\xe7\xa3\x64\xfc\xb7\xf3\x2f\xef\xff\x61\xde\x79\x9f\xbe\x0c\x4b\x1e\xf4\x3e\x67\x1f\xde\xbf\xaf\xca\xd2\x4b\x45\xf0\xd3\x89\xff\x97\xb1\x97\x9e\x11\x88\x95\xd4\x5a\x2a\x9e\x71\x11\x11\x26\xa4\x58\x17\xb2\xd2\x64\x36\x75\x5b\x85\x3b\x02\xdd\x1b\x06\xa7\x0b\x99\xac\xb7\x89\x62\x5c\xdc\x54\x5b\x5b\x31\x36\x5c\x0a\x88\x73\xa6\x75\x44\x62\x29\x8c\x92\xf9\x36\x1d\x87\x44\x13\xac\x79\x8c\xfb\x24\x1b\xe9\xda\xd1\x98\x63\x6c\x60\x22\x4b\x3b\x4c\x47\xc4\x7a\x4f\xa0\xd2\x04\x26\x1a\x99\x8a\x97\x6c\x91\x63\x44\x52\x96\x6b\x24\x30\x31\x2c\xcb\x76\x5b\x52\x9e\x1b\x54\x3b\x6d\xb5\x53\xc8\xc4\xfa\xb2\x43\xa5\x1d\xa8\xa5\xe1\x7a\xca\x3d\xba\xbb\x9d\xf2\xf7\x9b\xa5\xd1\x18\x2e\xb2\x83\x76\xdd\x12\x6f\xd5\x38\x20\xdc\x86\x1f\x5b\x60\x3e\x9b\xf2\xeb\x11\x29\x83\x94\x39\x78\x55\x32\x91\x38\x4c\x29\xb9\x6a\x7c\xc8\x67\x70\x8e\x5a\xe6\x95\xc5\x9b\x4c\xdd\x76\xd8\x61\xd8\xda\x51\x4c\x64\x08\x93\x9a\xe5\x15\x52\xbd\x16\xf1\x0d\x29\x6a\x83\x44\xc0\x48\x99\x1b\x5e\x46\x64\x29\x6b\x54\x9b\xba\x93\x70\x15\x91\x85\x34\x46\x16\x2d\x79\x0d\xde\x01\xab\x0f\xf3\xf7\x92\xa4\xe8\x4a\x74\x54\x7c\xa3\x78\xb6\x34\x02\xb5\x7e\x2e\x15\x8b\x0d\xd2\x1f\x8a\x0a\x96\x5c\x54\xda\x74\x6c\x9c\xda\x55\xc9\xb4\x79\x2e\x17\x71\x87\xf3\x87\x61\x42\x59\x2a\x4a\x96\xa3\x31\xb8\xe1\x22\x97\xea\x01\x44\x34\xfd\x90\x4a\x65\x37\xb5\x5c\xaa\x3b\x26\x6e\xe4\x6f\xf1\x9f\xf2\x3c\x77\x12\xc5\xcb\x76\xda\xc3\xf3\xdc\xab\x08\x17\x65\x65\xc0\x9e\x76\x11\x89\x97\x18\x5f\x2e\xe4\x15\xb9\x9e\xcc\xc8\x2c\xcb\xf7\x6d\x6b\xad\xce\xc0\x93\x8d\xfa\x7b\x2c\x7a\x79\x16\x74\x9d\x6d\x36\xc2\x3a\x73\xb8\xc8\xb9\x40\xc7\x69\x39\xe9\x1c\x61\x8b\x2b\xc7\x1f\x10\x60\x8a\x33\x67\xc9\x93\x04\x45\x44\x8c\xaa\x90\x40\xc2\x0c\x73\x4a\x85\x29\xbf\xb2\x54\x76\x0d\x3c\x96\x22\x22\xd7\x8e\x04\x25\xed\x8e\xce\x8b\x8c\xc0\x55\x91\xdb\xa3\xc1\x1e\xb3\x13\xd7\x5d\xad\x56\x74\x15\x52\xa9\x32\x37\xf0\x3c\xcf\xd5\x75\x46\xa0\xe6\xb8\xfa\x46\x5e\x45\xc4\x03\x0f\xfa\x7e\x60\xff\x3a\xdc\x94\x39\x3c\xd0\x75\x16\x11\x32\x9b\x26\x98\xea\x99\x3d\xed\x91\xa9\x77\x8a\x25\x1c\x85\x69\x08\xcc\x08\x64\x5d\xfd\x93\x62\x42\xa7\x52\x15\x11\x51\xd2\x30\x83\xaf\xfa\xaf\xc9\x6c\xaa\x8d\x2c\xa1\xb9\x0b\xd8\x2b\x81\x2c\x9d\x86\xac\x09\x7c\x35\x0f\x42\xef\x74\x48\x40\xa6\x69\x73\xb9\xf0\x8e\x88\x7b\x87\x78\xd0\x3f\x1d\x9f\x0c\x6e\xc4\x7d\xaf\x1d\xe0\xee\x6a\x35\x9b\xba\xad\xb2\x25\x33\x4b\xb0\xa1\x16\x91\x4a\xe5\xaf\xbe\xca\x5e\x13\x48\x22\xf2\x7d\xe0\xf5\x68\x08\xfd\x53\xdf\xeb\xd1\x31\x04\x3d\xda\xb3\xff\x46\xe0\xb7\xed\x34\x80\xc0\x0b\x69\x2f\x76\xc2\x21\xf8\xa3\x21\xf8\xa1\x4f\x87\x10\x06\x03\x2b\xd8\x1f\xd1\x11\x84\xde\x80\x0e\xa1\xe7\xd3\xc0\xb1\x8d\x03\x9f\xf6\x9c\x7e\x8f\x0e\xa0\x17\xd0\xbe\x33\xf6\xe9\xd0\x09\x42\xea\x3b\xbd\x3e\xed\xc1\x98\x8e\x9d\xf1\xc8\x8a\x79\x5d\x69\x39\x1c\xd3\x61\x1c\xf6\xe9\x08\x3c\x18\xf4\xe8\xc8\x09\xc6\x74\x60\x4b\x63\x67\xd0\xa7\xe1\x69\xdf\xf7\x69\x1f\xc6\x43\xea\x43\x38\x18\x51\xdf\x09\x06\x56\xd3\x56\xbf\x2f\xdf\x8f\x07\x10\x06\x5e\xec\xf8\x43\x3a\x04\xcf\x09\x03\xc7\xef\xd1\xd0\xfe\x86\x81\xee\x8a\xd0\xd4\x20\x0c\xc0\x36\x40\xdb\xd0\x15\xdb\x8e\x2f\x85\x6d\x09\x46\xcf\xc7\xf1\x83\x91\x33\xe8\xbd\x08\x0e\xbc\x00\x8e\x5d\xba\xd6\xf7\xf6\xee\x56\x67\x4f\xdc\x5d\xfe\x4b\x27\x51\xca\x73\xb4\xd7\x15\xa9\xae\x8f\xa3\xb9\x54\x05\x7b\xc8\x61\xb4\xbd\xf5\xd9\xd0\x97\xf7\xef\x7b\x76\x36\x3b\x80\x80\x60\x85\xbd\xf6\x6d\xea\x76\x25\x5f\xd7\x9c\x32\x49\x09\x34\x47\x5c\x44\x6c\x79\x67\x77\xdc\x91\xda\x6f\x4f\xdb\xe3\xf2\xd9\x54\x97\x4c\xcc\xce\xbe\x9d\x4f\xdd\xa6\xf4\xbf\x34\xea\xa2\xc4\x6c\x63\x55\x53\xd9\x6f\x56\xd7\x75\xcb\xae\x98\x15\xa8\xd8\xb6\x55\x7f\x3d\x7b\xfb\xee\x77\x60\x56\x29\x6e\xac\xb2\xe5\x03\xbe\x12\x7b\x6c\xe2\x05\xcb\x70\xc7\x51\x3f\xfc\x1e\x2c\x32\x3c\xbd\x09\xbf\xa6\xb2\xdf\xa6\xae\x6b\x5f\x00\x72\x51\x4b\x1e\xef\xd8\xf6\xe9\xfd\xfc\x41\x51\xf8\xb4\x75\x0f\x93\xee\xc0\xfa\xb9\xe6\x9a\x2f\x78\xce\xcd\x7a\x02\xaf\x6e\x5b\x0e\x51\x04\xc7\x65\x92\x1e\xc3\x2f\xbf\xc0\xde\x4e\x1b\x7d\xc7\xaf\xe1\x6b\x38\x6e\x80\x72\x3c\x86\x09\x1c\xb7\xb7\x80\xe3\x5f\x1f\xbf\xbf\x94\xa8\x62\x14\x66\xe3\x68\xb0\x61\x0b\x1f\x2a\xd6\x68\xf8\xcc\x1b\xef\xe7\x16\xe6\x3f\x7e\xe1\x7d\xc4\x03\x93\x35\xb5\x83\xef\xcb\x45\x65\x8c\x14\xbf\x7d\x1a\x35\x0f\xe6\x8e\xa1\x0f\x15\x8f\x2f\xe1\x4c\xa1\xbd\x1a\x4d\xdd\x6e\xc8\x5d\x78\xd7\x60\xa5\xe2\x05\x53\xeb\x03\x41\xa9\x79\x26\x98\xa9\xd4\xb5\x27\x3e\xc6\x4c\xc0\x19\xcb\xf0\xf0\x1c\x8f\x30\xdc\x4e\xa1\xed\xa2\xd3\x4b\xb9\xb2\xb7\xc1\x0c\x35\xcd\x51\x64\x66\x09\x33\xf0\x0e\x11\x92\xf0\x7a\x1b\xe1\xae\x00\x6b\xd7\x0e\x2f\xb2\xdf\xe4\x6e\xae\xd6\x34\xa9\xe2\x4b\xfb\x97\xc9\x26\x79\xc3\x2b\xf7\xeb\x2a\x6a\x44\x8e\xc2\x93\xa3\x60\x7e\x14\xcc\x79\x91\xd5\x81\xe3\x3b\x29\xd5\xb1\xe2\x8b\x84\x69\x8d\xc6\xde\xc5\x8b\xb6\xf3\x28\x98\x27\x32\xae\x0a\x14\xe6\x28\x98\xfb\x7e\x10\x8c\xfb\x5e\xd8\x3f\x0a\xe6\x6d\x5a\x86\xe5\x47\xc1\x3c\xf6\x12\x8c\x99\xcf\xfc\xd4\xca\xf4\x06\xbd\x91\xe7\x8f\x42\xef\x4f\x69\xe4\x37\xb9\x1a\xab\xe3\xc3\x36\x2d\x83\x57\x66\xb3\xcd\x58\x37\x80\xff\x00\xeb\x77\xfd\x6a\x78\x81\x5d\x7e\xe1\xae\xb9\xa7\x6e\xc2\xeb\x17\x09\xeb\x47\xfb\xf7\x81\x01\x9a\xc8\x95\xc8\x25\x4b\xba\xd0\xfc\xb6\xab\x3e\x3a\x32\x0f\x35\xdf\x32\xa7\x6c\x17\xd7\xbe\x3c\xd8\x56\x40\xc6\x4a\x96\x87\xec\xb2\x61\xd8\xe4\x18\x0f\x43\xdd\x86\x33\xb2\x24\xd7\xef\x8a\x25\xf2\x6c\x69\x26\x10\x7a\xe5\xd5\x1b\x6b\xf4\x5e\x07\xdd\x46\xc8\x31\x35\x1b\x88\x15\x4f\xcc\x72\x02\xc1\xa3\x10\xda\x14\x5a\x73\x62\xdc\xd2\x24\xf0\x2c\x10\x74\xa8\xa1\xd7\xc2\x1e\x8e\x47\x0b\xda\xbc\xd2\x12\xae\x5a\xc8\xc6\x42\xa7\xd1\xf1\x2e\x75\xee\x1a\xfd\xe4\x81\x4e\x93\x07\x7a\xda\xf0\xa7\x6b\xfc\x8c\x49\xdb\x43\xe9\x19\x6c\xdd\x9c\x6a\x4f\x9e\xfc\x7e\xf5\x1f\x18\x55\x0f\x00\xda\x92\xbe\x47\xf3\x3d\xcd\x7b\x56\xf5\xd4\xdd\x4d\x77\xef\x4d\xeb\x57\xa2\xbc\xcc\x9a\xf3\xa0\xae\xf0\xcf\x39\x33\xd8\x26\xdd\xb6\x33\xea\x0f\x1b\xdc\xe5\x9e\xef\xc7\x98\xad\xb8\x48\xe4\x8a\xb6\x4f\x2a\x1d\xfd\xfc\xeb\xd3\xa6\xcb\x79\x82\xca\x89\x65\x51\x4a\x81\xe2\x11\x13\xff\x58\xe1\xb9\xbd\xe4\x44\xbb\x8a\xfc\x93\xec\x83\x25\xff\x7a\x93\x60\x8e\x06\x61\x57\xfa\xcd\x3d\x4a\xb7\x65\x7a\xb1\xe7\xfb\x44\xfb\x51\x62\xea\xb6\x1f\x7d\xfe\x1d\x00\x00\xff\xff\xc4\xbe\x88\x4c\x0c\x1a\x00\x00")

func webIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_webIndexHtml,
		"web/index.html",
	)
}

func webIndexHtml() (*asset, error) {
	bytes, err := webIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/index.html", size: 6668, mode: os.FileMode(420), modTime: time.Unix(1538668848, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webScriptJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8a\x41\x0a\x83\x40\x0c\x45\xf7\x39\x45\xc8\x4a\xc1\x7a\x00\xa5\x97\x68\xc1\x4d\xe9\x22\x68\x10\x41\x33\x83\x13\xeb\xa2\x78\xf7\x32\x23\x9d\x2e\x9a\xd5\xcb\x7b\xbf\xdb\xa4\xee\xdd\xe2\x9d\x8a\x5a\x41\xaf\x4b\x90\x59\x7a\xa3\x0a\xbb\x4d\xee\x89\xeb\x4c\x65\x0b\x7f\xfb\x95\x75\x94\x73\x7e\x8b\x58\xb6\x00\xfb\xa4\x83\xdb\x6b\xf6\x1e\xaf\xa8\xb2\xc7\x58\xbc\x01\x11\x51\xe6\x06\x69\xe1\x49\xa9\x4a\xff\xc0\xc6\x0d\x9e\x2d\x9e\xf3\x36\x39\x0d\x51\x65\x77\x54\x19\x3d\x8f\x12\x1a\x7c\x10\xd3\xf3\x67\x43\xcf\xaa\xb2\xa6\x60\x12\xec\xdb\x0e\x38\x4a\xf8\x04\x00\x00\xff\xff\xb1\xd5\x67\xf9\xe2\x00\x00\x00")

func webScriptJsBytes() ([]byte, error) {
	return bindataRead(
		_webScriptJs,
		"web/script.js",
	)
}

func webScriptJs() (*asset, error) {
	bytes, err := webScriptJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/script.js", size: 226, mode: os.FileMode(420), modTime: time.Unix(1538666675, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _webStyleCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x58\x4b\x6f\xe3\xba\x15\xde\xfb\x57\x9c\x66\x10\x38\x46\x22\x85\x96\xed\xdc\x8c\x3c\xce\xa2\x17\x2d\xba\x98\xbb\x2a\x50\x14\xb8\x98\x05\x2d\x51\x32\x3b\x94\x28\x90\xf4\xab\x41\xfe\x7b\x41\x52\xd4\x93\x76\x92\x0e\x66\xe0\x48\xe2\xe1\xe1\x79\x7c\xe7\xc5\x9d\x2a\x58\x5c\x72\x75\x17\x66\xbc\x54\xf8\x48\x24\x2f\x48\x40\x23\x79\xc8\x03\x9c\x28\x7a\x20\x33\x08\xf5\x0b\x2d\x19\x2d\x49\x10\x64\x18\x5e\x21\xa5\xb2\x62\xf8\x1c\x43\xc9\x4b\xb2\x86\xb7\xc9\x44\xf3\x81\xd7\x09\x00\x80\x66\x14\x64\xb8\xa0\xec\x1c\xc3\xcd\xdf\xa9\xc0\xf0\x4f\x5c\xca\x9b\x07\x90\xb8\x94\x81\x24\x82\x66\x6b\x4b\x25\xe9\x7f\x49\x0c\xf3\xa8\x3a\xad\xcd\xd6\x2d\x4e\x7e\xe6\x82\xef\xcb\x34\x86\x2f\x04\x91\x94\xa4\x6b\x48\x38\xe3\x22\x86\x2f\xf3\xaf\xd1\x36\xda\xae\x27\x6f\x93\x2d\x4f\xcf\x50\x9f\x56\xe1\x34\xa5\x65\x1e\x03\x5a\x43\x81\x45\x4e\x4b\xfd\x68\x96\x1a\x29\x33\x46\x4e\x6b\xc0\x8c\xe6\x65\x40\x15\x29\x64\x0c\x09\x29\x15\x11\x6b\xf8\xcf\x5e\x2a\x9a\x9d\x83\x84\x97\x8a\x94\xaa\x59\x30\x0c\x0a\x5a\x06\x3b\x42\xf3\x9d\x8a\x61\x8e\xd0\x61\xa7\x0f\x9f\x14\x98\x96\xf5\xe1\x3d\x79\xb3\x2c\x5b\xc3\x96\x8b\x94\x88\x80\x91\x4c\xef\xa9\x4e\x20\x39\xa3\x29\x7c\x79\x5a\xe0\x15\x5e\xd5\x5a\x5a\x1a\x81\x53\xba\x97\x31\xb4\xda\xf3\x53\x20\x77\x38\xe5\xc7\x18\x10\xac\xaa\x13\x44\xfa\x27\xd0\x6c\x44\xbe\xc5\x77\xe8\xc1\xfc\x0b\xa3\xd5\xcc\xa7\xa1\xf9\xc4\x0f\x44\x64\x4c\xb3\xd8\xd1\x34\x25\xa5\x11\x39\xd4\xea\x09\xce\x64\x2d\xf7\xc0\x32\xfa\x37\x48\xa9\x20\x89\xa2\xbc\x8c\xb5\xc5\xf7\x85\xd9\x19\xf2\x4a\x7f\x1a\xda\x7a\xee\x84\xfb\xb0\xa1\x5b\x5e\x2f\x0c\x6f\x09\x8b\x33\x2a\xa4\x0a\x92\x1d\x65\x29\xbc\xc2\x91\xa6\x6a\x17\xc3\x33\x29\x34\x9a\x2e\x53\xde\x43\xcc\x70\x67\x9f\x91\x3c\x17\x5a\xdd\x79\x67\xe7\x80\x7b\x2d\x77\xa0\x78\x15\x5b\xb1\x3b\xa4\x3d\x7e\x8e\x72\xcb\x95\xe2\x45\x4b\x3c\x09\xb1\xb1\x8d\xec\x82\xdf\x6a\x5b\x7b\xd3\xed\x18\xfb\xdc\x51\x98\xd3\x3d\xcb\x6f\x2d\xf3\xed\x5e\xa9\xc6\xda\x2d\x9a\x2f\x1a\x5e\x07\x51\x0c\xb4\xdc\x11\x41\x95\xfd\x94\xec\x85\xd4\xe1\x52\x71\xda\x02\xd9\x4a\xe0\xc2\xf5\x6a\x94\x21\x84\x8c\xb3\xfa\x22\x99\x14\xd1\x31\xd5\x0c\x5e\x1b\x18\xd7\xf1\xf1\x01\xc5\xe2\x9d\x46\xa7\xde\xda\x95\x20\x5d\x90\x27\xf2\xe4\x21\x0f\x2b\x41\x0b\x2c\xce\xbe\x60\xfb\xba\x4a\xa2\x24\x1a\x88\x6d\x4c\xd2\x45\xc4\x58\x11\xc7\xd3\x2f\xca\xf3\xd7\x2d\xde\x62\xe3\xf0\x3f\xd5\xb9\x22\x9b\x64\x47\x92\x9f\x5b\x7e\xfa\x11\x2a\x9e\xe7\x8c\xf8\x72\xdf\x05\xca\x7b\x30\xf0\xf5\x09\xef\x6c\xee\xdc\x72\x19\x33\x2e\x47\xcc\x9f\x9c\xcb\x5d\x3e\x5a\xe8\xb4\xe1\xe2\x66\xf9\xec\x96\x1b\xd0\x40\x03\x92\x8a\x4b\x6a\xe3\x5a\x10\x86\x75\x4a\xbf\x80\x94\xf7\x54\x89\xb7\x24\xe3\x82\xd4\x1a\x35\x19\xf3\xe6\x66\x78\x0e\xde\x4a\xce\xf6\xaa\x3e\xc7\xe6\x42\xa4\xc5\x35\x21\x60\x9e\x5c\xb8\x98\x17\xaf\x4e\x0b\x7f\x55\xa8\x1d\xef\x4b\xa4\xad\x91\x94\xc0\xa5\x13\xc6\x3c\x67\x5c\x14\x80\xc2\x48\x7e\x40\x4b\x83\x8c\x46\xd7\x4b\x00\xf1\x33\x89\xcd\x07\x92\xfa\x4d\xd6\x88\x52\x4b\xc5\xb0\x22\xff\xbe\xd3\x62\xcf\x4c\x9a\xb6\x3c\xb5\x3a\xfc\x12\xe0\x9a\x93\xfb\x54\x7d\xb0\x7d\x2c\xbd\xff\x1f\x05\xb1\xcd\x43\x51\x75\x32\x15\xaa\x01\x9e\xf3\x5a\xb8\x22\xc5\x95\x62\x7e\x19\xf2\x3d\x8f\x5a\x3c\x98\x08\xf3\x79\x5a\x9f\x8c\xba\x18\xbf\x88\x65\xaf\x99\xde\xcb\x43\xd7\xb6\x76\xcb\x85\x47\xe4\xeb\x7a\x29\x5e\x59\xc2\xae\x26\x3d\x0a\x1b\x17\x3e\xa2\xeb\x8e\xbf\x87\xf7\x61\xd1\x11\x42\xc7\x64\xc3\x1e\xf9\x24\x18\x52\x5c\x3f\x7e\xd4\x22\xb6\x9d\x60\xdd\xe3\x85\xcf\x0e\x17\x36\x45\x35\x05\xf3\xb2\x72\xc3\x60\xba\x56\x07\x00\xe0\x63\x4c\xde\x4d\xfc\x61\x2a\x78\x95\xf2\x63\x19\x34\x11\xd8\x2d\x9f\xf0\x17\x5a\x54\x5c\x28\x5c\xaa\x61\x5d\xee\xae\x78\xc0\xdf\x5f\x7f\xa7\x6b\xf0\x11\xfb\x1b\x08\x1f\x65\xe3\xb7\xde\xe2\xdb\x24\x4c\x18\xc1\x62\x94\x53\x06\x44\x8d\x01\x0a\x52\xee\x8d\xfa\x6d\x5f\x3a\xa2\xff\x9c\x22\x83\x7e\xb9\xb7\xe8\x3a\xd7\xe0\x1c\x03\xde\x2b\x3e\x10\x6b\x12\x1e\xf6\x24\x90\x8c\xa6\x23\xf7\xed\x05\xbb\xbb\x49\xb1\xc2\x31\x2d\x70\x4e\x1e\xe5\x21\xbf\x3f\x15\xec\xe1\x9b\x3c\xe4\x70\x2a\x58\x29\x37\xd3\x9d\x52\x55\xfc\xf8\x78\x3c\x1e\xc3\xe3\x22\xe4\x22\x7f\x8c\x10\x42\x9a\x74\x0a\x07\x4a\x8e\x7f\xe5\xa7\xcd\x54\xe7\x94\x39\x2c\xa7\x2f\xdf\x2a\xac\x76\x20\x95\xe0\x3f\xc9\x66\x7a\x1b\x2d\xac\x32\x53\x48\x37\xd3\x3f\x10\x44\xbb\xf9\xf4\xf1\xe5\x9b\xde\xfd\x72\x33\xeb\x43\x62\x5c\x6d\xb5\x4d\x5b\xd9\x83\x4a\xf0\x84\x48\x39\x84\xa0\x45\xf2\xd0\x17\x9d\x7d\x29\x57\xef\x3a\xe3\x1d\x86\x57\x1a\x8e\xd1\x59\xd7\x02\xe5\x8a\x94\x8a\x73\xa6\x68\xe5\x4f\xae\x1f\x13\xa6\x47\x55\x37\x78\x75\x8f\x3b\x60\x70\xc5\x14\x1a\x30\x19\x65\xc4\x8d\x3c\xfd\xbe\xb0\x3b\x29\x39\xbc\x99\x6e\x51\x6f\xf1\x56\xd1\x71\x8a\xec\x4a\xde\x1f\xcf\x66\x0d\xab\x17\x5a\x56\x7b\x35\x1a\x54\xfa\xfd\x78\x93\x45\xcc\xb8\xb7\xac\xc7\x13\xb3\x5d\x56\x78\x34\x7c\xad\x6c\xe9\x1d\x94\xdd\xc8\xa5\xd7\x5f\x9a\x7a\xbb\xe7\xbe\xd0\x22\x6f\x46\x91\x53\xd0\x3b\xc9\x7c\x69\xda\xb6\xf0\xb7\xa6\xe6\x8f\x7c\xda\xb3\xcc\x7c\x35\x1b\x9c\xa1\x23\xf4\x75\x54\xbe\x8d\xf7\x2a\x41\x74\x60\x8e\xb2\x7e\xe0\x7a\xfe\xe4\xb7\x94\xa4\x64\xd8\x6c\x04\x26\x07\x5c\x4e\x0a\xb7\x8b\xdf\xf5\xa1\x46\x9d\xcd\x74\x89\xa6\x75\xfb\x69\x9f\x7b\xa9\x60\xa9\xff\x4f\x3f\x90\x41\x6e\x17\x7f\xbb\x5d\xfc\x9e\x43\x46\x19\x33\xe9\xc2\x66\xfc\xa9\xf9\x10\xf0\x0a\x27\x54\x9d\x37\x53\x14\x2e\xeb\x4f\x62\xcf\xc8\x66\x4a\x0e\xa4\xe4\x69\x5a\x6f\x37\x29\xc7\xe6\x97\x25\xfa\xbe\x44\x80\xfe\x11\xa1\xef\x08\x22\xf4\x87\x11\xe4\x5f\x11\xfa\x1e\x19\x89\x1e\xed\x86\xc7\xbc\xfe\x2b\x0f\xfa\xe9\x66\xe6\xbb\x46\xb0\x17\x09\xe3\x3b\x84\xe5\xec\xea\x7d\x81\xe0\x55\x6d\x78\x37\x62\x20\xed\x76\xe7\xf2\x27\x74\x01\x6e\x06\xe7\x47\x81\xab\x18\xf4\xef\xe5\x09\xe4\xcd\x1e\x62\x41\xe6\x9b\x1d\xea\x61\x61\xed\xe6\x87\x35\xd4\x2d\x16\xea\x4c\x0f\xb6\xa8\x69\x3e\x29\x3d\xd8\x7b\x2c\x49\x98\x6d\x74\x67\x83\x1c\x34\x54\xbf\xbb\x35\x34\xbc\x7d\x17\x0a\x0d\x85\xe2\xd5\x03\xb4\xaf\x56\x84\xf6\xea\x62\x8e\xd0\x6d\x7f\x43\x43\xd1\xb9\x42\x1a\x90\x34\xb2\xf6\x4c\xd0\x16\x0e\x17\x17\x05\xaf\xcb\xc8\x78\xa7\x7e\x1b\x26\x5b\x73\x13\xe5\xae\x54\xba\x03\xd6\x73\x75\xf2\xe6\x5d\x3d\x3d\x7b\x7a\xec\xb5\xdf\x2f\x9e\x29\xe6\x2e\x58\xa1\xdb\x07\xd0\xbf\xb3\x2b\x82\xb6\x53\x55\x77\x7a\xbc\xec\xfd\xc0\x08\x61\x01\x60\x9f\x6b\x0c\xd8\x17\x07\x83\xc0\x5d\xe9\x78\x4e\xfd\x53\xa7\x80\x76\xfa\xd9\xb8\xee\xf7\x07\xbc\x8e\x11\xe6\xec\x5d\x1e\x25\x09\x04\xd1\xad\xeb\x27\xf8\x8e\x59\xae\xb4\xc3\x1b\xa6\xf2\xf3\x2c\xed\x20\x30\x66\x6c\xa1\xd4\x70\x26\xf2\xf8\x49\xde\x3d\x13\x18\x31\x47\x46\x20\x9f\xe5\xd9\x97\xb5\xc3\xb4\x2f\xed\xa7\xf9\x76\x26\x92\x86\xbb\x65\x39\x76\xdc\xe7\x2d\x61\xb9\x7b\x19\xff\x9a\xfb\xba\xa3\x9c\x97\xfd\xc0\x89\x7d\xcc\xfd\x2f\x00\x00\xff\xff\x37\xa0\x5d\xd8\x99\x17\x00\x00")

func webStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_webStyleCss,
		"web/style.css",
	)
}

func webStyleCss() (*asset, error) {
	bytes, err := webStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "web/style.css", size: 6041, mode: os.FileMode(420), modTime: time.Unix(1541692968, 0)}
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
	"web/favicon.ico": webFaviconIco,
	"web/index.html":  webIndexHtml,
	"web/script.js":   webScriptJs,
	"web/style.css":   webStyleCss,
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
	"web": &bintree{nil, map[string]*bintree{
		"favicon.ico": &bintree{webFaviconIco, map[string]*bintree{}},
		"index.html":  &bintree{webIndexHtml, map[string]*bintree{}},
		"script.js":   &bintree{webScriptJs, map[string]*bintree{}},
		"style.css":   &bintree{webStyleCss, map[string]*bintree{}},
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
