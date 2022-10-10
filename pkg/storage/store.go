package storage

import (
	"bufio"
	"io"
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

	body := make([]byte, dataLenWidth)
	enc.PutUint64(body, uint64(len(p)))
	body = append(body, p...)

	if _, err = s.buf.Write(body); err != nil {
		return
	}

	s.size += uint64(len(body))
	return
}

func (s *store) read(offset uint64) ([]byte, error) {
	if s.size < offset+dataLenWidth {
		return nil, io.EOF
	}

	if err := s.buf.Flush(); err != nil {
		return nil, err
	}

	lenBytes := make([]byte, dataLenWidth)
	if _, err := s.file.ReadAt(lenBytes, int64(offset)); err != nil {
		return nil, err
	}

	if s.size < (offset + dataLenWidth + enc.Uint64(lenBytes)) {
		return nil, io.EOF
	}

	data := make([]byte, enc.Uint64(lenBytes))
	if _, err := s.file.ReadAt(data, int64(offset+dataLenWidth)); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *store) close() error {
	if err := s.buf.Flush(); err != nil {
		return err
	}

	if err := s.file.Close(); err != nil {
		return err
	}

	return nil
}
