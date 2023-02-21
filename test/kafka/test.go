package main

import (
	constantx "douyin/pkg/constant"
	kafkax "douyin/pkg/kafka"
	"douyin/pkg/kafka/msg"
	"douyin/pkg/logger"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto" // Must use google ver.
)

var l *zap.SugaredLogger

func init() {
	l = logger.NewLogger()
}

func main() {
	w := &kafkax.Writer{
		Writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  append([]string{}, constantx.KAFKA_TestBroker),
			Topic:    constantx.KAFKA_TestTopic,
			Balancer: &kafka.LeastBytes{},
		}),
	}
	w.Run()

	r := &kafkax.Reader{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:   append([]string{}, constantx.KAFKA_TestBroker),
			Topic:     constantx.KAFKA_TestTopic,
			Partition: constantx.KAFKA_PartitionNum,
			MinBytes:  10e3,
			MaxBytes:  10e6,
			//ReadBatchTimeout: 500 * time.Millisecond, // try to read's timeout def: 10s
			//ReadBackoffMin: ,  // Min interval def: 100ms
			//ReadBackoffMax: 500 * time.Millisecond, // Max interval def: 1s

		}),
		IsFetchMode: false, // do need commit?
		PullCb:      pullCb,
	}

	r.Run()

	for i := 0; i < 100; i++ {
		m := msg.Msg{
			Id:      int64(i),
			Content: "hello",
		}
		w.MsgChan <- m.ProtoReflect().Interface()
	}

}

func pullCb(m *kafka.Message) error {
	msg := &msg.Msg{}
	err := proto.Unmarshal(m.Value, msg)
	if err != nil {
		l.Error("[kafka-reader::readcb] Error when unmarshal a message")
	}

	l.Debugf("Id: %v, Content: %v", msg.Id, msg.Content)

	return nil
}
