// Code generated by go-bindata. (@generated) DO NOT EDIT.

 //Package v1 generated by go-bindata.// sources:
// pkg/api/v1/service.swagger.json
package v1

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

var _serviceSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x51\x73\xdb\xb6\x93\x7f\xcf\xa7\xd8\x51\x6f\xe6\xd2\x1b\x55\x6e\xf2\xff\xcf\xcd\x9d\x3b\x7d\x70\x13\xb7\xf1\x4d\xe2\x78\x1c\x27\x7d\xa8\x32\x36\x44\x2e\x25\x34\x14\xc0\x02\xa0\x1c\x35\xe3\xef\x7e\xb3\x00\x48\x82\x14\x65\x4b\xb2\xec\x48\x29\xf5\x90\xc8\x24\xb0\xd8\x5d\xec\xee\x6f\xb1\x04\xa1\x2f\x4f\x00\x7a\xfa\x9a\x8d\xc7\xa8\x7a\x87\xd0\x7b\x3e\xf8\xb1\xd7\xa7\x6b\x5c\x24\xb2\x77\x08\x74\x1f\xa0\x67\xb8\x49\x91\xee\x67\x9f\xc6\x07\x2c\xe3\x07\xb3\x67\x07\x1a\xd5\x8c\x47\x38\xc8\x94\x34\xd2\x76\x02\xe8\xcd\x50\x69\x2e\x05\x35\xf5\x5f\x41\x48\x03\x1a\x4d\xef\x09\xc0\x8d\x25\x1d\x49\xa1\xf3\x29\xea\xde\x21\xfc\xe1\x7a\xb1\x2c\x4b\x79\xc4\x0c\x97\xe2\xe0\x4f\x2d\x05\xb5\xfd\x68\xdb\x66\x4a\xc6\x79\xb4\x62\x5b\x66\x26\xba\xe2\x99\x78\x24\x56\xe5\xe8\x4f\x8c\xcc\xa5\x99\x67\x58\xdd\x05\xe8\x8d\xd1\x04\x7f\x02\xf4\x64\x86\xca\xd2\x3d\x89\x89\xff\xb7\xb6\xdf\x05\x75\xbb\x7c\xcd\xb5\xf1\x22\xda\xa6\x0a\x75\x26\x85\xae\x11\xb4\x37\x9e\xff\xf8\x63\xe3\x12\x40\x2f\x46\x1d\x29\x9e\x19\xaf\x98\x23\xd0\x79\x14\xa1\xd6\x49\x9e\x42\x41\x69\x10\x90\xb7\x9d\x74\x34\xc1\x29\x5b\x20\x06\xd0\xfb\x0f\x85\x09\xd1\xf9\xee\x20\xc6\x84\x0b\x4e\x74\xf5\xc1\xec\x19\x31\x19\x30\x7d\xee\x29\xf7\x6a\xfd\x6f\x82\xbf\x6e\xc2\x21\x7b\x31\x26\x2c\x4f\xcd\xdd\xec\x0b\xc8\x05\x7e\xce\x30\x32\x18\x03\x2a\x25\x55\x29\xc5\x7d\x85\x50\xb9\x30\x7c\x8a\xc7\x44\xf4\x16\xbe\x9f\xb4\x48\xd0\xcb\x98\x62\x53\x34\xa8\x2a\x63\x71\x9f\x86\x38\x82\x4d\xad\x29\xeb\x89\xbc\xbe\x8c\x31\x45\x83\x71\x93\x6f\x6e\x25\xfd\x2b\x47\x35\x6f\xde\x52\xf8\x57\xce\x15\x92\x89\x24\x2c\xd5\xd8\xb8\x4d\x66\x46\x7d\x47\x52\xa6\xc8\x44\xb3\x77\x22\xd5\x94\x99\xb0\xc1\xb2\xf9\x58\xc2\x35\x39\x0f\xd7\x06\x85\xd9\x11\x9e\xcb\xef\x1f\x83\xb9\x30\x6c\xdc\x9c\x85\xd0\x9f\x2a\x02\x1f\x9f\x34\x44\xef\x65\x52\xaf\xee\x97\x2f\x14\x32\x83\x3b\xef\x99\x15\xcb\x9d\x33\xb6\x9b\xf5\x48\xc6\x0b\x56\xeb\x0c\xba\xed\x4e\x60\xcf\x46\xe5\x4d\x73\x7e\xa0\xf9\xd9\xba\xa1\xe7\xab\xdb\xf9\xfb\x2c\xee\xec\xbc\xb3\xf3\xed\xcd\x8f\x33\xa8\x73\xfc\x2b\x47\x6d\x1e\xde\xe8\x9f\x04\x7a\x6b\x4d\xcc\x0e\xbe\x90\x86\x6e\x36\xcd\xcf\x7e\xc3\xdd\x4f\xcf\x3a\xe7\x28\x3f\x4b\x9c\xc3\xfe\xdf\xea\x1c\x94\xda\xaf\xe7\x1c\x45\x4e\xa3\x8d\xe2\x62\xfc\xa0\x19\x8b\xcb\x21\x57\xb6\xd5\x97\xae\x79\x67\xae\x9d\xb9\x56\x77\x1f\xde\x5c\x6f\x09\xc1\xb5\x65\xf1\x8a\xf9\xf7\x9e\xe5\xde\x9d\x0d\xef\x6e\x3e\xb2\xbd\xf4\x63\x0b\xf9\xf6\x9e\xe5\xda\x9d\x5d\xef\xae\x5d\x3f\x50\x8e\xbd\x6e\x70\x3f\xf8\x92\x28\x39\x1d\x10\xc2\xdc\xf8\xef\x36\xd7\x3e\x50\x98\x32\xc7\xf1\x97\xe2\xeb\xcd\xc1\x17\x23\x8b\xa6\x46\x0e\xd6\x4e\xca\xcf\x0b\x9a\x7b\x91\x92\x17\xdc\x76\x5e\xd4\xee\x45\xa5\xe1\x3c\x42\x9a\xb3\x42\xed\xb3\xb4\xdd\xdd\x60\xa7\x70\x9a\xdd\xe0\xc6\xfb\xed\xce\x30\xb3\x3b\xf3\x14\xd9\x44\xf5\x92\xbc\x66\xfb\x35\x73\xcf\xd2\xd2\x92\x39\x41\xc0\x0f\x76\xe8\x35\xb9\xce\x2d\x7a\xec\x1b\xd7\x6e\x2d\xbc\x53\x5c\xaf\x0c\xaf\x25\x76\xdd\x73\xb1\x5f\x61\xe0\x9e\x2c\xf5\x3b\x18\xf4\x9f\x0e\x06\x3b\x18\xec\x60\x70\x19\x47\x1d\x0c\xae\xcc\xf5\x37\x0f\x83\x77\x56\x09\x2b\x10\xdc\x93\x3a\x61\x07\x82\xfe\xd3\x81\x60\x07\x82\xdf\x2c\x08\x7e\xdd\xba\xe0\x0a\x31\x66\x2b\x61\xfa\x09\xdc\x56\x0e\x2c\x6c\x75\x10\xd6\x05\xeb\x17\x97\x16\x08\x07\x41\xa5\xb0\xbc\x56\x95\x0c\xc3\x4b\x0b\xb5\xc3\xbb\x0a\xf0\x15\x64\xec\x49\x09\xbe\x83\x0c\xff\xb9\x23\x28\x0e\x76\x0c\x3b\x16\x6d\x7d\xc7\xf8\xda\x2d\x34\x69\xfa\xf9\xee\x71\xb5\x3b\x53\xf8\x75\xf1\xa5\x78\xe2\xe4\x34\xb3\xc5\x27\x4f\x9b\x83\x8d\x05\x86\xf5\xb7\x75\xed\xc9\x96\x7b\xc7\x2c\xb1\xda\xed\xb6\xf7\x9f\x65\x59\xe0\xce\x04\x8e\x19\xc7\xeb\xc7\x5a\x94\xa3\xc8\xa7\x0d\xfd\xd8\xeb\xbf\x1c\xbd\x3b\x79\xd1\x68\x0c\xd0\x3b\x7d\x7b\xfe\xe6\xe8\xf5\xe2\xf5\xf3\x93\x17\xaf\xea\x13\xf4\xb1\xdf\xb4\x9a\xc2\x9c\x3c\xed\x35\x95\xd2\x2a\xfa\xd6\xb4\xb2\x26\x33\x19\x1b\xe3\xa5\x91\x9f\x70\x09\x00\x7e\x25\x8e\x34\xff\xfb\x01\xaa\x39\x5c\x18\x1c\xa3\x5a\x5e\xce\xe1\xc2\xfc\xeb\xf9\xba\x0c\xef\xe0\xcb\x2d\x2b\x83\xcd\xa6\xdb\x1c\xdc\x1a\xe4\xbf\x0e\xae\x99\x89\x26\x9b\x40\xce\xef\xb6\xe3\xc3\x62\xce\x53\x6d\x14\xb2\x29\x17\xe3\xf2\x9a\xfe\x7e\xf5\x40\x5e\xcc\x80\x13\x7a\x31\x52\x64\x8a\x64\x33\x7c\x81\xe5\x4a\xa2\x36\xd4\xf1\x77\x6f\x87\x39\xab\x9e\xe3\x19\x8a\x46\x4a\x01\x4d\x9b\xf4\xd4\x2c\x58\xad\x37\x94\x47\xa4\x77\x56\x47\x2d\xb8\x04\x0d\x6c\x6a\x1d\xba\x7a\x5b\xd1\xd1\x01\x27\x34\xc8\x04\xee\x12\xa5\xc3\xeb\x0e\xaf\x3b\xbc\xee\xf0\xba\xc3\xeb\x47\xc3\xeb\x4d\x5f\xfc\xd9\x8f\x1d\x86\xdd\x2e\x5d\xfb\xd9\x79\x90\xd9\x9d\x0a\xd2\x3f\x0f\xee\xb6\x10\x68\x8a\x8e\xab\xbf\x8b\xb5\x37\x9b\xb3\xba\x08\x62\x3f\x5d\x04\x69\x68\xed\x31\x50\xb9\x7a\x02\xb8\x09\x3e\x57\x45\xe3\x5d\xf7\x31\x57\xc0\x2d\xaa\xe6\x5d\x21\xd7\x7e\x3a\x8f\x5b\x99\x93\x0e\xb3\x61\x4b\x61\xa8\x08\x38\xeb\x1f\xd7\x54\xf8\xef\x9e\x1d\xd8\x54\x63\xbb\x8b\x3d\xfe\xb3\x3f\xab\xe4\x15\xa2\xc3\xfe\x1e\xd9\x54\x33\xce\xad\x6c\x07\xdd\xab\x63\x9b\x42\xa6\x3b\x97\x6c\x37\xee\xdd\xd8\x50\xb7\xad\xa3\x9b\x56\x37\xf8\x15\xf7\xb2\xed\xd5\xf1\x4d\x9d\xbd\xef\x93\xbd\x6f\xfd\xf5\xf2\xbb\x8c\x7f\x85\x84\xad\xa5\x9a\xbb\xc6\x0b\x73\x7b\x75\x3e\x4e\xe7\x2c\x77\x3b\xcb\xee\xac\xd0\x12\x25\xa7\xbb\xf2\x70\xaa\x3c\x2d\xf7\xab\x73\x12\x9a\xdd\x8e\xb0\xd4\xbd\xa5\xf6\x78\x5c\x7f\x0b\x6f\xa9\x6d\x04\x56\xfe\x4c\x14\x7b\xc6\x49\xf0\xd6\xc2\x86\x27\x9d\xec\x5d\xad\xa1\x2b\x33\xf8\xcf\xda\x60\xf1\xf8\xef\x4e\xed\x06\x1f\x8f\xf8\x42\xc2\xda\x11\x60\xb9\xf7\x97\x27\xce\x07\x86\x54\x5a\x5c\xcf\x9e\x5b\x3f\xca\x93\x23\x31\x0f\xfd\x7e\xc9\x0e\xbb\x65\x3b\xeb\x6c\xfb\xcb\x5c\xa5\x4d\x67\x5f\x26\x62\x68\x98\x33\x96\xe6\x78\x47\xc7\x9a\x27\x06\x75\x9d\xb9\x09\xfc\xf7\xa6\x35\xec\x15\x12\xfe\xca\x31\x8d\xdf\x30\xfd\xe9\x3e\x72\xd6\x4f\xd7\x6f\x10\x61\x4a\xb1\x3a\x50\xf4\xb8\xc1\xe9\xe2\x26\xc4\x55\xed\xb0\x19\x5a\x2e\x26\x08\x1a\xed\x1e\xbe\x84\xa4\x81\x29\xd3\x9f\xc0\xf2\x34\x68\x51\x43\x29\x4d\x83\x8c\xed\x70\x08\xc3\x5e\x32\x60\xc3\xde\x50\x50\x9b\xe0\xda\x68\x10\xd3\xd5\xa1\x78\x85\x0a\xe1\x2a\xb9\x02\x85\x99\x42\x8d\xc2\x68\x60\x7e\x64\x2e\x40\xcb\x29\x82\x92\xd2\xc0\x14\xb5\x66\x63\xec\xc3\x15\xbb\x02\x26\x62\xb8\x1a\x5d\x0d\x85\x6d\xa8\xa9\xa5\x99\x60\xd1\x06\x12\x99\x0b\xdb\xfd\x2a\xb9\xea\xbb\xc6\xf1\x55\x49\xb6\xbc\x6b\x26\x38\x14\x45\x1f\xdb\x7a\x30\xba\x1a\x10\x57\xbf\x96\x92\x6b\x60\x0a\x21\xd7\x18\x83\x91\xa0\x33\x8c\x78\x32\x07\x06\x3a\x1f\x85\x5a\xd2\x60\x26\xcc\x80\x9e\xc8\x3c\x8d\x61\x84\x43\xa1\xd0\xe4\x4a\x60\x0c\x23\x6a\x3e\xa6\xb6\x05\xac\x81\x54\x30\x95\x31\x4f\xb8\xbf\x2d\xc0\xa5\x35\x55\x93\x41\x83\x87\x54\x4b\x98\xb0\x19\x02\x83\x28\xd7\x46\x4e\xe1\xff\xde\xbd\x3d\x05\x14\x91\x8c\xb9\x18\xc3\x53\x8d\x08\x23\x4c\xe5\xf5\xf7\x56\x80\xef\xc0\x75\x7f\x63\xbb\x73\x01\x67\x4a\x92\xf5\x91\x63\xd2\xfd\xdf\x27\x28\x9c\x54\x5e\x75\x91\x14\x06\x3f\x5b\x81\x18\x64\x65\xe3\x3e\xb0\x12\x58\x4a\xf5\x4a\x35\x14\x3a\x1f\xfd\x50\xaa\x4e\x43\xc2\x53\x83\xca\x89\x43\xe4\x8e\xce\x4e\x48\x5f\x52\xa4\x73\x4b\x9a\xd9\x61\xa4\xc6\x42\x5d\x4c\x0f\x85\xd3\x26\xaf\x98\x20\x61\x07\xf0\xab\x54\x80\x9f\xd9\x34\x4b\xb1\x0f\x3c\x29\xef\x14\xad\x32\x85\x33\x2e\x73\x3d\x14\xbe\x15\x31\x60\x7f\xac\xc2\x4d\x52\x0b\xc7\x4c\x43\x22\xd3\x54\x5e\xeb\x43\x92\x9e\x8c\x31\x81\x2f\xee\x0b\x00\x83\x43\x78\xfe\xbc\xf8\x6b\x54\xdd\x00\x88\xe1\x10\x9e\x55\x7f\x7e\xa6\x96\xc5\x9f\x37\xc5\x97\x39\x35\xfa\x97\xfb\xcb\x5f\xfc\xfb\x10\xfe\x87\x46\x22\x6f\xf2\x9b\x62\xaf\x79\x9a\xda\x9f\xe8\x28\xf4\xe1\xc5\x8f\xc0\x06\x27\xe2\x50\x15\xda\xf9\xdc\x9f\x5b\xb3\xfd\x7b\x28\x9e\x9a\x09\x72\xe5\xda\x38\x1a\x23\xe7\xa0\x46\x5a\x6d\xf8\x94\xc1\x99\xb9\x9c\x72\x63\x9c\x3e\x6d\x48\x02\x9a\xd3\xa1\x90\xb9\xc9\x72\xf3\xbd\x95\x7d\x63\xe9\x6f\x4a\x01\x87\xe2\x88\xdc\x95\x96\x0f\x71\xe1\xaa\xda\x8a\xc6\x48\xc7\x94\x91\x7c\x8e\x30\x33\xc0\x8c\x65\x31\x65\xda\x40\x26\x35\x77\xc6\x9f\x00\x1b\x0a\x1b\x0b\xc0\xc5\x26\x6b\xb1\x27\x64\x78\x65\xfc\x04\x17\x2c\x0b\xba\x3e\x34\x90\x5c\x0d\x5f\xea\x3b\x2f\xae\x7c\xcb\x19\x82\xb6\x86\x90\xa6\x85\x42\x9f\x32\x4d\xa6\x54\x1b\x21\x09\x1a\x0c\xc5\x84\x91\xdf\x62\x39\x2d\x18\x3b\x4f\x3a\x95\x06\x9d\x73\xb3\x30\x20\xc6\x12\x1d\x6b\x02\x29\x87\x64\x8a\xa7\x73\x3b\xf6\xdc\x4f\xcc\x50\x18\x99\xfd\x90\xe2\x0c\xd3\x05\x7b\x1c\xc0\x89\x80\x88\x69\x74\xee\x76\x7e\xfc\xee\xa2\x55\xaa\x60\xbc\x42\xac\x98\x2b\x8c\x4c\x39\x4a\x49\xba\x0f\xa3\xdc\xea\xa7\x4e\x76\x28\x52\xae\x9b\x84\x0b\x6f\xd2\x06\x59\x1c\x2a\x0c\x59\x34\x01\x2e\x62\x3e\xe3\x71\xce\xd2\x82\xdd\xa1\xf0\x9e\x57\xc6\x34\x85\x5a\xe6\x2a\x42\x20\xea\x6d\xd2\xf8\x20\x35\x45\x33\x91\x71\x7f\x28\xa4\x99\xa0\x82\x20\x3b\x80\x29\x9b\x93\x25\x53\x0c\x1a\xc0\xef\x13\x82\x81\x92\xb5\x82\x25\x6f\xee\x43\x11\xa5\xc8\x54\x3a\x87\x58\x46\xf9\x14\x85\xb1\x8e\x3e\x46\x4b\xf4\x9a\x9b\x09\x70\xa3\x21\xc6\x28\x65\xde\x0a\x3c\xc3\x47\x67\x27\x03\x80\x13\x31\x14\x4c\xcc\x2d\x8f\x4e\x01\x98\x24\x64\x5d\x72\x89\x58\x07\xc5\x17\x4d\x06\x58\x24\x58\x43\x31\xc2\x09\x9b\x71\x72\x53\xa9\x88\xb6\x6e\x8f\xb5\xae\x5e\x08\x6f\x0b\xa5\x6b\xe7\x32\xc1\x6c\xf2\xc5\x88\xaf\x4b\xc3\xd3\x70\x3d\xe1\xd1\xa4\x30\x5d\x99\x78\x73\x62\x6a\x8c\x26\x54\x3e\x21\xd3\x58\x52\xf8\x37\xd2\xea\xd2\x92\x8c\x07\x70\xe1\x23\x70\x8d\xf9\x32\x1c\x4f\x98\x18\x3b\x65\xfb\xc0\xe3\x46\xa8\x22\x33\xb4\x06\x66\xd2\x61\x0c\x29\x12\x0a\xd1\x35\x3b\xa5\x1a\x72\x61\x64\x1e\x4d\x68\x58\xeb\xc2\x25\x77\x5c\x43\xc6\x74\x01\x31\x72\x28\x5c\x86\x30\x72\xbd\x3d\xaf\x9e\x85\x7e\x09\x1a\x7c\x2c\xa4\x22\x6b\xac\xb1\xc7\xd2\xb4\xc4\x7a\x17\x42\x67\x21\xda\x58\xdc\x28\x83\xc8\x62\x68\xaa\xe4\xa1\x89\x6b\x81\xdb\x3e\x08\xbc\x2e\xc6\x23\xab\xa3\xb9\x26\x33\x44\x11\x3b\x50\xb1\x56\xf3\x99\x6b\xe3\xde\x27\xa9\x8f\xe0\x94\xe4\x26\xa8\x54\xc0\x00\xca\xd0\x31\x14\x6d\x5c\xd9\xd9\x28\x22\xa6\xa7\xb1\x18\x2a\xe1\xca\x86\xca\xab\x85\x58\xd9\x40\xe0\x85\x39\x5b\x20\x55\xce\xb1\xb3\x41\xa7\x8a\xa1\x58\xd4\x85\xa1\xe4\xa0\x54\x48\x09\x3a\x53\x54\x63\x4b\xbe\xa9\x8e\x80\x93\x32\x54\x34\x95\x61\xb3\xaa\x10\xd8\xc7\x7c\x86\xb5\xa6\x9e\x42\x1b\x40\xd7\x41\xa9\x8e\xc8\x6d\x80\x1c\x1d\xc2\x1f\xcf\x3e\x86\x78\x25\xe2\x60\xda\xd7\x19\xe9\xc7\x56\xe2\xcf\x43\xe2\x56\x5d\x7c\x41\xbf\xdc\xe7\x1a\x3e\xeb\xfd\xc3\xa6\xbd\xc3\x5e\xdf\xe6\xbf\xd1\xb0\xf7\xb1\xec\x6b\x1a\x49\xc2\x68\x4d\xd6\x6e\xd3\x42\x1f\x9e\xd7\x15\x01\x9c\xf4\x4f\x41\xd4\x45\x4a\x8a\xc3\x99\x92\x33\x1e\x93\x0d\xb8\x48\x44\xa1\x62\x86\x4a\xd1\x35\x33\xe1\xba\x48\x32\x20\x0c\x80\x94\xe2\x7a\x93\xa6\xc0\x50\xe6\xde\xd6\x4b\x9d\x99\x52\xa6\x1b\xa3\x22\x7a\x04\xe0\x25\x84\xfe\xa7\xf6\xa6\xd5\xcc\x61\x02\x15\xe6\xda\x58\x27\x0c\x22\x90\x1d\x67\x31\xf5\xf1\xb4\xca\xec\xd0\xca\x12\x87\xa6\xf7\x0a\x45\x44\xb9\xe4\x22\x3f\x55\x86\x20\xc3\xe0\xd5\x2f\x55\xc2\x8a\x51\xc8\xb2\xb5\x61\x22\xc2\xc2\x97\xaa\x38\xec\xd9\x0a\xa8\x05\x5c\xf7\x29\xdd\x8f\xe5\x50\xb8\x44\xa6\x20\xeb\x04\x22\xd5\xba\xa8\x18\xbb\xf4\xbd\xf2\xef\x9a\x29\xd5\xb2\x20\x59\x58\xb2\x53\x58\x5b\xee\x33\x14\x6d\xc9\x4f\x40\xb2\x96\xfd\xc0\x84\xe9\x96\xe4\xa7\x4a\x7d\xaa\xcc\x5b\xa3\xd7\x80\x2b\x2b\x01\xce\x64\x9a\x17\x61\x83\x3b\x58\x9f\x22\x13\x3e\xee\x95\x8b\x23\x84\x28\xe5\xc4\x7c\x99\x38\x7d\x12\xf2\xda\xaa\x8e\x06\x27\x2c\xc1\x44\x2a\xb4\xb7\x12\x9e\xa6\x3e\xd2\x58\x17\xb1\x28\x86\xba\x74\x0f\x3f\x7b\xce\x0c\xb8\x2a\xa6\xc8\x22\x90\xe5\x82\x13\x2e\x5d\x33\x4a\x14\x2a\xd4\xa6\xb5\x4c\x99\x73\xfb\xdf\xd4\xb3\xfc\x7a\x8c\xa4\xb5\x95\xe3\xd1\xe6\x8e\xd7\x6c\x5e\x44\xd5\x39\x05\xf0\x4a\x79\xd6\x3a\xe2\x3c\xa2\xd8\xc7\x84\x2f\xbf\xf1\x84\x58\xb7\xd3\x77\xa4\x5d\x52\x52\x4b\xe9\x3c\xc6\xa5\x32\x62\x61\x40\x2e\x8d\xc8\x62\x7e\x05\x92\xba\x05\x25\xa1\x4c\xc3\x9c\x36\x0a\x97\x8b\x91\x70\x4a\xfb\x74\x26\x4c\x86\x3f\x71\x11\xdb\xf4\xac\x3d\xfd\x69\x89\x59\xe4\xd5\x2e\x63\xf0\x49\xc5\x44\x12\x20\x87\x0b\x3b\x97\xf4\x7c\x07\x2f\xa4\xd0\x3c\x2e\xd3\x17\xc2\x93\x57\x17\x17\x67\x3e\xef\x74\x6b\x1f\x7b\x81\xb8\xb0\x06\xb7\x08\xbc\x3e\xd5\xc9\x35\xea\xba\x7d\x96\xfe\xef\x27\xfa\xec\xe8\xe2\xc5\xab\x32\x67\x95\x09\x9c\xbd\xbf\xa8\x39\xb4\x66\x86\xeb\x64\xee\x46\xd4\x38\x65\xc2\xf0\x48\x0f\xc5\x53\x6a\x48\xd4\x1c\xde\xfa\x94\xd3\x2d\xb6\xf2\x34\xf5\x1c\xe9\x62\xd9\x6c\x97\xd5\xc7\xc5\xb2\x5a\x26\x61\x6e\xe7\x63\x1a\x35\xe9\x2f\x38\xa8\x5d\x8a\x53\x2c\x24\x41\x34\x17\xe3\x14\x3d\x64\xc3\xb5\xcd\x6e\xdd\x8a\x87\x29\x1c\x0a\x8d\x19\x25\xaa\x45\x71\x20\x92\xd3\x29\x1b\xb8\x91\x34\x08\x36\xb5\xc1\xcc\x26\xe3\xd4\xc9\xe6\x79\x91\x14\x33\x54\xc6\x25\x72\x07\x89\x92\x53\xa0\xc4\x41\xfd\x10\xb1\x29\xa6\xd4\x89\x46\xb2\xad\x84\x9d\x8f\xc2\x12\xc9\x40\x0b\xcc\x8d\xfc\x8c\xb9\x69\xb7\x0b\x62\xea\x55\xd9\x51\x99\x40\x57\xeb\xe4\xe2\xe6\x99\x92\x09\x4f\xb1\x02\xa1\xf7\x1a\x15\xe9\x52\xc1\xcf\xf0\xec\xa7\xe2\xea\xd9\x84\x96\x9d\x99\xfd\xf7\x67\x78\xfe\x53\x6d\x55\x5c\xd0\xb2\x5d\x4b\x42\x5e\x4b\x31\xd7\x59\xca\xe6\x97\x56\xfe\x90\xa4\xbf\xcf\xe2\x58\xa1\xd6\x75\xaa\x76\x42\xdc\x52\x97\x35\x93\x9b\x2b\xcf\xf2\x95\x75\xf2\x54\x4a\x1b\x6f\x75\x1e\x4d\x2a\xe1\xa8\x6d\xc9\x48\x59\x9f\x22\xa9\x06\x21\x3f\x45\x05\x2b\x68\x63\x45\x2c\xae\xdf\xd4\x4c\x83\xb4\xab\x49\x8a\xc2\x36\xca\xda\x96\xb3\x0f\x1b\xe8\x4b\x1e\xca\xe1\xa9\x75\x73\xf4\x53\x36\xc5\xfe\xe2\x50\xf5\x25\x07\x05\xd1\xb7\x02\x0b\x63\xd5\xcd\x0a\x96\x51\xc8\x4c\x80\x4c\x92\xda\x6a\xf8\x93\x9c\x82\x11\x77\xe3\x3c\x65\x45\xed\x61\x50\xfa\xb5\x5f\x83\x36\xcd\x64\xc1\x32\xde\x59\xeb\x7a\xe3\xff\x2a\xc5\xb1\xa3\x80\x41\x6d\x2e\xdd\xd7\x20\x7d\xf1\x53\xea\xa7\xfa\xdf\x3f\x55\x77\xde\xe5\xa3\x82\x92\xce\x47\x97\xc5\x18\x3f\xc3\xff\xfe\xd4\x56\x88\xb8\xa8\xc7\xaf\x88\x89\x30\x7b\x5a\x32\xbd\xe1\x8c\x5a\x2a\x6f\xd5\x5d\x5d\x02\x5e\x6a\x3d\x2b\x88\xf4\xe2\xce\x33\xb4\x62\x69\x78\x3a\xec\x55\xc2\x0f\x7b\x2e\x78\x73\x6d\xc3\xf0\xf7\xc4\x29\x21\x5d\x11\x8e\xb8\xf0\x25\x91\x22\xb8\x56\xf3\x0b\x1f\x50\x11\x64\x59\xbf\x2c\x64\x6e\x24\x71\x36\xb8\xce\xed\x02\xca\x2d\xb9\x7d\x5c\x9d\xd8\x68\x54\xd5\x3c\x2c\x7b\xe1\xa2\xc5\x85\x7b\x02\x13\x5f\xba\x9c\xd1\x60\x2e\xd4\x73\x11\xa5\x79\x5c\xae\x58\x2c\x7b\xae\xcc\xe4\x16\xcc\x76\xf5\x70\x75\x72\xfa\xe1\xe8\xf5\xc9\xcb\xcb\xa3\xf3\xdf\xde\xbf\x39\x3e\xbd\xb8\xaa\xf0\x90\x58\xb2\x01\xcc\xc2\xf1\x94\x65\x19\x1b\xa5\xc1\x23\xb6\xea\xf5\xf1\xab\x92\xc5\x46\x21\xd8\xd7\x58\xf5\x7c\x3a\x92\x29\x8f\xea\xac\x24\xd5\x7a\xe2\xb0\x57\xab\xc7\xd7\x9e\x61\xdd\xa3\x14\xdf\xf6\x62\xfd\x4a\xcf\x1b\x08\x0b\x96\xf5\x6b\x7b\xaf\x75\xf9\x5b\xad\x21\x55\x6f\x81\x9b\x30\x14\xa3\x61\x3c\xdd\xc2\x63\x85\xf6\x87\x86\xe1\x23\x9e\xf6\xa7\x84\xad\xcf\x4d\x5a\x4e\x21\xb8\xc7\x6c\x8d\x55\x16\x5d\x3e\x8c\xe6\x27\xc6\x64\x0f\x44\xfa\x1e\x93\x6a\xb9\xd2\x86\x99\x7c\xe9\xc4\xee\x9f\x4d\x2c\xff\x79\xe5\x7b\x3e\x3b\x7c\x30\x41\x97\xfd\x84\xdb\x6a\x92\x2e\xec\x15\xb8\x87\x98\x6d\x6f\x91\x6d\x59\xd4\xb6\xa3\x18\xd7\x13\x74\x5f\x26\x75\xd9\xbe\xc3\xbb\xa5\xf5\x6f\x92\xde\x53\xb4\x4d\x5c\xda\x6f\x23\xd8\x20\x14\x84\x4f\x6f\xd7\xee\xbe\x79\x0c\xa2\x9e\x1b\x71\x5c\xfd\xfe\xfd\x06\xcf\xf4\x73\x2e\xcc\x7f\xff\x7b\x09\x74\x07\xfb\xe2\x36\x22\xde\xb2\xdf\x2a\xa4\x1f\xee\x60\x7b\x08\xfa\x53\x34\x6c\xe9\x54\xb4\x9c\x66\xd4\x63\x71\x6c\x6d\x9e\xa5\x67\xcb\xcf\x34\xba\x23\xf4\xbd\x41\xc3\x3e\xd8\x3d\x16\x1b\x38\x4a\xed\x54\xbf\x7b\x38\xcd\x27\x2e\xe2\x4d\x4c\x69\xf1\xb7\x19\xe1\x21\xe0\x61\x89\x6a\x42\xef\xc5\xcf\xe6\x32\x38\xe3\x65\x55\x59\x6e\xd5\x2e\x4d\xcd\x7d\xb4\xfa\x95\x42\x8a\x2d\x79\x5d\xb6\xc5\xc1\x65\x5a\xb6\xf6\x57\x0f\xd6\x77\x6b\xe6\x43\x63\x63\xd0\xda\xea\xd9\x0e\x9f\x6b\x6e\x55\x5a\x51\xc0\x73\x4c\x50\xa1\x88\xee\x8d\xb2\x0f\x0e\x45\xb7\x8a\xf1\xae\x81\x11\x3b\x6d\xc2\x77\x4a\x52\xc3\xc9\x9d\x16\xa5\x09\xd3\x8f\x8e\x2a\x6e\xe2\x37\x40\x94\x8b\xba\xcd\xee\x93\x96\xf3\x96\xed\xcc\x8f\xa1\xe7\x5c\xdf\x8d\x4f\x5f\x2d\xb5\xb8\x9b\xb5\x87\xce\xda\xc2\x1d\xfc\x5b\xa3\x7f\xa7\x0d\xd7\xdf\x3c\xdb\x72\x10\x5f\x67\x31\xdb\x92\xc0\x4e\xeb\xdb\x5c\x6f\x21\xb9\xb8\x3b\x76\x45\x15\x6c\x4d\x7c\xd9\x5c\x90\xad\xa0\x80\x47\x10\x3e\x74\x2d\x66\xa2\xc9\xe5\x1d\x8b\x9a\xb6\x77\xd0\x6f\x7b\x03\xfd\x76\xed\x7e\xe0\x78\xdd\xa6\xd4\x86\x1d\x2f\x9e\x83\xb1\x70\x02\xc6\xe2\xd9\x17\xb5\x53\x2f\x3e\x06\x9b\x84\xdb\x4e\xba\x68\x32\x16\x9c\x19\x7a\xef\x39\xdf\x81\xac\x7e\x1d\xdf\xab\x44\x5f\x2d\x91\x2d\xab\x31\x5b\x28\x1a\x6d\x82\x58\xf6\x05\x8e\xb5\xdc\xaa\xca\x49\xdb\x95\x25\xb7\x4a\x6e\x27\x17\xc3\xdf\x50\xb1\xe1\x31\x61\xb1\x56\x8f\xfb\x0a\xc9\x5d\x9b\xad\xaf\xd2\x6f\xd1\xa4\x1f\x21\x95\xfc\x07\xa7\x6a\xdf\xaa\xcd\x7f\x95\x64\xb0\xbd\x08\xfe\xd8\xe9\x60\xfb\xcf\xad\x3c\x00\xe6\xad\xfe\xd4\xe3\xb1\x55\x50\x55\x8d\x36\x4a\xd9\xde\x5d\x9c\x9f\x9c\xfe\x16\x66\x68\x27\xa7\x17\xc7\xbf\x1d\x9f\x87\x97\x5e\xbe\x7d\xff\xcb\xeb\xe3\xf0\xca\x2f\x6f\xdf\xbe\x3e\x3e\x3a\xbd\x35\x93\xf3\xa4\x1b\xfc\x36\x32\x99\x8d\x98\x3e\x39\x3d\xb9\x08\xb9\x79\x71\x7e\x7c\x74\x51\xe3\xef\xfd\xd9\xcb\xc6\x95\x97\xc7\xaf\x8f\x2f\x8e\x6f\x65\xd8\x92\x7d\x52\x28\xf9\xe6\xc9\xcd\x93\xff\x0f\x00\x00\xff\xff\x9b\x7a\xb4\xa6\x9d\x97\x00\x00")

func serviceSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_serviceSwaggerJson,
		"service.swagger.json",
	)
}

func serviceSwaggerJson() (*asset, error) {
	bytes, err := serviceSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "service.swagger.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"service.swagger.json": serviceSwaggerJson,
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
	"service.swagger.json": &bintree{serviceSwaggerJson, map[string]*bintree{}},
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
