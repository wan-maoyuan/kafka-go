package storage

import (
	"bufio"
	"os"
)

const (
	dataLenWidth uint64 = 8
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
	offset = s.size

	lenBytes := make([]byte, dataLenWidth)
	enc.PutUint64(lenBytes, uint64(len(p)))
	lenBytes = append(lenBytes, p...)

	if _, err = s.buf.Write(lenBytes); err != nil {
		return
	}

	return
}

func (s *store) read(offset uint64) ([]byte, error) {

	return nil, nil
}

func (s *store) close() error {

	return nil
}
