package kafkax

import "github.com/segmentio/kafka-go"

type Manager struct {
	reader *kafka.Reader
	writer *kafka.Writer
	operator
}

func New() *Manager {
	return nil
}

func (m *Manager) GetWriter() (*kafka.Writer, error) {
	m.reader = kafka.NewReader(kafka.ReaderConfig{})
	return nil, nil
}

func (m *Manager) GetReader() (*kafka.Reader, error) {
	m.writer = kafka.NewWriter(kafka.WriterConfig{})
	return nil, nil
}
