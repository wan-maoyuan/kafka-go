package topic

import (
	"os"
	"testing"
)

func TestNewTopic(t *testing.T) {
	path := "./topic"
	_, err := NewTopic(path)
	if err != nil {
		t.Errorf("new topic error: %v", err)
		return
	}

	os.Remove(path)
}

func TestTopicCreate(t *testing.T) {
	path := "./topic"
	top, err := NewTopic(path)
	if err != nil {
		t.Errorf("new topic error: %v", err)
		return
	}
	defer top.Close()

	top.Create("test")
	if _, ok := top.topicMap["test"]; !ok {
		t.Error("create a test topic error")
	}

	os.Remove(path)
}

func TestTopicDelete(t *testing.T) {
	path := "./topic"
	top, err := NewTopic(path)
	if err != nil {
		t.Errorf("new topic error: %v", err)
		return
	}
	defer top.Close()

	top.Create("test")
	if _, ok := top.topicMap["test"]; !ok {
		t.Error("create a test topic error")
	}

	top.Delete("test")
	if _, ok := top.topicMap["test"]; ok {
		t.Error("delete a test topic error")
	}

	os.Remove(path)
}
