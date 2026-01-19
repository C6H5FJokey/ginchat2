package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserGroupStateModel = (*customUserGroupStateModel)(nil)

type (
	// UserGroupStateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserGroupStateModel.
	UserGroupStateModel interface {
		userGroupStateModel
	}

	customUserGroupStateModel struct {
		*defaultUserGroupStateModel
	}
)

// NewUserGroupStateModel returns a model for the database table.
func NewUserGroupStateModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserGroupStateModel {
	return &customUserGroupStateModel{
		defaultUserGroupStateModel: newUserGroupStateModel(conn, c, opts...),
	}
}
