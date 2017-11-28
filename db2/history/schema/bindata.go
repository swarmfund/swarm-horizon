// Code generated by go-bindata.
// sources:
// migrations/0010_trades.sql
// migrations/0011_history_price.sql
// migrations/0012_update_pending_tx_op_key.sql
// migrations/0013_signer_name.sql
// migrations/001_squashed.sql
// migrations/003_update_forfeit_request_created_at.sql
// migrations/004_change_asset_length.sql
// migrations/005_change_exchange_kyc.sql
// migrations/006_pending_submitter.sql
// migrations/007_move_forfeit_requests_to_payment.sql
// migrations/008_use_random_salt_in_tx.sql
// migrations/009_use_long_assets.sql
// migrations/014_balance_history.sql
// migrations/015_account_type.sql
// migrations/016_balance_updates_index.sql
// DO NOT EDIT!

package schema

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

var _migrations0010_tradesSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\x31\x6f\x83\x30\x10\x85\x77\x24\xff\x87\x37\x82\x4a\x86\x2e\x5d\x32\x91\xc6\xaa\x50\x29\x89\x5c\x32\x64\x42\x57\x38\x05\x0f\x60\x6a\x1f\x8d\xd2\x5f\x5f\x45\x74\x40\x55\x53\xa9\xab\xdf\xe7\xbb\x7b\xdf\x6a\x85\xbb\xde\x9e\x3c\x09\xe3\x30\xaa\x48\x45\x8f\x46\x67\x95\x46\x95\x6d\x0a\x8d\xce\x06\x71\xfe\x52\x8b\xa7\x96\x03\x62\x15\x01\x80\x6d\xb1\xc9\x9f\x5e\xb5\xc9\xb3\x02\xe5\xae\x42\x79\x28\x8a\x74\xce\xde\x28\x70\x4d\x21\xb0\xa0\xe9\xc8\x53\x23\xec\xf1\x41\xfe\x62\x87\x53\x7c\xff\x90\xfc\xe4\xdf\x27\x27\xff\xfa\x30\x2f\xe8\xdd\x34\xc8\xf5\x8a\xbc\xac\x6e\x8c\xfc\x0b\x19\xbd\x6d\xf8\x46\xd6\x78\x26\xe1\xb6\x26\x81\xd8\x9e\x83\x50\x3f\xe2\x6c\xa5\x73\xd3\xfc\x82\x4f\x37\xf0\x37\xbc\x37\xf9\x4b\x66\x8e\x78\xd6\x47\xc4\xb6\x4d\x17\xfd\xd3\x65\xb7\x44\x45\xc9\xfa\xaa\x77\x29\x7c\xeb\xce\x83\x8a\xb6\x66\xb7\xff\x55\xf7\xfa\x2b\x00\x00\xff\xff\x07\xa9\xe1\xb1\x9d\x01\x00\x00")

func migrations0010_tradesSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0010_tradesSql,
		"migrations/0010_trades.sql",
	)
}

func migrations0010_tradesSql() (*asset, error) {
	bytes, err := migrations0010_tradesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0010_trades.sql", size: 413, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations0011_history_priceSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x90\x31\x4f\xc3\x30\x10\x85\x77\x4b\xfe\x0f\x6f\x4c\x44\x3a\xb0\xb0\x74\x0a\x34\x03\x22\xb4\x95\x95\x0e\x9d\x2a\xd7\x9c\x9a\x93\x88\x6d\xec\x0b\x55\xf9\xf5\x08\x3a\x34\x8a\x58\xd8\x4e\xa7\xf7\x74\xdf\x7d\x8b\x05\xee\x06\x3e\x25\x2b\x84\x5d\xd4\x4a\xab\x27\xd3\xd4\x5d\x83\xae\x7e\x6c\x1b\xf4\x9c\x25\xa4\xcb\x21\x26\x76\x84\x42\x2b\x00\x38\xda\x4c\x07\x9b\x33\x09\x5c\x6f\x93\x75\x42\x09\x9f\x36\x5d\xd8\x9f\x8a\xfb\x87\x12\xeb\x4d\x87\xf5\xae\x6d\xab\x6b\xfe\x63\x0c\xf2\xaf\x82\xf0\x40\x59\xec\x10\x27\xd3\x99\xa5\x0f\xa3\xfc\x6e\xf0\x15\x3c\xcd\x4b\x57\xc4\xb7\x30\x1e\xdf\x09\x31\x91\xe3\xcc\xc1\xcf\x53\x5b\xf3\xfc\x5a\x9b\x3d\x5e\x9a\x3d\x8a\xdb\x23\xd5\x14\xb2\xba\x9d\x2d\xb5\x2a\x97\x3f\x56\xa6\x9e\x56\xe1\xec\xb5\x5a\x99\xcd\xf6\x2f\x4b\xcb\xef\x00\x00\x00\xff\xff\x91\x95\xf9\x33\x53\x01\x00\x00")

func migrations0011_history_priceSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0011_history_priceSql,
		"migrations/0011_history_price.sql",
	)
}

func migrations0011_history_priceSql() (*asset, error) {
	bytes, err := migrations0011_history_priceSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0011_history_price.sql", size: 339, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations0012_update_pending_tx_op_keySql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x92\x5f\x4b\xc3\x30\x14\xc5\xdf\x03\xf9\x0e\xf7\xad\x2b\x3a\xd0\xe1\x9e\xc6\x1e\xaa\x0d\x58\x08\x6d\xed\x12\xd4\xa7\x10\xb6\x30\x8b\x35\xad\x69\x3a\xeb\xb7\x97\x74\x13\x5b\x1d\xb8\xf5\xa9\x7f\x72\xcf\xb9\xbf\x1c\xce\x74\x0a\x17\x6f\xf9\xd6\x48\xab\x80\x57\x18\x05\x94\x91\x0c\x58\x70\x4b\x09\x54\x4a\x6f\x72\xbd\x15\xd6\x48\x5d\xcb\xb5\xcd\x4b\x5d\xc3\x7e\xc0\xb6\x42\xe9\x9d\x2a\xca\x4a\x01\x7b\x4e\x09\x58\xd5\xda\xc5\x39\x72\xa3\xea\xa6\xb0\x63\xc4\x65\xa5\x8c\x74\x3f\xc4\xab\xfa\x84\x30\x4b\x52\x88\x13\x06\x31\x2f\x8a\x05\x46\x18\xdd\x65\x24\x60\x04\x78\x1c\x3d\x70\x02\x51\x1c\x92\x27\x67\x27\x6c\x2b\x06\x52\xd1\xe8\xfc\xbd\x51\x22\xd7\x1b\xd5\x62\x04\x90\xc4\xc7\xd7\x4e\x06\x32\xdf\x8d\x3e\xde\x93\x8c\xfc\x3a\x80\x68\x75\x00\xa1\xd4\xef\x48\x30\xea\xc7\x1b\x96\x1f\x1a\xa3\x8e\xf7\x44\xa8\xce\x84\xa7\xa1\xbb\xce\x51\xb2\x15\x61\xc3\x38\x96\x5e\x59\xb9\xa7\x77\x20\xfc\x0b\xc8\x29\x1d\x1b\xb5\xdb\x36\x48\xfa\x3f\xb4\x5e\x4b\x96\x5e\xef\xc3\xbb\xfc\x69\x40\x77\xb0\x7f\xf5\xce\x2a\xd0\xb0\x7f\x3b\x69\xd6\x2f\xd2\x4c\xae\xaf\x66\x37\xfe\xe8\x22\x7e\xbb\xcc\xe6\x73\x67\xf2\x15\x00\x00\xff\xff\xe1\xa8\x25\xb5\x1d\x03\x00\x00")

func migrations0012_update_pending_tx_op_keySqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0012_update_pending_tx_op_keySql,
		"migrations/0012_update_pending_tx_op_key.sql",
	)
}

func migrations0012_update_pending_tx_op_keySql() (*asset, error) {
	bytes, err := migrations0012_update_pending_tx_op_keySqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0012_update_pending_tx_op_key.sql", size: 797, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations0013_signer_nameSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xcd\xa1\xae\xc3\x20\x14\x06\x60\x4f\xc2\x3b\xfc\xae\xe2\xa6\x4f\x50\xc5\x1d\x4c\x9d\xc1\xd2\x80\x26\x64\x23\x04\xd1\xd3\x06\x48\xb6\xc7\x9f\x98\x99\x98\x98\xfd\xcc\x37\xcf\xf8\xdb\x6a\x69\x69\x64\x84\x43\x0a\x45\xde\xac\xf0\xea\x9f\x0c\x8e\xcc\xf7\xca\x25\x8e\x96\xb8\xa7\xdb\xa8\x3b\xc7\x5e\x0b\xe7\xd6\xa1\xb4\xc6\xc9\x51\xb8\x58\xbc\x29\x72\xda\x32\x46\x7e\x0e\x58\xe7\x61\x03\x11\xb4\x39\xab\x40\x1e\xd3\xb4\x48\x21\xc5\xe7\xa5\xf7\x07\xff\xbe\xe9\xd5\x5d\xbf\x74\xcb\x2b\x00\x00\xff\xff\x84\x1c\x57\x20\xbe\x00\x00\x00")

func migrations0013_signer_nameSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations0013_signer_nameSql,
		"migrations/0013_signer_name.sql",
	)
}

func migrations0013_signer_nameSql() (*asset, error) {
	bytes, err := migrations0013_signer_nameSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/0013_signer_name.sql", size: 190, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations001_squashedSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x5b\xdf\x6f\xdb\xbe\x11\x7f\x0f\x90\xff\x81\x6f\x49\xb0\x64\x58\x81\x2e\xd8\x9c\x27\x37\x51\x57\x63\x8e\xbc\xc6\xf6\xda\xa0\x28\x08\x5a\xba\x38\x5c\x65\x51\x25\xe9\xc4\xee\xb0\xff\x7d\x90\xf5\xc3\x92\x28\x8a\x94\x23\xb9\xdf\xc7\x50\xc7\xbb\xfb\xf0\x7e\xf0\xee\xcc\x5c\x5d\xa1\x90\xd1\x50\x44\xe0\x49\xca\x42\x34\xfd\x19\xb8\xec\x8e\x48\x32\x65\x6b\xee\xc1\x28\xff\xf2\x91\xf1\x8f\x34\x80\xd3\x93\xd3\x93\xab\x2b\xf4\xa7\x15\x5d\x72\x22\x01\xcd\xa3\x78\x25\x5e\xdb\xad\xbb\x64\x05\x03\xf4\x4c\x85\x64\x7c\x8b\x89\xe7\xb1\x75\x28\xc5\x0d\x9a\x6d\x23\x18\xa0\xd9\xf0\xc3\xd8\xb9\x41\x53\xef\x19\x56\x64\x80\xa2\xf5\x22\xa0\xde\x0d\x9a\xbc\x86\xc0\x07\x28\x62\x42\x2e\x39\x88\x84\xd9\xe9\xc9\xed\x83\x33\x9c\x39\xc9\x26\x85\x25\x3a\x3f\x3d\x41\x08\x21\xea\xa3\x05\x5d\x0a\xe0\x94\x04\xc8\x9d\xcc\x90\x3b\x1f\x8f\x2f\x93\x6f\xc4\xf7\x39\x08\x81\xbc\x67\xc2\x89\x27\x81\xa3\x17\xc2\xb7\x34\x5c\x9e\x5f\xbf\xbf\x38\x3d\xb9\xb8\x69\x50\x7d\x41\x02\x12\x7a\xd0\xa5\xea\x19\xcb\xb2\xea\x34\x94\x55\xbd\x53\x42\x4c\xfd\x1a\xd5\xff\x7a\x7d\xa1\xe0\x14\x02\x64\x0d\xe9\x7b\x95\x32\x39\x3d\x6b\xce\x3f\xb6\x1e\xfa\x8f\x60\xe1\x22\xff\x60\x79\x6c\x98\xfa\x58\xc0\xcf\xec\xf4\xa6\xce\xe7\xb9\xe3\xde\xb6\x3e\xc0\x6c\x9f\x8e\x7f\xa2\xe5\x74\x36\x7c\x98\xa1\x2f\xa3\xd9\x27\xf4\x2e\x59\x19\xb9\xb7\x0f\xce\xbd\xe3\xce\xd0\x87\xc7\x6c\xcd\x9d\xa0\xfb\x91\xfb\xef\xe1\x78\xee\xec\x17\x86\x5f\x0b\x0b\xb7\xc3\xdb\x4f\x0e\x7a\xf7\x06\x7c\x68\xf2\xc5\x75\xee\xd0\x87\x47\x6b\xa0\xc3\xf1\xcc\x79\x30\xe2\xcc\xf9\x2a\x04\x7f\xa6\x7e\x93\xbe\x4f\x8c\x3f\x01\x95\x98\xc3\xcf\x35\x88\x4e\x23\xb1\xca\xda\xe8\xd6\x92\xf0\x65\xad\x9f\x5e\xd7\x38\xea\x2a\xf6\x53\x3b\x5a\x1a\x52\x49\x89\x04\x1f\x2f\xb6\x78\x2d\x80\xa3\x05\x63\x01\x90\xb0\xc6\xfb\x21\x92\xe0\x67\xdf\xd3\x65\x8f\xc3\x6e\x37\x91\x48\xd2\x15\x08\x49\x56\x11\x7a\xa5\xf2\x99\xad\x93\x15\xf4\x8b\x85\x90\x12\xaf\x23\xdf\x9e\x38\x3b\x22\xb9\x8d\x00\xd1\x50\xc2\x12\x38\xba\x73\x3e\x0e\xe7\xe3\x19\xfa\x8b\x6d\x4c\x55\x0f\xba\xaf\xd8\xd2\xc8\x39\x7a\x8c\x59\xe2\xed\x2c\xd6\x34\xf2\xd4\x98\xab\x12\x1a\x62\x2f\x00\x7f\x09\xbc\xcb\x90\x4b\x39\x66\x91\x26\x62\x3d\x42\x6f\xef\x5a\x15\x7f\x4f\xc8\xf1\x33\x11\xcf\x76\x91\x14\x71\x78\xa1\x6c\x2d\xb0\x71\x67\x16\xd2\x9c\x84\x82\xec\xea\x03\xbc\xbb\x59\x1a\xbc\x3c\xdd\xc2\x22\xe0\xa4\xcd\x06\x2f\x60\xc2\x18\x70\x4a\x4e\xc8\x92\x50\xf6\xf7\x2a\x62\x5c\x02\xc7\x2f\xc0\x45\x5c\xe7\x54\xc5\xbe\x53\xb2\x15\x93\x24\xc0\x5e\x5c\x1b\x69\xf2\xd9\x13\x00\x8e\x18\x0b\xb4\xb7\xb8\x00\xfc\x04\x5a\xeb\xec\xbe\x73\x10\xc0\x5f\xb4\x34\x2b\xb2\xc1\x72\x83\x05\x48\x2c\xe8\x2f\x95\xcc\x90\x39\xf6\x67\x1d\x11\x2e\xa9\x47\x23\xd2\x6d\x35\x56\x2f\xa0\x70\x13\x68\x70\xa9\xfb\xb5\xb7\x46\xa5\xf0\xd3\x13\x36\x56\x4d\xb1\xb3\x67\xa6\x3e\x3b\x1b\x0c\x14\x8a\x2a\x37\x78\x7a\x02\x4f\x8a\xa4\xfa\x39\xe8\x98\xfb\x4a\xd3\x8d\xd2\x8e\x9e\xac\x5b\x61\xef\x2c\x65\x37\x4a\x55\x13\x77\x3d\xb9\x21\x7d\xe7\x9b\x7a\x09\x18\x8b\x72\xa9\x90\x5b\xb5\x44\x24\x8a\x02\xea\x25\xd8\x18\xf7\x81\xeb\x22\xae\x54\x82\x54\xbe\xf9\x20\x09\x0d\x52\x5f\x4f\xd7\xc4\xae\xfd\xcb\xa2\xae\x8b\x88\x4a\xef\x94\x5d\x3a\xc7\xbb\xc4\xdd\x94\xcf\x33\xd6\x21\x7b\x3d\x57\x2b\x3e\x1f\x42\x49\x9f\x68\x5c\xea\xd5\x1e\x8b\x90\x71\x6f\x9a\xc2\xb5\x8d\xde\xd8\x7f\x32\xbe\xfd\x46\x6e\x55\xd2\xef\x8b\x5a\x13\xe6\xee\x23\xb6\x2a\xb1\x21\x5a\xe3\x08\xcd\x48\x9b\xd0\x44\x64\xbb\x82\xb0\x97\x26\xa7\xca\xda\x18\xb5\xd9\x06\x7d\xc4\x1e\xbb\x09\x29\x45\xb7\x6d\xe5\x50\xc5\xdd\xd7\x65\xa6\x91\x73\xf4\x80\xb0\xc4\xdb\x59\x38\x68\xe4\xa9\xc1\x50\x25\x34\x5c\x5a\x0a\xdf\xbd\x3f\x1e\xc7\x7e\x65\x79\xbf\xdf\x8e\xcd\xf8\xfb\xb3\x67\x59\xae\x85\x5d\xf7\x1b\x9a\xf0\x15\x8b\x82\x9e\x6a\x7a\x9d\x08\xfb\xaa\xde\xaa\x72\x31\xd7\xf5\x86\x14\xa5\xd3\xb3\x2f\x57\x37\xc8\x3b\xba\xab\xb7\xc4\xdf\x99\xab\x1b\xe4\xaa\xae\xae\xdb\x60\x48\x65\x85\x6d\x3d\xb9\x77\xee\xd2\x45\x0d\xed\xa7\x25\x69\x41\x6b\x9a\xc2\x58\x97\xe8\xcd\xb5\xb6\x66\x92\x9f\x4b\x6f\x98\x51\x10\x7d\xc8\xea\xa6\x31\x6f\xa9\xdd\x35\x23\x18\xb9\xc1\x10\xbe\x40\xc0\x22\x40\x12\x36\x6a\xbf\xb3\xc1\x1c\xc4\x3a\x90\xba\xaf\x2b\x90\x44\xf7\x2d\xc6\xa9\xfd\x2e\xe8\x32\x24\x72\xcd\xa1\xee\xa7\xa0\xbf\x5f\x5f\x7c\xfb\xbe\xef\x64\xfe\xfb\xbf\xba\x5e\xe6\xdb\x77\x65\x32\x03\x2b\x96\xcc\x75\xd5\xc6\x27\x67\x16\xb2\x10\x6c\x5a\xa3\x98\x99\xca\x27\x43\x47\x57\x80\x17\x6c\x1d\xfa\x22\x36\xcf\xdf\x38\x09\x97\x60\xf9\x33\x0c\xa2\x7e\x16\x3a\xa9\x52\x2d\xa3\x3f\x89\x9d\x89\x3b\x56\x7f\x79\x40\x09\xc5\xed\x64\x3c\xbf\x77\x63\x6b\x4f\x9d\xd9\xbe\x6d\x83\x8d\x7c\x21\xc1\xf9\x99\xe6\x07\x8d\xb3\xc1\x80\xc3\xd2\x0b\x88\x10\x09\x10\xab\x49\x70\x1f\x70\x14\x19\xad\x60\x69\x66\xc7\x0a\xbc\x56\xd3\x93\x3e\x60\x6a\x24\xb5\x02\xdb\x38\x75\x69\x0d\x59\x14\xda\xf8\xfe\xe0\x2a\x10\xf3\xc9\x81\x1d\xd4\x6a\xbb\x6a\x0b\x53\x69\x1c\x7b\xb0\xa9\x22\xa3\x95\x35\x35\x2d\xc8\xc1\x00\x0b\x05\xec\x71\x81\x16\x5a\xee\x76\x80\xcb\x35\xba\x2d\x70\x6d\x7d\xdc\x03\x70\xad\xac\x56\x96\x36\x54\x6c\xb6\xc0\xf3\xc4\xaf\xe4\xf4\xfd\xe8\x1d\xff\x80\x6d\x76\x0a\xb7\x13\x77\x3a\x7b\x18\x8e\xdc\x0e\x2f\x9d\xe4\x4a\x1c\xde\xdd\x15\xb8\x9b\xf4\x41\x73\x77\xf4\x79\xee\xa0\xf3\xfd\xf2\x81\x38\xa3\x3f\x04\xba\x58\x0b\xf4\xaf\x87\xd1\xfd\xf0\xe1\x11\xfd\xd3\x79\x44\xe7\x06\x40\xca\x15\xa7\xbd\xbd\xfa\x02\x58\x15\xd4\x08\xb4\x56\xab\x56\x80\x35\x97\x9d\xe1\x1e\xeb\x0b\x7c\xbd\xb8\xc6\x23\x68\xd0\xb0\xd5\x41\x28\x89\x53\x9f\x0b\x7b\x02\x5f\x15\xd4\x08\xbb\x56\xab\x56\x80\xb5\x09\xd3\x98\x0b\xfb\x3a\x00\x9d\xc0\xc6\x83\x68\xd4\xd2\xea\x40\x16\xf9\x64\x25\x03\x35\x72\xef\x9c\xaf\x6d\x7b\xe7\xdd\xa6\x02\x33\x34\x71\xeb\x3b\xe9\xf9\x74\xe4\xfe\x03\x2d\x24\x07\x40\xe7\x29\xf1\xa5\xd2\xa5\xea\x54\x8d\x9b\xee\x6e\xf4\xdc\xb5\xef\x56\x4a\x56\x9b\x7e\x9d\x6e\x49\xef\xdb\x8d\x76\x09\x2f\x3b\xfd\x2a\xe3\x85\x4b\x75\x90\xa0\x8d\x04\xcc\x22\x1c\x15\x0a\xb1\x83\x94\x4e\x6f\xcd\x44\xf7\x22\xd7\xa2\xfa\x9a\x4c\x5b\x02\xa2\x4e\xfa\x2e\x6b\x1f\x00\x5c\x22\x8b\x0b\x1a\xcb\x4d\x1f\xd0\x52\xae\x1a\xcb\x1c\x08\xae\x3c\x07\xad\x87\xc4\x22\x1c\x7b\x2d\x7b\x23\xa2\x14\xca\x9e\xdb\xa1\x46\x2a\x1a\xa4\x5e\xe3\xfc\x75\xd0\x62\xdb\xad\x1d\xca\x8c\x8b\x00\xb2\x77\x4f\x25\x8d\xf5\xfa\x15\xcf\xbd\x7b\x25\x15\xee\x76\xe1\xac\x53\x57\x26\x26\x93\xdd\x38\xc0\x9e\xdb\xe1\xae\x6c\xe1\xb6\x34\xf4\x61\x83\xab\x6f\xb2\x31\x0b\x71\xfa\xe6\xba\xbb\x13\x37\x8a\x2a\x02\xcd\x5f\x87\x97\xaf\xa3\x84\xb0\x25\x92\x2e\xdd\xa6\x49\x8a\x59\x7f\x2b\x23\xa4\x31\x12\xf3\xcc\x5f\xcc\x75\xe0\x51\x66\x21\xc6\x40\xcd\x29\xdb\xa1\xe8\xef\xfc\x4b\x42\x0e\xcd\x33\x7a\x96\x95\xa7\x86\x7d\x1a\x41\x79\xd6\x68\x04\x53\xd9\xd0\x0e\x5a\xe1\x59\xe8\x11\x6c\x53\x7c\x84\x6a\xc2\x55\xa0\x6d\x07\xa9\xee\xc9\xeb\x11\xb0\xd5\xbe\xb4\x35\x81\xac\xdb\xd4\x0e\x6d\x56\x4c\x1e\x01\x61\xfe\xc3\x94\x09\x55\x63\x6f\x50\x66\x5f\x98\xca\xf6\x9a\x20\xaa\x72\x6a\x8b\xa9\x43\xd2\x44\x99\x71\xf9\xa2\xed\x3c\x4f\x34\x09\xb3\x41\xd4\xba\x0e\xa8\x08\xdc\x46\x6f\xf3\x33\x0b\x4c\xdb\x08\xac\x90\x6c\x23\x1b\xf7\x2a\x16\x6f\xfd\x3a\x98\x2a\xa9\x7d\x11\xb9\xff\x1f\xbc\x3b\xf6\x1a\xc6\x6b\x3e\x67\x11\x92\x64\x11\x80\x5a\x50\x78\x44\x78\xc4\x87\x9b\x5a\xaa\x7c\xe2\xd9\x48\xa5\x8c\x11\x1b\xa9\xb3\x58\x6f\x24\xd2\x74\x26\x76\x7b\x0c\x74\xca\xec\xab\x91\x5a\x5b\x22\xdb\xee\xda\x53\xfe\x3f\x00\x00\xff\xff\x50\xcc\x93\x12\x53\x39\x00\x00")

func migrations001_squashedSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations001_squashedSql,
		"migrations/001_squashed.sql",
	)
}

func migrations001_squashedSql() (*asset, error) {
	bytes, err := migrations001_squashedSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/001_squashed.sql", size: 14675, mode: os.FileMode(438), modTime: time.Unix(1511895155, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations003_update_forfeit_request_created_atSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xcf\xb1\x6a\x85\x30\x14\xc6\xf1\x3d\x90\x77\x38\x63\x4b\xf1\x09\x9c\xd4\xa4\x56\x50\x23\x69\x84\x76\x92\xd0\x1e\x6b\x06\x8d\x4d\x8e\x48\xfb\xf4\x77\x70\x11\xee\x1d\x5c\x3f\xfe\xc3\xef\x4b\x12\x78\x99\xdd\x4f\xb0\x84\xd0\xaf\x9c\x65\xb5\x91\x1a\x4c\x96\xd7\x12\x26\x17\xc9\x87\xbf\x61\xf4\x61\x44\x47\x43\xc0\xdf\x0d\x23\x45\x38\xa2\x42\xd5\x7d\xd3\xc2\x57\x40\x4b\xf8\x3d\x58\x02\xf3\xd9\x49\xc8\xab\xb2\x6a\x0d\xf4\xef\x55\x5b\x82\xfc\x30\x3a\x2b\xcc\x93\xec\x54\xf1\x06\xaf\x5a\x35\xa7\xfe\x39\xe5\x8c\xb3\xb3\x40\xf8\x7d\xb9\x68\x10\x5a\x75\xf7\x84\xf4\xea\x03\x21\x1e\xf8\xc9\xcd\x18\xc9\xce\x2b\xec\x8e\x26\xbf\x1d\x0b\xfc\xfb\x05\x53\xce\x6e\x01\x00\x00\xff\xff\xe2\x21\x0a\x66\x2b\x01\x00\x00")

func migrations003_update_forfeit_request_created_atSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations003_update_forfeit_request_created_atSql,
		"migrations/003_update_forfeit_request_created_at.sql",
	)
}

func migrations003_update_forfeit_request_created_atSql() (*asset, error) {
	bytes, err := migrations003_update_forfeit_request_created_atSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/003_update_forfeit_request_created_at.sql", size: 299, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations004_change_asset_lengthSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe5\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xc8\x2c\x2e\xc9\x2f\xaa\x8c\x4f\x4a\xcc\x49\xcc\x4b\x4e\x2d\x56\x80\x48\x3a\xfb\xfb\x84\xfa\xfa\x29\x24\x16\x17\xa7\x96\x28\x84\x44\x06\xb8\x2a\x24\x67\x24\x16\x25\x26\x97\xa4\x16\x29\x94\x25\x16\x55\x66\xe6\xa5\x6b\x98\x68\x5a\xf3\x72\xf1\x72\x21\x9b\xee\x92\x5f\x9e\x47\x35\xf3\x8d\x35\x15\x42\x83\x3d\xfd\xdc\x15\x8a\x4b\x93\x8a\x4b\x8a\x34\xc0\x8a\x75\x14\x0c\x75\x14\x8c\x35\xad\x01\x01\x00\x00\xff\xff\xc6\x71\xf2\x2d\xd4\x00\x00\x00")

func migrations004_change_asset_lengthSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations004_change_asset_lengthSql,
		"migrations/004_change_asset_length.sql",
	)
}

func migrations004_change_asset_lengthSql() (*asset, error) {
	bytes, err := migrations004_change_asset_lengthSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/004_change_asset_length.sql", size: 212, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations005_change_exchange_kycSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe5\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xc8\x2c\x2e\xc9\x2f\xaa\x8c\x4f\x4a\xcc\x49\xcc\x4b\x4e\x2d\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\xae\x4c\xb6\xe6\xe5\xe2\xe5\x42\x36\xc0\x25\xbf\x3c\x8f\x80\x11\x8e\x2e\x2e\x48\x26\x28\x64\x15\xe7\xe7\x25\x59\x03\x02\x00\x00\xff\xff\x92\x1b\xba\x5b\x85\x00\x00\x00")

func migrations005_change_exchange_kycSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations005_change_exchange_kycSql,
		"migrations/005_change_exchange_kyc.sql",
	)
}

func migrations005_change_exchange_kycSql() (*asset, error) {
	bytes, err := migrations005_change_exchange_kycSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/005_change_exchange_kyc.sql", size: 133, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations006_pending_submitterSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x4d\x6b\x32\x51\x0c\x85\xf7\x03\xf3\x1f\xb2\x54\xde\x57\x68\xa5\x76\xe3\xca\xb6\xee\xc4\x42\xd1\xf5\x25\xce\x84\x3b\xa1\x63\xee\x25\x37\xa3\x4e\x7f\x7d\xb1\x63\x3f\xc0\xcf\xed\xe1\x39\x27\x9c\x24\x83\x01\xfc\x5b\xb3\x57\x34\x82\x65\xcc\xb3\x3c\x7b\x7e\x9b\x4e\x16\x53\x58\x4c\x9e\x66\x53\x88\x24\x25\x8b\x77\xa6\x28\x09\x0b\xe3\x20\x2e\xb1\x17\xd2\x04\xbd\x3c\x03\x00\xe0\x12\x56\xec\x13\x29\x63\xfd\xbf\x93\x4e\xb9\x3a\x8c\xc5\x60\xfe\xba\x80\xf9\x72\x36\x3b\xc0\x5d\x9c\xe3\x92\xc4\xd8\xda\xcb\x54\x6c\x56\x35\x17\xee\x9d\x5a\x28\x2a\x54\x2c\x8c\x14\x36\xa8\x2d\x8b\xef\x3d\x3e\xf4\x7f\x6c\x79\xd6\x1f\xdf\xd2\xe6\x52\x0d\xdb\xb9\x0a\x53\x75\x65\xd2\x2f\x4c\xb2\xa1\x3a\x44\x3a\x61\xb8\xbf\x1b\x1e\x5b\x42\x24\xc5\xaf\xe5\x58\x1b\x09\x58\x8c\x3c\xe9\x79\xea\x7a\xeb\xef\x65\xd9\xfe\x9c\x67\xf2\x6c\xe7\x94\x52\x53\xdb\x89\xac\xe1\x68\x74\x14\x56\x28\xa1\x51\xe9\xd0\xc0\x78\x4d\xc9\x70\x1d\x61\xcb\x56\x85\xa6\x53\xe0\x23\x08\x1d\xe0\x26\x96\xb7\xc3\x2c\x6c\x8c\x16\xf4\xc6\x5b\xfe\xfd\xd5\x97\xb0\x95\xbd\x56\x6a\x88\x60\xb8\xaa\xe9\xd2\xaf\x8e\xaf\x81\x7b\xe2\x33\x00\x00\xff\xff\x39\x71\x2f\xbb\x0b\x03\x00\x00")

func migrations006_pending_submitterSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations006_pending_submitterSql,
		"migrations/006_pending_submitter.sql",
	)
}

func migrations006_pending_submitterSql() (*asset, error) {
	bytes, err := migrations006_pending_submitterSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/006_pending_submitter.sql", size: 779, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations007_move_forfeit_requests_to_paymentSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\xc1\x4e\xeb\x30\x10\x45\xf7\x96\xfc\x0f\x77\xd9\xea\xbd\x4a\x2c\x10\x9b\xae\x42\x13\x56\xa6\x45\x55\xb2\x8e\xdc\x64\x9a\x8e\x44\x6c\xe3\x4c\xa8\xc2\xd7\x23\xda\x06\x50\x41\x22\xdb\x33\xc7\xf6\xe8\x5e\x2f\x16\xf8\xd7\x72\x13\xad\x10\x8a\xa0\x55\xba\xdd\x3c\x21\x4f\xee\x4d\x86\x03\x77\xe2\xe3\x50\xee\x7d\xdc\x13\x4b\x19\xe9\xa5\xa7\x4e\xba\xa5\x56\x89\xc9\xb3\xed\x95\x16\xec\xd0\x92\xfb\xd2\x90\xa4\x29\x56\x1b\x53\x3c\xae\x71\x61\xa5\x0c\x81\xc0\x4e\x96\x5a\x69\xf5\xfd\xe9\xd4\x1f\xdd\x07\x5b\x6d\xb3\x24\xcf\xfe\x58\x00\x33\xad\x00\x80\x6b\xec\xb8\x61\x27\x58\x6f\x72\xac\x0b\x63\xfe\x9f\x07\x62\x63\x43\x82\xea\x60\xa3\xad\x84\x22\x5e\x6d\x1c\xd8\x35\xb3\xbb\xdb\xf9\xb5\x6b\x5b\xdf\xbb\x89\x2e\x3b\x16\xb6\x42\x75\xb9\x1b\xca\xbe\xa3\x88\x9d\xf7\xcf\x64\xdd\x8f\x4b\xab\x8a\x82\x50\x3d\xce\x2f\xb8\x8a\x74\x3a\x6d\x05\xc2\x2d\x75\x62\xdb\x80\x23\xcb\xc1\xf7\x67\x82\x37\xef\xe8\x22\xf7\xa1\x9e\x2e\x8f\x11\x8d\x01\x53\x43\x11\x69\xf6\x90\x14\x26\xc7\xcd\xe7\x7a\x5a\xcd\x4f\xd1\x4f\xea\xef\xf4\x15\x7e\x29\x70\xf9\x1e\x00\x00\xff\xff\xc2\x8d\x39\x01\x34\x02\x00\x00")

func migrations007_move_forfeit_requests_to_paymentSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations007_move_forfeit_requests_to_paymentSql,
		"migrations/007_move_forfeit_requests_to_payment.sql",
	)
}

func migrations007_move_forfeit_requests_to_paymentSql() (*asset, error) {
	bytes, err := migrations007_move_forfeit_requests_to_paymentSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/007_move_forfeit_requests_to_payment.sql", size: 564, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations008_use_random_salt_in_txSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe5\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xc8\x2c\x2e\xc9\x2f\xaa\x8c\x2f\x29\x4a\xcc\x2b\x4e\x4c\x2e\xc9\xcc\xcf\x2b\x56\x08\x72\xf5\x73\xf4\x75\x55\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x4c\x4e\xce\x2f\xcd\x2b\x89\x2f\x4e\x2d\x2c\x4d\xcd\x4b\x4e\x55\x08\xf1\x57\x28\x4e\xcc\x29\xb1\xe6\xe5\xe2\xe5\x42\x36\xdb\x25\xbf\x3c\x0f\x24\x46\xa2\xf9\x20\xb3\x40\x66\xa2\xdb\x63\xcd\xcb\x05\x08\x00\x00\xff\xff\xce\x55\x14\xb5\xba\x00\x00\x00")

func migrations008_use_random_salt_in_txSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations008_use_random_salt_in_txSql,
		"migrations/008_use_random_salt_in_tx.sql",
	)
}

func migrations008_use_random_salt_in_txSql() (*asset, error) {
	bytes, err := migrations008_use_random_salt_in_txSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/008_use_random_salt_in_tx.sql", size: 186, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations009_use_long_assetsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe5\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xc8\x2c\x2e\xc9\x2f\xaa\x8c\x4f\x4a\xcc\x49\xcc\x4b\x4e\x2d\x56\x80\x48\x3a\xfb\xfb\x84\xfa\xfa\x29\x24\x16\x17\xa7\x96\x28\x84\x44\x06\xb8\x2a\x24\x67\x24\x16\x25\x26\x97\xa4\x16\x29\x94\x25\x16\x55\x66\xe6\xa5\x6b\x18\x9a\x69\x5a\xf3\x72\xf1\x72\x21\x1b\xef\x92\x5f\x9e\x47\x35\x0b\xcc\x34\x15\x42\x83\x3d\xfd\xdc\x15\x8a\x4b\x93\x8a\x4b\x8a\x34\xc0\x8a\x75\x14\x0c\x75\x14\xcc\x34\xad\x01\x01\x00\x00\xff\xff\xe4\x8d\xd8\x28\xd5\x00\x00\x00")

func migrations009_use_long_assetsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations009_use_long_assetsSql,
		"migrations/009_use_long_assets.sql",
	)
}

func migrations009_use_long_assetsSql() (*asset, error) {
	bytes, err := migrations009_use_long_assetsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/009_use_long_assets.sql", size: 213, mode: os.FileMode(438), modTime: time.Unix(1511895194, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations014_balance_historySql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xb1\x4e\xf4\x40\x0c\x84\xeb\x8b\x94\x77\x98\x32\xd1\x7f\xe9\x7e\xd1\x5c\x4b\x45\x01\x05\xa2\x8e\x9c\xc4\x97\x58\x6c\xbc\x2b\xaf\xc3\x29\x3c\x3d\x5a\x10\x70\x50\x50\xce\x78\x3c\xb6\xbe\xae\xc3\xbf\x55\x66\x23\x67\x3c\xa5\xba\xaa\xab\xd1\xb8\x08\xa7\x21\x30\x16\xc9\x1e\x6d\xef\x07\x0a\xa4\x23\xf7\x5b\x9a\xc8\x39\xa3\xa9\xab\x83\x4c\x18\x64\xce\x6c\x42\x01\xc9\x64\x25\xdb\xf1\xcc\xfb\xb1\xae\x0e\x9f\x79\x99\xf0\x42\x36\x2e\x64\xcd\xcd\xff\x16\xc6\x67\x36\xd6\x91\xf3\xef\xe6\x8c\xe6\x7b\xa7\x85\x46\x87\x6e\x21\x94\x2e\x5a\xe3\xa6\x5e\x6e\x89\xfa\x8f\xc9\xc7\x37\x53\x4f\x0e\x97\x95\xb3\xd3\x9a\x70\x11\x5f\xde\x25\x5e\xa3\xf2\x57\x1e\x5d\x77\xe5\x4a\x86\xa8\xb3\xba\x44\xa5\x70\xc4\x39\x1a\x98\xb2\xb0\xe1\xee\xf1\xe1\x1e\x89\x2c\x8b\xce\x75\xd5\x9e\x0a\x93\x6b\x4a\xb7\xf1\xa2\xc5\x9b\x2c\xa6\xbf\x29\x9d\xde\x02\x00\x00\xff\xff\xb6\xb6\x0a\xcb\x5d\x01\x00\x00")

func migrations014_balance_historySqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations014_balance_historySql,
		"migrations/014_balance_history.sql",
	)
}

func migrations014_balance_historySql() (*asset, error) {
	bytes, err := migrations014_balance_historySqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/014_balance_history.sql", size: 349, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations015_account_typeSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xcd\x31\x0a\xc2\x40\x10\x85\xe1\x3e\x90\x3b\xbc\x5e\x02\xf6\x69\xbd\x82\x75\x18\x77\x47\x5d\x98\xcc\x2c\x9b\xb7\x48\x6e\x2f\x82\x85\x85\x90\xf6\x2f\xfe\x6f\x9a\x70\x5a\xcb\xa3\x09\x15\xd7\x3a\x0e\xe3\x20\x46\x6d\xa0\xdc\x4c\xf1\x2c\x1b\xa3\xed\x8b\xa4\x14\xdd\xb9\x41\x72\x46\x0a\xeb\xab\xe3\xdb\x16\xee\x55\x51\x9c\xf0\x20\xbc\x9b\x21\xeb\x5d\xba\x11\xe7\xf9\xf3\xfb\x15\x2e\xf1\xf2\x43\x23\xb7\xa8\xff\x90\xf9\x1d\x00\x00\xff\xff\x63\x5e\x52\x1f\xac\x00\x00\x00")

func migrations015_account_typeSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations015_account_typeSql,
		"migrations/015_account_type.sql",
	)
}

func migrations015_account_typeSql() (*asset, error) {
	bytes, err := migrations015_account_typeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/015_account_type.sql", size: 172, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations016_balance_updates_indexSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe5\xe2\xe5\x72\x0e\x72\x75\x0c\x71\x55\xf0\xf4\x73\x71\x8d\x50\xc8\x48\x2a\x8d\x4f\x4a\xcc\x49\xcc\x4b\x4e\x8d\xcf\x4c\x89\xcf\xcc\x4b\x49\xad\x50\xf0\xf7\x53\xc8\xc8\x2c\x2e\xc9\x2f\xaa\x84\xcb\x95\x16\xa4\x24\x96\xa4\x16\x2b\x84\x06\x7b\xfa\xb9\x2b\x78\x38\x06\x7b\x28\x68\x20\xf4\x69\x5a\x83\x0c\x46\xb6\xca\x25\xbf\x3c\x0f\x24\xe6\x12\xe4\x1f\x80\xc7\x2a\x6b\x5e\x2e\x40\x00\x00\x00\xff\xff\x06\x8c\x62\xc1\xa1\x00\x00\x00")

func migrations016_balance_updates_indexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations016_balance_updates_indexSql,
		"migrations/016_balance_updates_index.sql",
	)
}

func migrations016_balance_updates_indexSql() (*asset, error) {
	bytes, err := migrations016_balance_updates_indexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/016_balance_updates_index.sql", size: 161, mode: os.FileMode(438), modTime: time.Unix(1511894109, 0)}
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
	"migrations/0010_trades.sql": migrations0010_tradesSql,
	"migrations/0011_history_price.sql": migrations0011_history_priceSql,
	"migrations/0012_update_pending_tx_op_key.sql": migrations0012_update_pending_tx_op_keySql,
	"migrations/0013_signer_name.sql": migrations0013_signer_nameSql,
	"migrations/001_squashed.sql": migrations001_squashedSql,
	"migrations/003_update_forfeit_request_created_at.sql": migrations003_update_forfeit_request_created_atSql,
	"migrations/004_change_asset_length.sql": migrations004_change_asset_lengthSql,
	"migrations/005_change_exchange_kyc.sql": migrations005_change_exchange_kycSql,
	"migrations/006_pending_submitter.sql": migrations006_pending_submitterSql,
	"migrations/007_move_forfeit_requests_to_payment.sql": migrations007_move_forfeit_requests_to_paymentSql,
	"migrations/008_use_random_salt_in_tx.sql": migrations008_use_random_salt_in_txSql,
	"migrations/009_use_long_assets.sql": migrations009_use_long_assetsSql,
	"migrations/014_balance_history.sql": migrations014_balance_historySql,
	"migrations/015_account_type.sql": migrations015_account_typeSql,
	"migrations/016_balance_updates_index.sql": migrations016_balance_updates_indexSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"0010_trades.sql": &bintree{migrations0010_tradesSql, map[string]*bintree{}},
		"0011_history_price.sql": &bintree{migrations0011_history_priceSql, map[string]*bintree{}},
		"0012_update_pending_tx_op_key.sql": &bintree{migrations0012_update_pending_tx_op_keySql, map[string]*bintree{}},
		"0013_signer_name.sql": &bintree{migrations0013_signer_nameSql, map[string]*bintree{}},
		"001_squashed.sql": &bintree{migrations001_squashedSql, map[string]*bintree{}},
		"003_update_forfeit_request_created_at.sql": &bintree{migrations003_update_forfeit_request_created_atSql, map[string]*bintree{}},
		"004_change_asset_length.sql": &bintree{migrations004_change_asset_lengthSql, map[string]*bintree{}},
		"005_change_exchange_kyc.sql": &bintree{migrations005_change_exchange_kycSql, map[string]*bintree{}},
		"006_pending_submitter.sql": &bintree{migrations006_pending_submitterSql, map[string]*bintree{}},
		"007_move_forfeit_requests_to_payment.sql": &bintree{migrations007_move_forfeit_requests_to_paymentSql, map[string]*bintree{}},
		"008_use_random_salt_in_tx.sql": &bintree{migrations008_use_random_salt_in_txSql, map[string]*bintree{}},
		"009_use_long_assets.sql": &bintree{migrations009_use_long_assetsSql, map[string]*bintree{}},
		"014_balance_history.sql": &bintree{migrations014_balance_historySql, map[string]*bintree{}},
		"015_account_type.sql": &bintree{migrations015_account_typeSql, map[string]*bintree{}},
		"016_balance_updates_index.sql": &bintree{migrations016_balance_updates_indexSql, map[string]*bintree{}},
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

