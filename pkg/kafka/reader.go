package kafkax

import (
	"context"
	"errors"

	"github.com/segmentio/kafka-go"
)

const (
	// For pull message
	KindReader uint = 0
	// For push message
	KindWriter = 1
)

type Reader struct {
	reader *kafka.Reader
	readOps
}

type readOps interface {
	CreateTopics()
	StartPull()
	PrintTopics(broker string) error
}

func MakeReader(topic string, r *kafka.Reader) error {
	if topic != r.Config().Topic {
		return errors.New("Topic not the same")
	}

	return nil
}

func Run(r *kafka.Reader, cb func(kafka.Message) error) error {
	for {
		m, e := r.ReadMessage(context.Background())
		if e != nil {
			l.Errorln("error occur when reader pull a message")
		}
		if e = cb(m); e != nil {
			return e
		}
	}
}
