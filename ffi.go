package lightec_ffi

/*
#cgo LDFLAGS: -L ./rust/target/release -llightecffi
#include <stdlib.h>
#include "./lightecffi.h"
*/
import "C"

func (c *GenProof) add(a, b int) (int, error) {
	ca := C.int(a)
	cb := C.int(b)
	result := C.add_numbers(ca, cb)
	r := int(result)
	return r, nil
}
