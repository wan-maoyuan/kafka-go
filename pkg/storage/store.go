package storage

import (
	"bufio"
	"os"
)

type store struct {
	file *os.File
	buf  bufio.Writer
	size uint64
}

func newStore(f *os.File) (*store, error) {
	fileInfo, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}

	return &store{
		file: f,
		buf:  *bufio.NewWriter(f),
		size: uint64(fileInfo.Size()),
	}, nil
}

func (s *store) write(p []byte) (offset uint64, err error) {

	return 0, nil
}

func (s *store) read(offset uint64) ([]byte, error) {

	return nil, nil
}
