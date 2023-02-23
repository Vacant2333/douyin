package kafkax

import (
	"context"
	constantx "douyin/pkg/constant"
	"time"

	"github.com/segmentio/kafka-go"
)

type Reader struct {
	Reader *kafka.Reader
	// Callback funtion when get a msg
	PullCb      func(m *kafka.Message) error
	err         chan error
	IsFetchMode bool
}

func NewReader(topic string, brokers []string, pullCb func(m *kafka.Message) error, isFetchMode bool) (*Reader, error) {
	return NewGroupReader(topic, brokers, "", pullCb, isFetchMode)
}

func NewGroupReader(topic string, brokers []string, groupID string, pullCb func(m *kafka.Message) error, isFetchMode bool) (*Reader, error) {
	r := &Reader{
		PullCb:      pullCb,
		IsFetchMode: isFetchMode,
		err:         make(chan error),
	}

	config := kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topic,
		MinBytes:  10e3,
		MaxBytes:  10e6,
		MaxWait:   3 * time.Second,
		Partition: constantx.KAFKA_PartitionNum,
		// ReadLagInterval: -1,
	}

	if groupID != "" {
		config.GroupID = groupID
	}

	if isFetchMode {
		//config.ReadBackoffMax = 10 * time.Millisecond
		config.ReadBackoffMin = 10 * time.Millisecond
	}

	r.Reader = kafka.NewReader(config)

	return r, nil
}

func (r *Reader) Run() {
	CreateTopics(r.Reader.Config().Brokers, append([]string{}, r.Reader.Config().Topic))
	if r.PullCb == nil {
		l.DPanic("have not appoint pull callback function in struct")
	}

	if r.IsFetchMode {
		r.fetchRun()
	} else {
		r.run()
	}

	// Listen for errors
	go func() {
		for {
			e := <-r.err
			l.Errorln("erroer occured in Kafka reader: ", e.Error())
		}
	}()
}

/*
* @bref: Fetch a single message then commit ever single time
  - will make performance lost
*/
func (r *Reader) fetchRun() {
	go func() {
		defer r.Reader.Close()

		for {
			m, err := r.Reader.FetchMessage(context.Background())
			if err != nil {
				r.err <- err
				continue
			}
			if err := r.PullCb(&m); err != nil {
				r.err <- err
				continue
			}
			if err := r.Reader.CommitMessages(context.Background(), m); err != nil {
				r.err <- err
			}
		}
	}()
}

func (r *Reader) run() {
	go func() {
		defer r.Reader.Close()
		for {
			m, err := r.Reader.ReadMessage(context.Background())
			if err != nil {
				r.err <- err
				continue
			}
			if err := r.PullCb(&m); err != nil {
				r.err <- err
			}
		}
	}()
}
