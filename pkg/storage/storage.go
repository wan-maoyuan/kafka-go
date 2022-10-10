package storage

type Storage struct {
}

// 根据主题保存消息数据
func (s *Storage) SaveMessage(topicName string, message []byte) error {
	return nil
}

// 根据主题和消息偏移量获取消息
func (s *Storage) GetMessage(topicName string, offset uint64) ([]byte, error) {

	return nil, nil
}

// 某个主题删除之后，对应的消息二进制数据也需要删除
func (s *Storage) DeleteDataByTopic(topicName string) {

}
