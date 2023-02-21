package kafkax

import (
	"context"
	"errors"
	"strings"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

type Writer struct {
	Writer *kafka.Writer
	// Callback funtion when get a msg
	Postcb  func(*Writer, proto.Message) error
	MsgChan chan proto.Message
}

func (w *Writer) Run() {

	// FIXME: maybe need trim port suffix
	brokers := strings.Split(w.Writer.Addr.String(), ",")
	CreateTopics(brokers, append([]string{}, w.Writer.Topic))

	if w.MsgChan == nil {
		w.MsgChan = make(chan proto.Message)
	}

	go func() {
		defer w.Writer.Close()
		for {
			var e error
			m := <-w.MsgChan
			if w.Postcb == nil {
				e = templateCb(w, m)

			} else {
				e = w.Postcb(w, m)
			}
			if e != nil {
				l.Errorf("error when write %v", e)
			}

		}
	}()

}

func templateCb(w *Writer, m proto.Message) error {
	pbbytes, err := proto.Marshal(m)
	if err != nil {
		return errors.New("cannot marshal meesage to protobuf binary bytes")
	}
	l.Debugln("[kafka-writer] Marshal from message struct to pb bytes success")

	err = w.Writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: pbbytes,
		},
	)
	if err != nil {
		return err
	}

	l.Debugf("[kafka-writer] Post one message to kafka <%v>", w.Writer.Topic)

	return nil
}
