package mnist

import (
	"testing"
)

func TestReadLabelFile(t *testing.T) {
	ll, err := ReadLabelFile("data/t10k-labels-idx1-ubyte.gz")
	if err != nil {
		t.Fatalf("read (%s)", err)
	}
	if len(ll) != 10000 {
		t.Errorf("unexpected count %d", len(ll))
	}
}
