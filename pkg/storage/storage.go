package storage

import (
	"fmt"
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
	msgChan      chan []byte // 消息发送队列
	currentCount uint32      // 数据 index 最大值
	fileName     string      // 存储的文件名，出去后缀
	idx          *index      // 二进制文件
	sto          *store      // 二进制文件索引文件
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
	if _, ok := s.sm[topicName]; !ok {
		if err := s.initTopicFolder(topicName); err != nil {
			return err
		}
	}

	if err := s.saveMessage2File(topicName, message); err != nil {
		return err
	}

	return nil
}

// 根据主题和消息偏移量获取消息
func (s *Storage) GetMessage(topicName string, offset uint64) ([]byte, error) {

	return nil, nil
}

// 某个主题删除之后，对应的消息二进制数据也需要删除
func (s *Storage) DeleteDataByTopic(topicName string) {}

func (s *Storage) initTopicFolder(topicName string) error {
	s.topicMap[topicName] = struct{}{}

	topicDir := filepath.Join(utils.C.Log.FilePath, topicName)
	os.RemoveAll(topicDir)
	if err := os.Mkdir(topicDir, 0777); err != nil {
		return err
	}

	fileName := fmt.Sprintf("%016d", 0)
	idx, err := newIndex(filepath.Join(topicDir, fileName+".index"))
	if err != nil {
		return fmt.Errorf("topic: %s create index file error: %v", topicName, err)
	}

	sto, err := newStore(filepath.Join(topicDir, fileName+".store"))
	if err != nil {
		idx.close()
		return fmt.Errorf("create topic: %s store file error: %v", topicName, err)
	}

	s.sm[topicName] = &topicInfo{
		msgChan:      make(chan []byte, messageCacheSize),
		currentCount: 0,
		fileName:     fileName,
		idx:          idx,
		sto:          sto,
	}

	return nil
}

func (s *Storage) saveMessage2File(topicName string, message []byte) error {
	info := s.sm[topicName]
	offset, err := info.sto.write(message)
	if err != nil {
		return fmt.Errorf("Storage write message to store error: %v", err)
	}

	if err := info.idx.write(info.currentCount, offset); err != nil {
		return fmt.Errorf("Storage write message to index error: %v", err)
	}

	info.currentCount += 1
	return nil
}

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
			msgChan:      make(chan []byte, messageCacheSize),
			currentCount: 0,
			fileName:     strings.Repeat("0", 16),
		}

		for _, topicFile := range topicFiles {
			if strings.HasSuffix(topicFile.Name(), indexSuffix) {
				length := len(topicFile.Name())
				name := topicFile.Name()[:length-6]

				messageIndex, err := strconv.ParseUint(name, 10, 32)
				if err != nil {
					logrus.Errorf("convert file name to int error: %v, name: %s", err, name)
					continue
				}

				if info.currentCount < uint32(messageIndex) {
					info.currentCount = uint32(messageIndex)
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
		if num, err := info.idx.readLast(); err != nil {
			info.currentCount = num + 1
		}

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
