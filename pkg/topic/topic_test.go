package topic

import (
	"os"
	"sort"
	"testing"
)

func TestNewTopic(t *testing.T) {
	path := "./topic"
	_, err := NewTopic(path)
	if err != nil {
		t.Errorf("new topic error: %v", err)
		return
	}

	defer os.Remove(path)
}

func TestTopicCreate(t *testing.T) {
	path := "./topic"
	top, err := NewTopic(path)
	if err != nil {
		t.Errorf("new topic error: %v", err)
		return
	}
	defer os.Remove(path)
	defer top.Close()

	top.Create("test")
	if _, ok := top.topicMap["test"]; !ok {
		t.Error("create a test topic error")
	}
}

func TestTopicDelete(t *testing.T) {
	path := "./topic"
	top, err := NewTopic(path)
	if err != nil {
		t.Errorf("new topic error: %v", err)
		return
	}
	defer os.Remove(path)
	defer top.Close()

	top.Create("test")
	if _, ok := top.topicMap["test"]; !ok {
		t.Error("create a test topic error")
	}

	top.Delete("test")
	if _, ok := top.topicMap["test"]; ok {
		t.Error("delete a test topic error")
	}
}

func TestTopicGetAll(t *testing.T) {
	path := "./topic"
	top, err := NewTopic(path)
	if err != nil {
		t.Errorf("new topic error: %v", err)
		return
	}
	defer os.Remove(path)
	defer top.Close()

	topList := []string{"one", "two", "three", "four", "five"}
	sort.Strings(topList)

	for _, t := range topList {
		top.Create(t)
	}

	a := top.GetAll()
	sort.Strings(a)
	for index := range a {
		if topList[index] != a[index] {
			t.Errorf("topic GetAll error, top: %s, a: %s", topList[index], a[index])
			break
		}
	}
}
