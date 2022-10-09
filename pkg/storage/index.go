package storage

import (
	"encoding/binary"
	"io"
	"os"
)

const (
	indexWidth  uint64 = 4                        // 下标的字节数
	offsetWidth uint64 = 8                        // 数据文件偏移数
	rowWidth    uint64 = indexWidth + offsetWidth // 一条索引记录的字节数
)

var (
	enc = binary.BigEndian
)

// 数据索引
type index struct {
	file *os.File
	size uint64
}

func newIndex(f *os.File) (*index, error) {
	fileInfo, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}

	return &index{
		file: f,
		size: uint64(fileInfo.Size()),
	}, nil
}

func (i *index) write(index uint32, offset uint64) error {
	if err := binary.Write(i.file, enc, index); err != nil {
		return err
	}

	if err := binary.Write(i.file, enc, offset); err != nil {
		return err
	}

	i.size += rowWidth
	return nil
}

func (i *index) read(index uint32) (offset uint64, err error) {
	if i.size == 0 || i.size < uint64(index+1)*rowWidth {
		err = io.EOF
		return
	}

	b := make([]byte, offsetWidth)
	if _, err = i.file.ReadAt(b, int64(uint64(index)*rowWidth+indexWidth)); err != nil {
		return
	}

	offset = enc.Uint64(b)
	return
}

func (i *index) readLast() (index uint32, err error) {
	if i.size == 0 {
		err = io.EOF
		return
	}

	index = uint32((i.size / rowWidth) - 1)
	return
}

func (i *index) close() error {
	if err := i.file.Sync(); err != nil {
		return err
	}

	return nil
}
