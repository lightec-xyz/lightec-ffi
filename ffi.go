package lightec_ffi

/*
#cgo CFLAGS: -I./rust/target/release
#cgo LDFLAGS: -L ./rust/target/release -llightecffi
#include <stdlib.h>
#include "./lightecffi.h"
*/
import "C"

func (c *GenProof) GenProof(a, b int) (int, error) {
	ca := C.int(a)
	cb := C.int(b)
	result := C.gen_proof(ca, cb)
	r := int(result)
	return r, nil
}
