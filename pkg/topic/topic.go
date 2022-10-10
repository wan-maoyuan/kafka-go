package topic

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type topic struct {
	path     string
	topicMap map[string]struct{}
}

func NewTopic(path string) (*topic, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	m := make(map[string]struct{})
	reader := bufio.NewReader(f)
	stop := false
	for !stop {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		if strings.Trim(string(line), " ") != "" {
			m[string(line)] = struct{}{}
		}
		stop = isPrefix
	}

	return &topic{
		path:     f.Name(),
		topicMap: m,
	}, nil
}

func (t *topic) Create(name string) {
	t.topicMap[name] = struct{}{}
}

func (t *topic) Delete(name string) {
	delete(t.topicMap, name)
}

// // 定时将内存中的数据同步到硬盘上
// func (t *topic) CronSync() {
// 	c := cron.New()

// 	c.AddFunc("* 5 * * * *", func() { // 每个5分钟同步一次
// 		topicSync(t.path, t.topicMap)
// 	})

// 	c.Start()
// 	select {}
// }

func (t *topic) Close() {
	topicSync(t.path, t.topicMap)
}

func topicSync(path string, m map[string]struct{}) {
	dir := filepath.Dir(path)
	tmp := filepath.Join(dir, ".topic.tmp")
	tmpFile, err := os.OpenFile(tmp, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		logrus.Errorf("create and open topic tmp file error: %v", err)
		return
	}

	for key := range m {
		tmpFile.WriteString(key + "\n")
	}
	tmpFile.Close()

	if err := os.Remove(path); err != nil {
		return
	}

	os.Rename(tmpFile.Name(), path)
}
