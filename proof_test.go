package lightec_ffi

import "testing"

func TestGenProof(t *testing.T) {
	proof := GenProof{}
	sum, err := proof.GenProof(1, 2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sum)
}
