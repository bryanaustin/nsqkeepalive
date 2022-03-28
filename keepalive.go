// A handler wrapper for keeping alive (touching) NSQ messages on a fixed interval.
package nsqkeepalive

import (
	"github.com/nsqio/go-nsq"
	"time"
)

// HandlerWrapper implements nsq.Handler and touches the message on the specified interval
type HandlerWrapper struct {
	Interval time.Duration
	Child    nsq.Handler
}

// Handler is a continence method for creating HandlerWrapper
func Handler(i time.Duration, h nsq.Handler) nsq.Handler {
	return &HandlerWrapper{Interval: i, Child: h}
}

// HandleMessage implements nsq.Handler for HandlerWrapper
func (h HandlerWrapper) HandleMessage(m *nsq.Message) error {
	t := time.NewTicker(h.Interval)
	defer t.Stop()
	cm := h.channeledMessage(m)
	for {
		select {
		case <-t.C:
			m.Touch()
		case e := <-cm:
			return e
		}
	}
	return nil
}

// channeledMessage will convert the message handler into channel for evented communication
func (h HandlerWrapper) channeledMessage(m *nsq.Message) <-chan error {
	ec := make(chan error)
	go func() {
		defer close(ec)
		ec <- h.Child.HandleMessage(m)
	}()
	return ec
}
