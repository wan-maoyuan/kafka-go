package topic

type topic struct{}

func NewTopic() (*topic, error) {
	return nil, nil
}

func (t *topic) Create(name string) error {
	return nil
}

func (t *topic) Delete(name string) {

}

func (t *topic) Close() error {
	return nil
}
