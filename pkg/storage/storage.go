package storage

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/wan-maoyuan/kafka-go/pkg/utils"
)

const (
	messageCacheSize        = 1024     // 消息缓冲区的大小
	indexSuffix      string = ".index" // 二进制文件索引文件后缀
	storeSuffix      string = ".store" // 二进制文件后缀
)

type Storage struct {
	topicMap map[string]struct{}   // 所有的topic 的 map
	sm       map[string]*topicInfo // 以 topic 为 key 的map，存储对象和索引对象
}

type topicInfo struct {
	msgChan  chan []byte
	maxCount int64
	fileName string
	idx      *index
	sto      *store
}

func NewSorage() (*Storage, error) {
	topics, err := getTopicsFromDataDir()
	if err != nil {
		return nil, err
	}

	topicMap, err := getTopicBinaryInfo(topics)
	if err != nil {
		return nil, err
	}

	return &Storage{
		topicMap: topics,
		sm:       topicMap,
	}, nil
}

// 根据主题保存消息数据
func (s *Storage) SaveMessage(topicName string, message []byte) error {
	_, ok := s.sm[topicName]
	if !ok {
		if err := s.createTopicFolder(topicName); err != nil {
			return err
		}
	}

	return nil
}

// 根据主题和消息偏移量获取消息
func (s *Storage) GetMessage(topicName string, offset uint64) ([]byte, error) {

	return nil, nil
}

// 某个主题删除之后，对应的消息二进制数据也需要删除
func (s *Storage) DeleteDataByTopic(topicName string) {}

func (s *Storage) createTopicFolder(topicName string) error {
	s.topicMap[topicName] = struct{}{}

	return nil
}

func (s *Storage) saveMessage2File() {}

// 读取数据文件夹，将所有的主题读取出来
func getTopicsFromDataDir() (map[string]struct{}, error) {
	dirList, err := os.ReadDir(utils.C.Log.FilePath)
	if err != nil {
		return nil, err
	}

	topics := make(map[string]struct{}, len(dirList))
	for _, file := range dirList {
		if file.IsDir() {
			topics[file.Name()] = struct{}{}
		}
	}

	return topics, nil
}

// 根据 topic 生成 topicInfo 结构体
func getTopicBinaryInfo(topics map[string]struct{}) (map[string]*topicInfo, error) {
	topicInfos := make(map[string]*topicInfo, 0)

	for topic := range topics {
		topicDir := filepath.Join(utils.C.Log.FilePath, topic)
		topicFiles, err := os.ReadDir(topicDir)
		if err != nil {
			continue
		}

		var info = topicInfo{
			msgChan:  make(chan []byte, messageCacheSize),
			maxCount: 0,
			fileName: strings.Repeat("0", 16),
		}

		for _, topicFile := range topicFiles {
			if strings.HasSuffix(topicFile.Name(), indexSuffix) {
				length := len(topicFile.Name())
				name := topicFile.Name()[:length-6]

				messageIndex, err := strconv.ParseInt(name, 10, 64)
				if err != nil {
					logrus.Errorf("convert file name to int error: %v, name: %s", err, name)
					continue
				}

				if info.maxCount < messageIndex {
					info.maxCount = messageIndex
					info.fileName = name
				}
			}
		}

		indexPath := filepath.Join(utils.C.Log.FilePath, topic, info.fileName+indexSuffix)
		i, err := newIndex(indexPath)
		if err != nil {
			logrus.Errorf("create topic: %s index error: %v", topic, err)
			continue
		}
		info.idx = i

		storePath := filepath.Join(utils.C.Log.FilePath, topic, info.fileName+storeSuffix)
		s, err := newStore(storePath)
		if err != nil {
			logrus.Errorf("create topic: %s store error: %v", topic, err)
			continue
		}
		info.sto = s

		topicInfos[topic] = &info
	}

	return topicInfos, nil
}
