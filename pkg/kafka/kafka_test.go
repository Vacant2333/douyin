package kafkax

import (
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
)

func TestWithNoTopicInstace(t *testing.T) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Topic:          "hello",
		read
	})

}
