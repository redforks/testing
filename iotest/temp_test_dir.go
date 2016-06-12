package iotest

import (
	"io/ioutil"
	"os"

	"github.com/redforks/testing/reset"
)

// Create a temporary directory used for testing, the directory deleted automatically in spork/testing/reset.
type TempTestDir string

// Create a new TempTestDir.
func NewTempTestDir() TempTestDir {
	if dir, err := ioutil.TempDir(``, `tst`); err != nil {
		panic(err)
	} else {
		r := TempTestDir(dir)
		reset.Add(func() {
			r.Close()
		})
		return r
	}
}

// The absolute path of temp directory.
func (ttd TempTestDir) Dir() string {
	return string(ttd)
}

// Close() delete the whole temp directory, including the files and sub directories in
// it. Ignore any IO error. Normally do not need call .Close(), it automatically called in spork/testing/reset.
func (ttd TempTestDir) Close() {
	_ = os.RemoveAll(string(ttd))
}
