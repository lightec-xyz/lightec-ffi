package lightec_ffi

import "testing"

func TestAdd(t *testing.T) {
	proof := GenProof{}
	sum, err := proof.add(1, 2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sum)
}
