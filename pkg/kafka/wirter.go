package kafkax

import (
	"context"
	"errors"
	"strings"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type Writer struct {
	writer *kafka.Writer
	// Callback funtion when get a msg
	postCb  func(*kafka.Writer, *proto.Message) error
	msgChan chan proto.Message
}

func NewWriter(brokers []string, topic string) *Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &Writer{
		writer: writer,
	}
}

func (w *Writer) WithPostCb(postCb func(*kafka.Writer, *proto.Message) error) *Writer {
	w.postCb = postCb
	return w
}

func (w *Writer) WithMsgChan(msgChan chan proto.Message) *Writer {
	w.msgChan = msgChan
	return w
}

func (w *Writer) Run() {
	brokers := strings.Split(w.writer.Addr.String(), ",")
	CreateTopics(brokers, append([]string{}, w.writer.Topic))

	if w.msgChan == nil {
		w.msgChan = make(chan proto.Message)
	}

	go func() {
		defer w.writer.Close()
		for {
			var e error
			m := <-w.msgChan
			if w.postCb == nil {
				e = defaultWriteCb(w.writer, &m)
			} else {
				e = w.postCb(w.writer, &m)
			}
			if e != nil {
				l.Errorf("error when write %v", e)
			}
		}
	}()
}

func (w *Writer) GetMsgChan() chan proto.Message {
	return w.msgChan
}

func defaultWriteCb(writer *kafka.Writer, m *proto.Message) error {
	pbbytes, err := proto.Marshal(*m)
	if err != nil {
		return errors.New("cannot marshal message to protobuf binary bytes")
	}
	l.Debugln("[kafka-writer] Marshal from message struct to pb bytes success")

	err = writer.WriteMessages(context.Background(), kafka.Message{Value: pbbytes})
	if err != nil {
		return err
	}

	l.Debugf("[kafka-writer] Post one message to kafka <%v>", writer.Topic)
	return nil
}
