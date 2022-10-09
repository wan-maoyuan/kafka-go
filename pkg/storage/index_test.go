package storage

import (
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	filePath := "./0001.index"
	os.Remove(filePath)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		t.Fatalf("open index file error: %v", err)
	}

	i, err := newIndex(file)
	if err != nil {
		t.Fatalf("new index error: %v", err)
	}
	defer i.close()

	if err := i.write(1, 0); err != nil {
		t.Errorf("index write data error: %v", err)
	}

	os.Remove(filePath)
}

func BenchmarkWrite(b *testing.B) {
	filePath := "./0001.index"
	os.Remove(filePath)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		b.Fatalf("open index file error: %v", err)
	}

	i, err := newIndex(file)
	if err != nil {
		b.Fatalf("new index error: %v", err)
	}
	defer i.close()

	for j := 0; j < b.N; j++ {
		if err := i.write(uint32(j), 0); err != nil {
			b.Errorf("index write data error: %v", err)
		}
	}

	os.Remove(filePath)
}

func TestRead(t *testing.T) {
	filePath := "./0001.index"
	os.Remove(filePath)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		t.Fatalf("open index file error: %v", err)
	}

	i, err := newIndex(file)
	if err != nil {
		t.Fatalf("new index error: %v", err)
	}
	defer i.close()

	if err := i.write(0, 0); err != nil {
		t.Errorf("index write data error: %v", err)
	}

	if err := i.write(1, 12); err != nil {
		t.Errorf("index write data error: %v", err)
	}

	if err := i.write(2, 28); err != nil {
		t.Errorf("index write data error: %v", err)
	}

	if offset, err := i.read(0); err != nil && offset == 0 {
		t.Errorf("index read 0 , offset: %d", offset)
	}

	if offset, err := i.read(1); err != nil && offset == 12 {
		t.Errorf("index read 0 , offset: %d", offset)
	}

	if offset, err := i.read(2); err != nil && offset == 28 {
		t.Errorf("index read 0 , offset: %d", offset)
	}

	os.Remove(filePath)
}

func BenchmarkRead(b *testing.B) {
	filePath := "./0001.index"
	os.Remove(filePath)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		b.Fatalf("open index file error: %v", err)
	}

	i, err := newIndex(file)
	if err != nil {
		b.Fatalf("new index error: %v", err)
	}
	defer i.close()

	if err := i.write(0, 0); err != nil {
		b.Errorf("index write data error: %v", err)
	}

	if err := i.write(1, 12); err != nil {
		b.Errorf("index write data error: %v", err)
	}

	if err := i.write(2, 28); err != nil {
		b.Errorf("index write data error: %v", err)
	}

	for j := 0; j < b.N; j++ {
		if offset, err := i.read(0); err != nil && offset == 0 {
			b.Errorf("index read 0 , offset: %d", offset)
		}

		if offset, err := i.read(1); err != nil && offset == 12 {
			b.Errorf("index read 0 , offset: %d", offset)
		}

		if offset, err := i.read(2); err != nil && offset == 28 {
			b.Errorf("index read 0 , offset: %d", offset)
		}
	}

	os.Remove(filePath)
}
