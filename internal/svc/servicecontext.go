// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	constant "ginchat2/common"
	"ginchat2/internal/ai"
	"ginchat2/internal/config"
	"ginchat2/internal/kafka"
	"ginchat2/models"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type ServiceContext struct {
	Config       config.Config
	Gateway      ai.Gateway
	UserModel    models.UserModel
	MessageModel models.MessageModel
	PersonaModel models.PersonaModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.Database.DataSource)
	gateway, err := ai.NewGateway(
		c.OpenAI.BaseURL,
		c.OpenAI.Model,
		c.OpenAI.Token,
	)
	if err != nil {
		logx.Error("gateway init fail", err)
	}
	if c.MsgChannel.ChannelType == constant.KAFKA {
		kafka.SetTopic(c.MsgChannel.KafkaTopic)
		if err := kafka.InitProducer(c.MsgChannel.KafkaHosts); err != nil {
			logx.Error("Kafka Producer init fail:", err)
			c.MsgChannel.ChannelType = constant.GO_CHANNEL
		}
		if err := kafka.InitConsumer(c.MsgChannel.KafkaHosts); err != nil {
			logx.Error("Kafka Comsumer init fail:", err)
			c.MsgChannel.ChannelType = constant.GO_CHANNEL
		}
	}
	return &ServiceContext{
		Config:  c,
		Gateway: *gateway,
		UserModel: models.NewUserModel(
			conn,
			c.Cache,
			cache.WithExpiry(time.Hour*24),
		),
		MessageModel: models.NewMessageModel(
			conn,
			c.Cache,
			cache.WithExpiry(time.Hour*24),
		),
		PersonaModel: models.NewPersonaModel(
			conn,
			c.Cache,
			cache.WithExpiry(time.Hour*24),
		),
	}
}
