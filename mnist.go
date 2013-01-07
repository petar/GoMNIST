package mnist

import (
	"compress/gzip"
	"encoding/binary"
	"io"
	"os"
)

type Label uint8

func ReadLabelFile(name string) ([]Label, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	z, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	return readLabelFile(z)
}

func readLabelFile(r io.Reader) ([]Label, error) {
	var (
		magic int32
		n     int32
	)
	if err := binary.Read(r, binary.BigEndian, &magic); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &n); err != nil {
		return nil, err
	}
	println("Reading", n, "records")
	ll := make([]Label, n)
	for i := int32(0); i < n; i++ {
		var l Label
		if err := binary.Read(r, binary.BigEndian, &l); err != nil {
			return nil, err
		}
		ll[int(i)] = l
	}
	return ll, nil
}
