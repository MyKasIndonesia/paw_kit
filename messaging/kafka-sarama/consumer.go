package kafka_sarama

import (
	"context"
	"github.com/MyKasIndonesia/paw_kit/messaging"
)

func (l *Kafka) AddTopicListener(topic string, callback messaging.CallbackFunc) {
	l.mu.Lock()
	defer func() {
		l.mu.Unlock()
	}()
	functions := l.CallbackFunctions[topic]
	functions = append(functions, callback)
	l.CallbackFunctions[topic] = functions
	l.Option.ListTopics = append(l.Option.ListTopics, topic)
}

func (l *Kafka) Listen() {
	if l.Consumer != nil {
		return
	}

	var err error
	l.Consumer, err = l.NewListener(l.Option)
	if err != nil {
		return
	}

	go func() {
		for err := range l.Consumer.Errors() {
			l.Option.Log.Info(context.Background(), "Error: %s\n", err.Error())
		}
	}()

	go func() {
		for ntf := range l.Consumer.Notifications() {
			l.Option.Log.Info(context.Background(), "Rebalanced: %+v\n", ntf)
		}
	}()

	go func() {
		for {
			select {
			case msg, ok := <-l.Consumer.Messages():
				if ok {
					l.Consumer.MarkOffset(msg, "") // mark message as processed
					for _, function := range l.CallbackFunctions[msg.Topic] {
						err := function(msg.Value)
						if err != nil {
							l.Option.Log.Error(context.Background(), "consumer.Listen() ERROR at CallBackFunc", nil, err)
						}
					}
				} else {
					l.Option.Log.Info(context.Background(), "error", nil)
				}
			}
		}
	}()
}
