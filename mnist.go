package mnist

import (
	"compress/gzip"
	"encoding/binary"
	"io"
	"os"
)

const (
	imageMagic = 0x00000803
	labelMagic = 0x00000801
)

type Image []byte

func ReadImageFile(name string) (rows, cols int, imgs []Image, err error) {
	?
}

func readImageFile(r io.Reader) (rows, cols int, imgs []Image, err error) {
	var (
		magic int32
		n     int32
		nrows int32
		ncols int32
	)
	if err = binary.Read(r, binary.BigEndian, &magic); err != nil {
		return nil, err
	}
	if magic != imageMagic {
		return nil, os.ErrInvalid
	}
	if err = binary.Read(r, binary.BigEndian, &n); err != nil {
		return nil, err
	}
	if err = binary.Read(r, binary.BigEndian, &nrows); err != nil {
		return nil, err
	}
	if err = binary.Read(r, binary.BigEndian, &ncols); err != nil {
		return nil, err
	}
	imgs = make()
	for i := int32(0); i < n; i++ {
		…
	}
	return imgs, nil
}

type Label uint8

// ReadLabelFile opens the named label file (training or test), parses it and
// returns all labels in order.
func ReadLabelFile(name string) (labels []Label, err error) {
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

func readLabelFile(r io.Reader) (labels []Label, err error) {
	var (
		magic int32
		n     int32
	)
	if err = binary.Read(r, binary.BigEndian, &magic); err != nil {
		return nil, err
	}
	if magic != labelMagic {
		return nil, os.ErrInvalid
	}
	if err = binary.Read(r, binary.BigEndian, &n); err != nil {
		return nil, err
	}
	println("Reading", n, "records")
	labels = make([]Label, n)
	for i := int32(0); i < n; i++ {
		var l Label
		if err := binary.Read(r, binary.BigEndian, &l); err != nil {
			return nil, err
		}
		labels[int(i)] = l
	}
	return labels, nil
}
