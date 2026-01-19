// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Database sqlx.SqlConf
	Cache    cache.CacheConf
	Auth     struct {
		AccessSecret string
		AccessExpire int64
	}
	MsgChannel struct {
		ChannelType string
		KafkaHosts  string
		KafkaTopic  string
	}
	OpenAI struct {
		BaseURL string
		Model   string
		Token   string
	}
}
