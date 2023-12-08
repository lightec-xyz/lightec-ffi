package main

import lightec_ffi "github.com/lightec-xyz/lightec-ffi"

func main() {
	proof := lightec_ffi.GenProof{}
	sum, err := proof.GenProof(1, 2)
	if err != nil {
		panic(err)
	}
	println(sum)
}
