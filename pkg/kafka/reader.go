package kafkax

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Reader struct {
	Reader *kafka.Reader
	// Callback funtion when get a msg
	PullCb func(m *kafka.Message) error
	//err        chan error
	IsFetchMode bool
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
}

/*
* @bref: Fetch a single message then commit ever single time
  - will make performance lost
*/
func (r *Reader) fetchRun() {
	go func() {
		defer r.Reader.Close()
		for {
			// Fetch one msg
			m, e := r.Reader.FetchMessage(context.Background())
			if e != nil {
				//r.err <- errors.New("error occur when reader <fetch> a message")
				l.Errorln("error occur when reader <fetch> a message")
			}
			// Handle it
			if e = r.PullCb(&m); e != nil {
				// r.err <- e
				l.Errorln("error occur when handle message")
			}
			// Commit when no exception
			e = r.Reader.CommitMessages(context.Background(), m)
			if e != nil {
				//r.err <- errors.New("error cocur when commit a message")
				l.Errorln("error occur when commit a message")
			}
		}
	}()
}

func (r *Reader) run() {
	go func() {
		defer r.Reader.Close()
		for {
			m, e := r.Reader.ReadMessage(context.Background())
			if e != nil {
				//r.err <- errors.New("error occur when reader <read> a message")
				l.Errorln("error occur when reader <fetch> a message")
			}
			// Handle it
			if e = r.PullCb(&m); e != nil {
				//r.err <- e
				l.Errorln("error occur when handle message")
			}
		}
	}()
}
