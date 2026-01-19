package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StateModel = (*customStateModel)(nil)

type (
	// StateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStateModel.
	StateModel interface {
		stateModel
	}

	customStateModel struct {
		*defaultStateModel
	}
)

// NewStateModel returns a model for the database table.
func NewStateModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StateModel {
	return &customStateModel{
		defaultStateModel: newStateModel(conn, c, opts...),
	}
}
