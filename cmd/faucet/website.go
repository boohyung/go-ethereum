// Code generated by go-bindata.
// sources:
// faucet.html
// DO NOT EDIT!

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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\xef\x72\xdb\x36\x12\xff\xac\x3c\xc5\x86\x77\xad\xa5\xb1\x49\xda\x71\x26\xed\xc8\xa4\x3a\x99\x34\x97\xf6\xe6\xa6\xed\xb4\xe9\xdc\x75\xda\xce\x0d\x48\x2e\x49\xc4\x20\xc0\x02\x4b\xc9\xaa\x47\xef\x7e\x03\x80\xa4\x28\x59\x4e\xd3\x4b\xbf\xc8\x04\xb0\xf8\xed\x62\x77\xb1\x7f\xe0\xe4\xe9\x97\xdf\xbe\x7a\xfb\xd3\x77\xaf\xa1\xa6\x46\xac\x9e\x24\xf6\x0f\x08\x26\xab\x34\x40\x19\xac\x9e\xcc\x92\x1a\x59\xb1\x7a\x32\x9b\x25\x0d\x12\x83\xbc\x66\xda\x20\xa5\x41\x47\x65\xf8\x79\xb0\x5f\xa8\x89\xda\x10\x7f\xeb\xf8\x3a\x0d\xfe\x13\xfe\xf8\x32\x7c\xa5\x9a\x96\x11\xcf\x04\x06\x90\x2b\x49\x28\x29\x0d\xbe\x7e\x9d\x62\x51\xe1\x64\x9f\x64\x0d\xa6\xc1\x9a\xe3\xa6\x55\x9a\x26\xa4\x1b\x5e\x50\x9d\x16\xb8\xe6\x39\x86\x6e\x70\x01\x5c\x72\xe2\x4c\x84\x26\x67\x02\xd3\xab\x60\xf5\xc4\xe2\x10\x27\x81\xab\xfb\xfb\xe8\x1b\xa4\x8d\xd2\xb7\xbb\xdd\x12\xde\x70\xfa\xaa\xcb\xe0\x1f\xac\xcb\x91\x92\xd8\x93\x38\x6a\xc1\xe5\x2d\xd4\x1a\xcb\x34\xb0\x32\x9b\x65\x1c\xe7\x85\x7c\x67\xa2\x5c\xa8\xae\x28\x05\xd3\x18\xe5\xaa\x89\xd9\x3b\x76\x17\x0b\x9e\x99\x98\x36\x9c\x08\x75\x98\x29\x45\x86\x34\x6b\xe3\xeb\xe8\x3a\xfa\x2c\xce\x8d\x89\xc7\xb9\xa8\xe1\x32\xca\x8d\x09\x40\xa3\x48\x03\x43\x5b\x81\xa6\x46\xa4\x00\xe2\xd5\xff\xc7\xb7\x54\x92\x42\xb6\x41\xa3\x1a\x8c\x9f\x47\x9f\x45\x97\x8e\xe5\x74\xfa\xfd\x5c\x2d\x5b\x93\x6b\xde\x12\x18\x9d\x7f\x30\xdf\x77\xbf\x75\xa8\xb7\xf1\x75\x74\x15\x5d\xf5\x03\xc7\xe7\x9d\x09\x56\x49\xec\x01\x57\x1f\x85\x1d\x4a\x45\xdb\xf8\x59\xf4\x3c\xba\x8a\x5b\x96\xdf\xb2\x0a\x8b\x81\x93\x5d\x8a\x86\xc9\xbf\x8c\xef\x63\x36\x7c\x77\x6c\xc2\xbf\x82\x59\xa3\x1a\x94\x14\xbd\x33\xf1\xb3\xe8\xea\xf3\xe8\x72\x98\x78\x88\xef\x18\x58\xa3\x59\x56\xb3\x68\x8d\x9a\x78\xce\x44\x98\xa3\x24\xd4\x70\x6f\x67\x67\x0d\x97\x61\x8d\xbc\xaa\x69\x09\x57\x97\x97\x9f\xdc\x9c\x9a\x5d\xd7\x7e\xba\xe0\xa6\x15\x6c\xbb\x84\x52\xe0\x9d\x9f\x62\x82\x57\x32\xe4\x84\x8d\x59\x82\x47\x76\x0b\x3b\xc7\xb3\xd5\xaa\xd2\x68\x4c\xcf\xac\x55\x86\x13\x57\x72\x69\x3d\x8a\x11\x5f\xe3\x29\x5a\xd3\x32\xf9\x60\x03\xcb\x8c\x12\x1d\xe1\x91\x20\x99\x50\xf9\xad\x9f\x73\xd7\x78\x7a\x88\x5c\x09\xa5\x97\xb0\xa9\x79\xbf\x0d\x1c\x23\x68\x35\xf6\xf0\xd0\xb2\xa2\xe0\xb2\x5a\xc2\x8b\xb6\x3f\x0f\x34\x4c\x57\x5c\x2e\xe1\x72\xbf\x25\x89\x07\x35\x26\xb1\x8f\x58\x4f\x66\x49\xa6\x8a\xad\xb3\x61\xc1\xd7\x90\x0b\x66\x4c\x1a\x1c\xa9\xd8\x45\xa2\x03\x02\x1b\x80\x18\x97\xc3\xd2\xc1\x9a\x56\x9b\x00\x1c\xa3\x34\xf0\x42\x84\x99\x22\x52\xcd\x12\xae\xac\x78\xfd\x96\x23\x3c\x11\x8a\x2a\xbc\x7a\x36\x2c\xce\x92\xfa\x6a\x00\x21\xbc\xa3\xd0\xd9\x67\xb4\x4c\xb0\x4a\xf8\xb0\xb7\x64\x50\xb2\x30\x63\x54\x07\xc0\x34\x67\x61\xcd\x8b\x02\x65\x1a\x90\xee\xd0\xfa\x11\x5f\xc1\x34\xee\x0d\x61\xef\x65\x47\x35\x4a\x7b\x4e\xc2\xa2\x0f\x82\x70\x0c\x5b\x71\xaa\xbb\x2c\x64\x82\x1e\x05\x4f\xe2\xfa\x6a\x38\x52\x5c\xf0\x75\xaf\x91\xc9\xe7\x91\x72\x1e\x3f\xff\xe7\xd0\x7f\xa8\xb2\x34\x48\xe1\x44\x1d\x13\x62\x2e\xdb\x8e\xc2\x4a\xab\xae\x1d\xd7\x67\x89\x9b\x05\x5e\xa4\x41\xc5\x0d\x05\x40\xdb\xb6\xd7\x5d\x30\x1e\x49\xe9\x26\xb4\xa6\xd3\x4a\x04\xd0\x0a\x96\x63\xad\x44\x81\x3a\x0d\x7a\x9d\xbc\xe1\x86\xe0\xc7\xef\xff\x05\xbd\x81\xb9\xac\x60\xab\x3a\x0d\xaf\xa9\x46\x8d\x5d\x03\xac\x28\xac\x73\x47\x51\x34\xe1\xed\x3c\xfd\xa1\x74\x61\x46\x72\x4f\x35\x4b\xb2\x8e\x48\x8d\x84\x19\x49\xc8\x48\x86\x05\x96\xac\x13\xa3\xc4\x9e\x28\x00\x25\x73\xc1\xf3\xdb\x34\xb8\xbf\xe7\x25\x44\xdf\x63\xce\x5a\xca\x6b\xb6\xdb\x55\x7a\xf8\x8e\xf0\x0e\xf3\x8e\x70\xbe\xb8\xbf\x47\x61\x70\xb7\x33\x5d\xd6\x70\x72\x63\x59\xec\x76\xc1\xea\x0d\x5f\x23\x34\xe8\x0f\xf0\x34\x89\x3d\xfc\x5e\xf4\xd8\xca\x3e\x6a\xd9\x19\xed\x01\xc3\x13\x36\xa8\xc2\x51\x88\x00\x0a\x46\x2c\x34\x9c\xf0\x16\xb7\x56\xde\xe9\xde\x7e\x35\x67\x42\x64\xcc\x1e\xc7\x4b\x38\x6e\xfa\x1d\xad\xca\xd6\xdc\xb8\x22\x60\x35\x48\xe0\xa4\xff\x33\x4e\x75\x74\xe3\x48\xb5\x4b\xb8\x7e\x36\xb9\x6e\xa7\xfc\xed\xc5\x91\xbf\x5d\x9f\x24\x6e\x99\x44\x01\xee\x37\x34\x0d\x13\xc3\xf7\x60\xb8\xbd\x32\x8f\x37\x85\x36\xb8\x8c\xa2\x8d\x41\xea\xf2\x06\xd4\x1a\x75\x29\xd4\x66\x09\xac\x23\x75\x03\x0d\xbb\x1b\x03\xf5\xf5\xe5\xe5\x54\x6e\x5b\xbc\xb0\x4c\xa0\xf3\x6d\x8d\xbf\x75\x68\xc8\x8c\x3e\xed\x97\xdc\xaf\x75\xed\x02\xa5\xc1\xe2\x48\x1b\x96\xa3\x55\xad\xa3\x9a\x98\x7e\x54\xe6\x49\xd9\x4b\xa5\xc6\xd8\x37\x15\xa3\x87\x9e\x84\xe9\x60\x95\x90\xde\xd3\xcd\x12\x2a\xfe\x54\xec\xd2\xb6\x36\x79\x2c\x74\xf9\xcb\x65\xcf\xde\x22\x6a\x9f\x18\xad\xcb\x82\x1b\x26\x31\x15\x1f\xc1\xd9\x3a\x61\xc6\x0c\x7e\x08\x7b\x97\xa2\xf6\xec\xdd\xf0\x63\xf9\xd7\xc8\x34\x65\xc8\x1e\x8f\xae\x13\x01\xca\x4e\x16\x93\xf3\xbb\x1b\xfd\xb1\x02\x74\x92\xaf\x51\x1b\x4e\xdb\x0f\x95\x00\x8b\xbd\x08\x7e\x7c\x28\x42\x12\x93\x7e\xbf\xaf\x4d\x07\x7f\xd1\xe5\xfe\xa3\x5c\x7a\xbd\xfa\x4a\x6d\xa0\x50\x68\x80\x6a\x6e\xc0\x66\xc2\x2f\x92\xb8\xbe\x1e\x49\xda\xd5\x5b\xbb\xe0\x94\x0a\xa5\x4f\x86\xdc\x80\xee\xa4\x4b\x02\x4a\x02\xd5\x78\x98\x47\xa5\xff\x8a\xe0\xad\xb2\xb5\xc8\x1a\x25\x41\xc3\x04\xcf\xb9\xea\x0c\xb0\x9c\x94\x36\x50\x6a\xd5\x00\xde\xd5\xac\x33\x64\x81\x6c\xf8\x60\x6b\xc6\x85\xbb\x4b\xce\xa4\xa0\x34\xb0\x3c\xef\x9a\xce\xd6\x52\xb2\x02\x94\xaa\xab\xea\x5e\x16\x52\xd0\xa8\x4e\x12\x08\x25\xab\x51\x1e\xd3\xb2\x06\x18\x11\xcb\x6f\xcd\x05\x0c\x51\x01\x98\x46\x20\x8e\x85\xdd\xd5\xa7\x34\x96\xe7\x76\xbb\x89\xe0\xa5\xdc\x2a\x89\x50\xb3\xb5\x13\xe4\x88\x00\x1a\xb6\x1d\x80\x7a\xb9\x36\x9c\x6a\xee\x0f\xde\xa2\x6e\x6c\x71\x5c\x80\xe0\x0d\x27\x50\x25\x24\x86\xb4\x92\x95\xed\xa9\x5e\x3a\x09\x77\x3b\x2f\xf2\xdc\x2c\x20\xb6\xaa\xfa\x0e\x35\x57\xc5\x6e\x67\xeb\x2e\x47\x1a\x3d\x48\x2d\xf0\xb6\xc6\x13\xea\x1e\x33\x02\x68\x7c\xe5\x69\xa1\xd5\x8a\x30\xb7\x55\x24\xb0\x8a\x71\x69\x08\x32\x45\x26\xea\x93\x45\x12\xb7\x53\x63\xaa\xc3\xb3\x5c\x80\xe1\x4d\x2b\xb6\x90\x6b\x64\x84\xc0\x20\x61\x47\x8d\x96\x2d\x1b\x22\x5f\xef\xb8\x52\x3d\x00\x62\xba\xb2\x6d\xec\x7f\x59\xa6\x3a\x5a\x66\x82\xc9\x5b\x9b\x51\xc7\x52\x21\x89\xd9\xca\x69\xe9\x74\x91\x00\x2d\x33\x56\x65\x5c\x92\x72\x5a\xec\xfb\x56\x03\x73\x3b\x2a\xb9\x40\xd7\xda\x3a\xc7\x94\x67\xd6\x04\xb6\xff\x58\x5c\x40\xae\xda\xad\xdf\xed\xf6\x59\xd1\x8c\xab\x4b\x46\x28\x96\xa9\x35\x82\x2f\x7a\x32\x75\x07\x4c\x16\x50\x72\x8d\xc0\x36\x6c\xfb\x14\x7e\x52\x1d\xe4\x4c\x02\x69\x96\xdf\x7a\xde\x9d\xd6\xd6\x43\x5b\x94\x36\x0b\xed\x7d\x26\x43\xa1\x36\x8e\xc4\xa3\x95\x1c\x85\x73\x20\x83\x08\xb5\xda\x40\xd3\xe5\xee\x80\xd6\x73\xd0\x2e\x6c\x18\x27\xe8\x24\x71\xe1\xcf\x4d\x9d\x96\x90\xab\x06\x4d\xb4\xb7\xc2\xc9\x9b\x3d\x7e\xf5\x1f\xfb\xde\xc9\x2d\xc7\x31\xbc\x11\x2a\x63\x02\xd6\x36\x18\x65\xc2\xde\x57\x05\xb6\x48\x3b\x38\x83\x21\x46\x9d\xb1\x4e\x48\xa3\xfb\xd8\xfd\x6b\xa6\xed\xa5\xc0\xa6\x25\x48\xfb\xca\xdf\xce\x19\xd4\x6b\xdb\xcf\xf4\x3c\xbe\xc4\x92\x4b\xaf\xd9\xb2\x93\xde\xa5\xa8\x66\x04\xbe\x36\x31\xc0\x9c\xc6\xa1\xd3\x02\x7a\x75\x7b\x84\x11\xcf\xd1\x41\x3a\x6e\x9f\x3f\x70\xec\xfe\xa3\x77\xce\x45\xdf\xa8\x78\x98\xc8\xa0\x2c\xe6\xff\xfc\xe1\xdb\x6f\x22\x43\x9a\xcb\x8a\x97\xdb\xf9\x7d\xa7\xc5\x12\xfe\x3e\x0f\xfe\xe6\xea\xd7\xc5\xcf\x97\xbf\x46\x6b\x26\x3a\x7c\x00\x7d\x01\xfd\xe7\x12\x0e\xb9\xec\x16\x8b\x9b\xd3\xc5\xdb\xa4\x64\xd4\x68\x90\xe6\x96\x70\xac\xb1\x76\x37\x87\x8a\x61\xd0\x20\xd5\xca\x39\x81\xc6\x5c\x49\x89\x39\x41\xd7\x2a\xd9\xeb\x01\x84\x32\x66\x50\xc6\x9e\x62\xa2\x8f\xe1\xc0\xbc\x84\xf9\x60\x91\x4f\xe0\x19\xa4\x29\x5c\x0e\x6b\xbd\x36\x20\x05\x89\x1b\xf8\x37\x66\x3f\xa8\xfc\x16\x69\x1e\x6c\x8c\xbd\x8f\x01\x9c\x83\x50\x39\xb3\x78\x51\xad\x0c\xc1\x39\x04\x31\x6b\x79\xb0\xf0\x2d\xde\x0e\x6c\xcd\xfb\xc7\x60\x1f\x84\xe5\x9b\x60\x2f\xe9\xf9\xb9\x77\x95\xc1\x5c\x4a\x36\x68\x0c\xab\x70\x7a\x42\x17\xef\xc7\xa3\x58\x45\x34\xa6\x82\x14\x9c\x59\x5b\xa6\x0d\x7a\x92\xc8\xd6\x18\x3d\x17\xa7\x0e\x47\x96\xa6\x20\x3b\x21\xc6\xfd\x33\x8d\xf6\x16\xf5\x64\xbb\x27\x07\xe4\x91\x0f\xc7\x4f\xd3\x14\x6c\xc2\xb5\x36\x2a\xf6\x3b\xad\xcb\xf8\xd2\x60\x11\xd9\x9c\xbf\xdf\xb1\x18\xe1\x1e\xa0\x61\xf1\x47\x70\x58\x1c\xe3\x61\xf1\x08\xa0\xab\xc4\xde\x87\xe7\x2b\xb7\x09\x9c\x9b\x78\x04\x4d\x76\x4d\x86\xfa\x7d\x70\xbe\x12\xeb\xe1\x9c\xaa\xbf\x96\x34\xd9\x7b\x01\x57\x2f\x16\x8f\xa0\xa3\xd6\xea\x51\x70\xa9\x68\x3b\xbf\x17\x6c\x6b\xc3\x3d\x9c\x91\x6a\x5f\xb9\xc2\xe9\xec\x02\x2c\xaf\x25\x8c\x08\x17\xae\x5b\x5b\xc2\x99\x1b\x9d\xed\x1e\xe1\x66\xba\x3c\xb7\x89\xe0\x63\xf8\xf5\x18\x23\xc7\x7e\xfc\x28\xcf\x31\xb0\x1f\x30\x85\x4f\x3f\x85\x07\xab\x87\x2e\x68\x7d\xb8\xcf\x50\x90\x42\x10\xf4\xf0\xb3\x52\x69\x98\xdb\x45\x9e\x5e\xde\x00\x4f\xa6\x30\x91\x40\x59\x51\x7d\x03\xfc\xfc\x7c\x8f\x34\x1b\x60\xce\x53\x08\x6c\x6f\x90\x50\xb1\x72\x35\x9a\x2f\xe4\x7e\x09\x6c\x2f\x68\x7b\x64\x59\x2c\x6d\x98\x9d\x9f\xed\xb3\xf0\x24\x01\x9f\x1f\x88\xfc\x33\xff\x35\xea\x0c\x6a\x97\x32\xcf\x21\x88\x5a\x59\x7d\xe1\x3a\xc8\x17\xcf\xcf\x16\x37\xb0\xc7\x74\x7d\xe5\x12\x72\xdb\x65\xdd\x80\xef\x54\x5c\xbd\x08\x63\x8f\xe5\x46\x99\xd2\x05\xea\x50\xb3\x82\x77\x66\x09\xcf\xdb\xbb\x9b\x5f\x86\x1e\xd4\x55\xb5\x4e\xee\x56\xe3\xea\x94\x2c\x43\xe1\x74\x0e\x41\x12\x5b\xa2\x61\xcb\x78\xca\xe9\x53\x16\x9c\xa8\xc7\x61\x7c\x68\xea\xe7\x1b\x5e\x14\x02\xad\x10\x8e\xa1\x7f\x11\x2c\x3a\xed\x02\xd7\xdc\x8f\xe7\xc7\x72\x10\x6f\x70\x11\x75\x92\xdf\xcd\x17\x61\x4f\x33\x8c\x2f\xe0\xcc\xd8\xf8\x5c\x98\xb3\x45\x54\x77\x0d\x93\xfc\x77\x9c\xdb\xe2\x7e\xe1\xe5\xb6\x12\xdb\x8a\x7d\xb4\xf6\x6e\x72\xd1\xc6\x6e\x73\x11\xd5\xd4\x88\x79\x90\x90\x7b\x2e\xb3\xc2\x8d\x26\x76\x28\x7e\xfa\xd0\x23\x77\x87\x31\x34\x17\xca\xe0\x51\x8e\x00\x83\xf4\x96\x37\xa8\x3a\x9a\x8f\x79\xe4\xc2\x76\xc0\x97\x8b\x1b\xd8\xed\x5f\x15\xe3\x18\x5e\x1b\xdb\x53\x70\x53\x03\x83\x0d\x66\xc6\xc5\x77\xe8\xf7\xb8\x14\xee\x53\xf5\xcb\xef\xbe\x9e\xa4\xeb\x11\x75\xee\x84\x1b\x5f\x55\x4f\xe5\xc9\x93\xcf\xb8\x9b\xcd\x26\xaa\x94\xaa\x84\x7f\xc0\x1d\x13\xa9\xcd\x1e\xd1\x3b\xdb\xb8\x9a\xad\xcc\xa1\xc0\x12\xf5\x6a\x02\xdf\x67\xd7\x24\xf6\x0f\x8c\x49\xec\xff\x79\xf2\xbf\x00\x00\x00\xff\xff\x82\x9c\x59\xe7\x4d\x19\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"faucet.html": faucetHtml,
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
	"faucet.html": &bintree{faucetHtml, map[string]*bintree{}},
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
