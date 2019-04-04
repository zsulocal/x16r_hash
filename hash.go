package x16r_hash

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lx16r_hash
// #include "x16r.h"
import "C"

import (
	"unsafe"
)

func GetPowHash(hash []byte) []byte {
	powhash := make([]byte, 32)
	C.x16r_hash((*C.char)(unsafe.Pointer(&hash[0])), (*C.char)(unsafe.Pointer(&powhash[0])))
	return powhash
}
