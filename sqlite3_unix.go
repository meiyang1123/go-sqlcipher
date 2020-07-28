// +build !windows

package sqlite3

/*
#cgo darwin LDFLAGS: -L/usr/local/opt/openssl@1.1/lib
#cgo LDFLAGS: -lcrypto -ldl -lz
#cgo linux LDFLAGS: -Wl,--allow-multiple-definition
*/
import "C"
