package kafka

import (
	"strings"

	"github.com/IBM/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

var producer sarama.AsyncProducer

func InitProducer(hosts string) (err error) {
	config := sarama.NewConfig()
	config.Producer.Compression = sarama.CompressionGZIP
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if nil != err {
		logx.Error("init kafka client error")
		return
	}

	producer, err = sarama.NewAsyncProducerFromClient(client)
	if nil != err {
		logx.Error("init kafka async client error")
		return
	}
	return nil
}

func Send(data []byte) {
	if producer == nil {
		logx.Error("kafka producer not initialized")
		return
	}
	be := sarama.ByteEncoder(data)
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: be}
}

func Close() {
	if producer != nil {
		producer.Close()
	}
}
