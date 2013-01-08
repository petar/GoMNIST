package GoMNIST

import (
	"path"
)

// Set represents a data set of image-label pairs held in memory
type Set struct {
	NRow   int
	NCol   int
	Images []Image
	Labels []Label
}


// ReadSet reads a set from the images file iname and the corresponding labels file lname
func ReadSet(iname, lname string) (set *Set, err error) {
	set = &Set{}
	if set.NRow, set.NCol, set.Images, err = ReadImageFile(iname); err != nil {
		return nil, err
	}
	if set.Labels, err = ReadLabelFile(lname); err != nil {
		return nil, err
	}
	return
}

// Count returns the number of points available in the data set
func (s *Set) Count() int {
	return len(s.Images)
}

// Get returns the i-th image and its corresponding label
func (s *Set) Get(i int) (Image, Label) {
	return s.Images[i], s.Labels[i]
}

// Sweep is an iterator over the points in a data set
type Sweep struct {
	set *Set
	i   int
}

// Next returns the next image and its label in the data set
// If the end is reached, present is set to false
func (sw *Sweep) Next() (image Image, label Label, present bool) {
	if sw.i >= len(sw.set.Images) {
		return nil, 0, false
	}
	return sw.set.Images[sw.i], sw.set.Labels[sw.i], true
}

// Sweep creates a new sweep iterator over the data set
func (s *Set) Sweep() *Sweep {
	return &Sweep{set: s}
}

// Load reads both the training and the testing MNIST data sets, given
// a local directory dir, containing the MNIST distribution files.
func Load(dir string) (train, test *Set, err error) {
	if train, err = ReadSet(path.Join(dir, "train-images-idx3-ubyte.gz"), path.Join(dir, "train-labels-idx1-ubyte.gz")); err != nil {
		return nil, nil, err
	}
	if test, err = ReadSet(path.Join(dir, "t10k-images-idx3-ubyte.gz"), path.Join(dir, "t10k-labels-idx1-ubyte.gz")); err != nil {
		return nil, nil, err
	}
	return
}
