package wrapper

import (
	"io/ioutil"
)

type Reader interface {
	Read(path string) ([]byte, error)
}

type FileReader struct {
}

func (fr FileReader) Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
