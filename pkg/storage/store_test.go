package storage

import (
	"os"
	"testing"
)

func TestStoreWrite(t *testing.T) {
	path := "./test.store"
	os.Remove(path)

	s, err := newStore(path)
	if err != nil {
		t.Errorf("new store error: %v", err)
		return
	}
	defer s.close()

	offset, err := s.write([]byte("123456789"))
	if err != nil {
		t.Errorf("store write binary data error: %v", err)
		return
	}

	data, err := s.read(offset)
	if err != nil {
		t.Errorf("store read binary data by offset error: %v", err)
		return
	}

	if string(data) != "123456789" {
		t.Error("store data not equal get data")
	}

	os.Remove(path)
}

func BenchmarkWrite(b *testing.B) {
	path := "./test.store"
	os.Remove(path)

	s, err := newStore(path)
	if err != nil {
		b.Errorf("new store error: %v", err)
		return
	}
	defer s.close()

	for i := 0; i < b.N; i++ {
		_, err := s.write([]byte("hello world"))
		if err != nil {
			b.Errorf("store write binary data error: %v", err)
			return
		}
	}

	os.Remove(path)
}

func BenchmarkStoreRead(b *testing.B) {
	path := "./test.store"
	os.Remove(path)

	s, err := newStore(path)
	if err != nil {
		b.Errorf("new store error: %v", err)
		return
	}
	defer s.close()

	offset, err := s.write([]byte("123456789"))
	if err != nil {
		b.Errorf("store write binary data error: %v", err)
		return
	}

	for i := 0; i < b.N; i++ {
		data, err := s.read(offset)
		if err != nil {
			b.Errorf("store read binary data by offset error: %v", err)
			return
		}

		if string(data) != "123456789" {
			b.Error("store data not equal get data")
		}
	}

	os.Remove(path)
}
