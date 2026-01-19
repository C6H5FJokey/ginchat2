package kafka

import (
	"strings"
	"sync"

	"github.com/IBM/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

var consumer sarama.Consumer

type ConsumerCallback func(data []byte)

// 初始化消费者
func InitConsumer(hosts string) (err error) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if nil != err {
		logx.Error("init kafka consumer client error:", err)
		return
	}

	consumer, err = sarama.NewConsumerFromClient(client)
	if nil != err {
		logx.Error("init kafka consumer error:", err)
		return
	}
	return nil
}

// 消费消息，通过回调函数进行
func ConsumerMsg(callBack ConsumerCallback) {
	if consumer == nil {
		logx.Error("kafka consumer not initialized")
		return
	}
	if topic == "" {
		logx.Error("kafka topic is empty")
		return
	}
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		logx.Error("kafka list partitions error:", err)
		return
	}
	if len(partitions) == 0 {
		logx.Error("kafka no partitions for topic:", topic)
		return
	}

	var wg sync.WaitGroup
	for _, p := range partitions {
		partitionID := p
		pc, err := consumer.ConsumePartition(topic, partitionID, sarama.OffsetNewest)
		if err != nil {
			logx.Error("ConsumePartition error, partition:", partitionID, " err:", err)
			continue
		}
		wg.Add(1)
		go func(partitionConsumer sarama.PartitionConsumer) {
			defer wg.Done()
			defer partitionConsumer.Close()
			for msg := range partitionConsumer.Messages() {
				if callBack != nil {
					callBack(msg.Value)
				}
			}
		}(pc)
	}
	// block forever (or until all partition consumers exit)
	wg.Wait()
}

func CloseConsumer() {
	if nil != consumer {
		consumer.Close()
	}
}
