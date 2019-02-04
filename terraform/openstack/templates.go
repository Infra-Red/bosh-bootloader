// Code generated by go-bindata.
// sources:
// templates/provider-vars.tf
// templates/provider.tf
// templates/resources-outputs.tf
// templates/resources-vars.tf
// templates/resources.tf
// DO NOT EDIT!

package openstack

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

var _templatesProviderVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd1\xc1\x6a\xe3\x30\x10\x06\xe0\xbb\x9f\x62\xd6\x7b\xd9\x25\x0b\x79\x82\x3d\x84\x5c\x03\x85\x9a\x9e\xc3\x44\x1e\xa3\x69\xe4\x19\x31\x1a\x25\x94\xd2\x77\x2f\x76\x4a\x49\x82\xa1\xb9\x8a\xff\xff\xf4\x0b\xfd\x06\x0c\x81\x4a\x81\xa0\x6a\x3d\x0b\x3a\x95\x75\x30\xea\x49\x9c\x31\x95\xe6\x84\xc6\x78\x48\x04\x2d\x56\x8f\xfb\x6a\xa9\x85\xf7\x06\xa0\xa7\x12\x8c\xb3\xb3\x0a\xfc\x87\x76\x53\x3d\x4e\x95\x80\xf3\x09\x49\x9f\x95\xc5\xe1\xe5\x79\x07\x83\x1a\x3c\x65\x92\xce\x31\x1c\x21\x9b\x9e\xb8\x27\x83\x3f\x2a\xe9\x0d\x4a\x88\x34\xd2\x2a\x6a\xf1\x55\x56\xf3\x7f\x70\xa8\x0e\x67\xf6\xa8\xd5\x21\xa3\xc7\x5f\x7f\xdb\xe6\xa3\xb9\x1a\x52\x0b\xd9\x5e\x70\xa4\xc5\x25\x57\x37\x71\xa6\xc4\x42\xe0\x14\xa2\x70\xc0\x04\x53\x15\xe6\xea\x2d\x99\xb1\x94\xb3\x5a\xff\x83\x38\xd7\xbf\xb3\xb7\x44\xaf\x23\xb2\x3c\xb2\xeb\x92\x5c\x9a\xe1\x24\x28\xfe\xd0\xdb\x4c\x5f\x29\xf8\xfa\xd2\x58\xb2\x58\x0a\x85\x6a\xcb\x50\x77\xe4\x0c\x5d\xb7\x83\x13\x19\x0f\x5f\xbf\xd6\xce\xb9\x01\x6b\xf2\x29\x33\x60\x2a\xf7\x68\xc0\x40\xe6\xfb\x81\xd3\xb2\xbb\xad\xc5\x75\x84\xed\x06\xa6\xdc\x45\xa6\x3b\x77\x22\x3f\x03\x00\x00\xff\xff\x47\xcf\xad\x74\x77\x02\x00\x00")

func templatesProviderVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesProviderVarsTf,
		"templates/provider-vars.tf",
	)
}

func templatesProviderVarsTf() (*asset, error) {
	bytes, err := templatesProviderVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/provider-vars.tf", size: 631, mode: os.FileMode(480), modTime: time.Unix(1549295905, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesProviderTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcd\x41\x0e\x02\x21\x0c\x85\xe1\x3d\xa7\x68\x88\x6b\x6f\xe0\x59\x26\x0d\xd4\x48\x9c\x29\xa4\xc0\xb8\x98\x70\x77\x63\x23\x23\xd8\xe5\x9f\x2f\xaf\x49\xe2\x1e\x3c\x09\xd8\x98\x88\x73\x41\xf7\xb4\x70\x18\x00\xac\xe5\xb1\x54\x59\xe1\x73\x37\xb0\x97\x63\x47\xb9\xf6\xda\xac\x01\xa8\x99\x64\x61\xdc\x68\x22\x67\x55\x93\x30\xe7\x57\x14\x3f\xcf\xf4\xaa\xa4\x10\x23\x97\xef\xd0\x49\x86\xaa\xca\xc7\x0d\x03\xff\xab\xa1\xaa\x0a\x9c\xc9\x55\xa1\xf9\x5d\xaf\x4a\x1c\x3a\x92\xb2\xdc\xc3\x4a\x3f\x31\xc4\x66\x4d\x33\xef\x00\x00\x00\xff\xff\x54\x73\x4c\x37\x17\x01\x00\x00")

func templatesProviderTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesProviderTf,
		"templates/provider.tf",
	)
}

func templatesProviderTf() (*asset, error) {
	bytes, err := templatesProviderTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/provider.tf", size: 279, mode: os.FileMode(480), modTime: time.Unix(1549295712, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResourcesOutputsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\xcf\x8a\xdb\x30\x10\xc6\xef\x7e\x0a\x61\x7a\x68\xa1\x98\x24\xd0\x1e\x0a\x7d\x92\x52\xc4\xd8\x9e\x75\x26\xb1\x25\x31\x1a\x29\x9b\x5d\xf2\xee\x45\xb5\xbd\x89\x77\x95\x36\x6d\xf6\x16\x33\xf3\xfd\xbe\xf9\x67\xc7\x06\x71\x41\x54\xd9\xe2\x03\x84\x5e\xf4\x1e\x8f\xda\xc0\x80\xa5\x7a\x2e\x94\x8a\xd0\x07\x54\xdf\x55\xf9\xe1\xd9\x3a\x34\x5e\xa0\xd9\xeb\xc6\x0e\x2e\x08\xa6\x54\x07\xc4\x3a\x6e\xaa\xda\xfa\x6d\x95\x64\xa7\xb2\x38\x15\xc5\x0c\x75\x4c\x11\xc6\xcc\xff\xe0\x5d\xa8\x4f\x65\xa1\x94\x47\xe3\x49\x28\x26\x80\x70\xc0\x4b\x27\x7c\x14\x64\x03\xbd\x26\xf7\x07\x27\x83\x72\xb0\xbc\x27\xd3\xe9\x87\xde\x82\x90\xe9\xc8\x25\xbf\x5d\x5d\x41\xdb\x32\x7a\xbf\x6c\x20\x0e\x5e\x7b\x6c\x02\x93\x1c\x75\xc7\x36\x38\xbf\xc4\xff\xb8\xc6\xf7\xd8\xfc\xce\x4f\xf4\x38\xf8\x69\x38\x3f\x2f\xe1\xbb\x30\xb8\xda\x3e\x6a\x3d\xcf\xfe\x7e\xa7\x5d\x9d\x33\x6a\x89\xb1\x11\xcb\xef\xe9\x74\xb1\xf1\x85\x17\x99\x69\x11\x0d\xb5\x7c\xdb\x2a\x7c\xa8\x0d\xca\x0c\x9d\x1e\xab\xa4\x5f\x2e\xe3\x05\xdd\x1d\xee\x01\x77\x20\x78\x80\xa3\x26\xb7\xc4\xa7\x54\x6a\x6f\x23\x4f\x3f\x5f\x06\x41\xed\x95\x52\x33\xe7\x98\x1a\xdb\x5a\x2f\x1f\xff\x71\x18\x9f\xd5\x7a\xf5\x69\x69\xc3\x36\x08\xf2\xcd\x55\x4f\xe9\x33\x7b\x7c\x7c\x53\xfb\xf9\x5a\xae\x77\xb1\x5e\x55\xab\x6a\x5d\x7d\x2d\xb3\xf7\xfc\x57\xdd\x97\xac\x2e\x70\x7f\xe7\xbb\xfb\x6d\xb3\x59\x90\x21\xc8\x36\x8b\x8d\xc0\xd5\x1c\x5c\x76\x0f\x4f\xf9\xe4\x08\xd4\x43\x4d\x7d\x7a\x69\x9e\xac\x79\xf5\xa1\x3b\x97\xe9\xd8\xee\xb0\x91\x2c\x44\xd0\x80\x11\xfd\xf6\x3b\x79\x96\xb7\x76\x00\x32\x59\xf5\x18\xca\xa8\x19\x3b\xb2\x79\xcd\x18\xca\x68\xd0\xc4\xdc\xd9\x24\xcd\x18\xba\x72\x13\xd9\xbf\x86\x57\xaa\x5f\x01\x00\x00\xff\xff\xdc\x6a\x0a\x6b\x53\x06\x00\x00")

func templatesResourcesOutputsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResourcesOutputsTf,
		"templates/resources-outputs.tf",
	)
}

func templatesResourcesOutputsTf() (*asset, error) {
	bytes, err := templatesResourcesOutputsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resources-outputs.tf", size: 1619, mode: os.FileMode(480), modTime: time.Unix(1549297079, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResourcesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xc1\x4a\x03\x31\x10\x86\xef\x79\x8a\x9f\x3c\x80\x67\x2f\xbd\x79\x29\x88\x16\x7a\x14\x59\xa6\xbb\xb3\x65\x68\x9c\x2c\x93\xe9\xda\x2a\xbe\xbb\x64\x17\xa5\x2e\x82\x96\xdc\x26\xdf\xfc\x7f\xf2\x8d\x64\x42\xbb\xc4\x88\xac\x63\x23\x5d\xc4\xfb\x47\x08\x17\xd3\x93\x37\xca\x3e\xdf\x04\xa0\xe3\xd2\x9a\x0c\x2e\x59\xb1\x42\x7c\x1c\x58\xb7\x4e\xed\x01\x7c\x72\x36\xa5\x04\x65\x7f\xcd\x76\x80\x74\xf0\x8c\xd6\x98\x9c\x61\xf9\xe8\x6c\x10\x75\xb6\x9e\x5a\xc6\x90\xcd\x63\xf8\x51\x45\x23\x49\xa2\x9d\x24\xf1\x73\xf3\x96\x95\xff\x68\xbc\xe4\x51\x79\x28\xbd\xf0\x22\xb4\xd3\xd2\xd4\x71\x61\x1b\xd9\xca\xaf\x91\xf7\x52\x1c\xb9\xc7\xdd\xc3\x16\x33\x87\xf5\xa6\xc4\x09\xec\xe9\x98\x1c\x2b\x3c\xc5\xdb\x9b\xe9\xc4\xe7\x00\xf8\x79\xe0\xba\x99\xa4\x2c\x7f\xf1\x25\x6c\x7a\xcb\x95\xca\xea\x4e\x95\x66\xbc\x97\x52\x7d\xf5\x29\x93\x8b\xee\xb1\xde\x2c\x6a\x2a\x92\xf5\x3f\x2d\x33\xf9\xed\xe6\x33\x00\x00\xff\xff\xe0\x14\x40\x4e\xf1\x01\x00\x00")

func templatesResourcesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResourcesVarsTf,
		"templates/resources-vars.tf",
	)
}

func templatesResourcesVarsTf() (*asset, error) {
	bytes, err := templatesResourcesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resources-vars.tf", size: 497, mode: os.FileMode(480), modTime: time.Unix(1549295885, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResourcesTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xc1\x6e\xdb\x38\x10\xbd\xfb\x2b\x06\xca\x1e\x76\x0f\x51\x1d\xd5\x4e\x7c\xc9\xc1\x69\x81\xb6\x0b\xb4\x5b\x6c\x9a\xc5\xde\x04\x5a\x9a\x28\x4c\x24\x92\x18\x52\x6e\x84\x22\xff\xbe\x20\x25\xd9\x72\x2c\xdb\xb2\xb3\x1b\xdb\x8b\xe6\xe4\x0c\x47\xe4\x9b\xf7\xde\x50\xa4\x4e\xe0\x01\x0b\x50\x8c\x93\xee\x11\x6a\x99\x53\x84\xe0\x49\x85\x42\x1b\x16\x3d\x84\x91\xcc\x54\x6e\x30\x7c\xc0\xc2\x26\x85\xd3\xc0\x03\x6f\x22\xf5\x9d\x07\x3f\x7a\x00\x84\x09\x97\x02\xec\xdf\x25\x78\xbf\xfc\x98\x32\xf2\xcb\x58\x28\x58\x86\x4f\x5e\x0f\xc0\xfe\x00\xa8\x73\xec\xb3\xa7\x65\xa2\x41\xc1\x84\xa9\x13\x9f\x7a\xbd\x13\x10\x68\xbe\x4b\x7a\x68\xc7\x52\x0d\x72\x91\xd4\x3f\xd7\xc0\xe9\x0a\xa9\x01\xcb\x8e\xb0\x38\xe3\x22\xd4\x86\x19\x0c\x73\x65\x47\x0c\xe5\xe8\xd0\x6d\x80\xa4\xf3\x89\x40\x33\x47\x54\x05\x5a\x81\xad\x83\x56\x55\xc6\xe3\x66\xe6\x06\x12\x7c\xbb\xa0\xcf\x63\x37\x43\xc4\x63\x02\x58\x5c\xeb\xac\xef\xf7\xfd\x33\xbf\xff\x26\x18\xd8\x14\xae\xc2\x29\x92\x9e\x43\xba\x84\xc1\x32\x2d\x33\x62\x6c\x29\x8e\x9c\x34\x95\x11\x33\x16\xaf\x92\x32\xd5\x70\xe9\x8a\x03\xd0\x86\x91\x69\xac\x13\xf4\xfb\x9e\x1b\x40\x11\x2f\x00\x08\x86\x6e\xfd\xa7\x1e\x40\xc2\x0c\x7e\x67\x45\xc8\xd5\x12\xcc\x33\x9b\x83\x82\x4d\x52\x0c\xe3\xbb\x48\xcd\xc6\x4b\x2d\x00\x62\xa1\x1d\x63\x1a\xc9\xd6\xd1\xa0\xf3\xd9\x48\xed\x2b\x92\xb9\x41\xda\x24\x61\x99\xd5\x90\xb0\x0c\x6c\x2d\xe1\x0a\x1a\x4f\xab\xe9\x96\x6d\xb6\x50\x1c\x3e\x1a\x24\xc1\xd2\xb0\x61\x85\xd9\x5a\xf8\x68\x6c\x3c\x74\x5a\x6f\x36\x65\x55\x11\x17\x06\xe9\x96\x45\xd8\xa8\x4d\x49\x5a\x32\xe7\xca\x92\xea\x79\xe2\x35\x76\x9c\xb1\xe7\x37\xb8\xab\x4d\x59\x75\xc7\xda\x09\x66\x1d\xe4\x37\xfa\xc7\xaf\x2b\x3d\x81\xdb\x54\x32\xc3\x45\x02\x5c\x6d\xdc\x20\xea\x5c\xeb\x74\x5b\xf2\xfd\x64\xa1\xd6\x95\x85\x5a\x5f\x37\x89\xa8\xf9\x6e\x6e\x52\x1a\xa3\x9c\xb8\x29\x20\x21\x99\x6f\x86\xa2\x31\x72\x89\xdb\x01\x71\x26\xba\x74\xf9\xd6\xf0\xa8\x23\xe2\xca\x54\x8f\xfc\x9e\x67\x6a\x22\x1f\xe1\xba\x46\xf2\xc1\x2e\x60\xe1\x75\xc5\x42\x79\x8a\x35\xa0\xf2\x9f\xb3\x12\xd7\xb3\x95\xc6\xa2\x90\x02\xc1\x48\x98\xad\x79\xfd\xd1\xeb\x80\x3f\xe6\x84\x51\x3d\x0b\x17\x09\xa1\xd6\xce\xdd\xe6\x0e\xc9\x14\xca\x15\xf7\xe9\xeb\xd4\xed\x07\x8a\xa4\x91\x91\x4c\x5d\x1b\x44\xaa\x14\x82\x4c\x48\x4c\x24\x18\x66\xdc\x4e\x12\x04\xcf\xa2\xec\xb1\x8e\x12\x66\xd2\x60\xc8\x55\xa8\x08\x6f\xb9\x8d\x7b\x76\x27\xe9\xfb\xfd\x37\x6e\x23\xaa\x15\x0b\xcb\xe2\xd7\xbb\x70\x2e\x97\x7f\x3f\xa9\xed\xb7\x3b\xaf\x41\x47\x5e\xc7\x09\x0a\x03\xbf\xde\x4a\x02\x6b\x7f\x88\x08\x99\xc1\x53\x14\x53\x38\x8d\xe1\xbe\xcc\xfa\x6d\x2f\xd4\x9f\x8f\xce\x47\x6d\xe4\x57\xf1\x57\xa0\x7f\xbb\x1e\x9b\x1f\x08\x9e\xb1\x7e\xf5\xc7\xf5\x47\x78\xef\xe8\x91\xb4\xd4\x3d\xdd\x9b\xb2\x3c\x2d\xec\xe2\x8a\x72\x6b\x5c\xdd\x6f\xb5\x1b\x8c\x9c\xe3\xfc\xf2\xe1\xd3\x97\xbf\xf7\xd3\x73\xc3\xe1\x70\xd8\xda\x76\xd5\x40\x25\xfd\x8e\xba\xee\xee\x8c\xf9\x79\xe7\x65\x22\xb4\x37\x67\x9b\x08\x9f\xaf\x6e\xae\x0f\xb8\xf9\x8e\x56\x81\xb7\x9d\x15\x38\xdc\x17\xcf\xd1\x92\x3f\xe8\x4c\xfe\xcd\x78\xbc\x17\xf2\x47\x83\xc1\xdb\x36\xfa\xab\xf8\xb1\x0b\x30\xec\x2c\xc0\x3b\xc2\xf8\xae\xbc\x86\xbd\xbe\x08\xa3\xc1\xa0\x55\x84\x32\x7e\xec\x22\x9c\xaf\x3e\x2b\xc4\xa8\x52\x59\x60\x0c\x7f\x7d\xd6\x8b\xef\xe4\xf1\xb7\xfd\xbc\x0e\x06\x41\xfb\x7e\x54\xc5\x77\xd3\x62\x9a\xe9\x43\x11\xe3\x62\x07\x31\xfe\xc4\x84\x6b\x43\xc5\x9e\xce\x48\x17\x17\x17\xed\x67\xa4\x72\xe0\xe8\x25\x19\xed\x20\xc9\x55\x2a\x27\xda\x48\xc2\x3d\x69\x12\x0c\xfb\xed\x9a\x94\x03\x07\xa1\xc9\x76\x97\x9a\x69\xa6\xb7\xfb\x72\x60\x1f\xe8\x22\xdb\xbf\xf1\x11\x61\x9a\xe9\x75\xb7\x9a\x56\xaf\x2c\x07\xbf\xbd\xfb\xfa\x1a\x76\xd9\x9b\xf8\xb3\x29\x5e\x44\x71\xfb\x9d\xa5\x23\xc5\x37\xef\xff\x5b\x8a\xf3\xf8\xff\x40\xf1\xc6\x4b\x49\x4b\x17\xfd\xbc\x9d\xb4\x68\xd0\xfb\x27\x00\x00\xff\xff\x24\xa4\x8f\x74\x60\x1a\x00\x00")

func templatesResourcesTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResourcesTf,
		"templates/resources.tf",
	)
}

func templatesResourcesTf() (*asset, error) {
	bytes, err := templatesResourcesTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resources.tf", size: 6752, mode: os.FileMode(480), modTime: time.Unix(1549296603, 0)}
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
	"templates/provider-vars.tf": templatesProviderVarsTf,
	"templates/provider.tf": templatesProviderTf,
	"templates/resources-outputs.tf": templatesResourcesOutputsTf,
	"templates/resources-vars.tf": templatesResourcesVarsTf,
	"templates/resources.tf": templatesResourcesTf,
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
		"provider-vars.tf": &bintree{templatesProviderVarsTf, map[string]*bintree{}},
		"provider.tf": &bintree{templatesProviderTf, map[string]*bintree{}},
		"resources-outputs.tf": &bintree{templatesResourcesOutputsTf, map[string]*bintree{}},
		"resources-vars.tf": &bintree{templatesResourcesVarsTf, map[string]*bintree{}},
		"resources.tf": &bintree{templatesResourcesTf, map[string]*bintree{}},
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
