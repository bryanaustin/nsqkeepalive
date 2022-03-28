# NSQ Keep Alive
[![Go Reference](https://pkg.go.dev/badge/github.com/bryanaustin/nsqkeepalive.svg)](https://pkg.go.dev/github.com/bryanaustin/nsqkeepalive)

By default [NSQ](https://github.com/nsqio/nsq) will requeue a message if there is no response after 1 minute. This library will touch the message on a regular interval (I use 50 seconds) to let NSQ know it's still being processed.

## Example
```go
package main

import (
	"github.com/bryanaustin/nsqkeepalive"
	"github.com/nsqio/go-nsq"
	"time"
)

func main() {
	nsqconfig := nsq.NewConfig()
	consumer, err = nsq.NewConsumer("topic", "channel", nsqconfig)
	if err != nil {
		// Handle err
	}
	handlerobj := nsq.HandlerFunc(handler)
	wrappedhandler := nsqkeepalive.Handler(time.Second * 50, handlerobj)
	consumer.AddHandler(wrappedhandler)
	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		// Handle err
	}
	// ...
}

func handler(m *nsq.Message) error {
	// This message will be touched until it returns
}
```
