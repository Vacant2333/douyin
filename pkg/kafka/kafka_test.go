package kafkax

import (
	"strings"
	"testing"

	"github.com/segmentio/kafka-go"
)

func TestInitSequence(t *testing.T) {
	w := kafka.NewWriter(
		kafka.WriterConfig{
			Brokers: append([]string{}, "192.168.1.1", "172.0.0.1"),
			Topic:   "Comment",
		},
	)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: append([]string{}, "192.168.1.1", "172.0.0.1"),
		Topic:   "Comment",
	})

	//t.Logf(w.Addr.String())

	ss1 := strings.Split(w.Addr.String(), ",")
	ss2 := r.Config().Brokers

	for i := 0; i < len(ss1); i++ {
		tmp := strings.TrimSuffix(ss1[i], ":9092")
		t.Log(tmp, "--", ss2[i])
		if !strings.EqualFold(tmp, ss2[i]) {
			t.Error("error when get message from reder/writer ")
		}
	}
	if w.Topic != r.Config().Topic {
		t.Log(w.Topic, "---", r.Config().Topic)
		t.Log(w.Topic, "---", r.Stats().Topic)
		t.Error("error when get message from reder/writer ")
	}
}

func TestLog(t *testing.T) {
	t.Logf("str1: %v" "str2: %v", 1, 2)
}
