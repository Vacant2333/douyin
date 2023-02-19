package kafkax

import (
	constantx "douyin/pkg/constant"
	"douyin/pkg/logger"
	"fmt"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var l *zap.SugaredLogger

func init() {
	l = logger.NewLogger()
}

func printTopics(broker string) error {

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

func createTopics(brokers, topics []string) error {

	for _, broker := range brokers {

		// Create topic 'test_proto' explicitly, when KAFKA_AUTO_CREATE_TOPICS_ENABLE='false'
		conn, err := kafka.Dial("tcp", broker)
		if err != nil {
			return err
		}
		defer conn.Close()

		// Get controller
		controller, err := conn.Controller()
		fmt.Println(controller.Host)
		fmt.Println(controller.Port)
		if err != nil {
			return err
		}

		// Dial the header
		var ctrlConn *kafka.Conn
		ctrlConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
		if err != nil {
			return err
		}
		defer ctrlConn.Close()

		// Make topics create configs
		topicConfs := []kafka.TopicConfig{}
		for _, topic := range topics {
			toppicconf := &kafka.TopicConfig{
				Topic:             topic,
				NumPartitions:     constantx.KAFKA_PartitionNum,
				ReplicationFactor: constantx.KAFKA_ReplicationNum,
			}
			topicConfs = append(topicConfs, *toppicconf)
		}

		err = ctrlConn.CreateTopics(topicConfs...)
		if err != nil {
			return err
		}

		printTopics(broker)
	}

	return nil
}
