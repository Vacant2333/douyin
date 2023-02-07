package kafkax

import (
	"github.com/segmentio/kafka-go"
)

// TODO: define and implent
type operator interface {
	writerOnce(topic, broker string, msg interface{})
	readOnce()
	CreateTopics()
	Pull()
	Push()
	PrintTopics(broker string) error
}

func (m *Manager) PrintTopics(broker string) error {

	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		return err
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return err
	}

	mp := map[string]struct{}{}

	for _, p := range partitions {
		mp[p.Topic] = struct{}{}
	}

	for k := range mp {
		println(k)
	}

	return nil
}
