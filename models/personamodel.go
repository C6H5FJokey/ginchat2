package models

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PersonaModel = (*customPersonaModel)(nil)

var (
	cachePublicPersonaPrefix = "cache:public:persona:"
)

type (
	// PersonaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPersonaModel.
	PersonaModel interface {
		personaModel
		FindOneByUserAndName(ctx context.Context, userId int64, name string) (*Persona, error)
		FindAllByUser(ctx context.Context, userId int64) ([]*Persona, error)
	}

	customPersonaModel struct {
		*defaultPersonaModel
	}
)

// NewPersonaModel returns a model for the database table.
func NewPersonaModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PersonaModel {
	return &customPersonaModel{
		defaultPersonaModel: newPersonaModel(conn, c, opts...),
	}
}

func (m *customPersonaModel) FindOneByUserAndName(ctx context.Context, userId int64, name string) (*Persona, error) {
	publicPersonaUserIdAndNameKey := m.formatPublicPersonaUserIdAndNameKey(userId, name)
	var resp Persona
	err := m.QueryRowIndexCtx(ctx, &resp, publicPersonaUserIdAndNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (any, error) {
		query := fmt.Sprintf("select %s from %s where user_id = $1 and name = $2 limit 1", personaRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPersonaModel) Update(ctx context.Context, data *Persona) error {
	publicPersonaIdKey := fmt.Sprintf("%s%v", cachePublicPersonaIdPrefix, data.Id)
	var persona Persona
	err := m.QueryRowCtx(ctx, &persona, publicPersonaIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", personaRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, data.Id)
	})
	if err != nil {
		return nil
	}
	m.DelCacheCtx(ctx, m.formatPublicPersonaUserIdAndNameKey(persona.UserId, persona.Name))
	m.DelCacheCtx(ctx, m.formatPublicPersonaUserIdAndNameKey(data.UserId, data.Name))
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, personaRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.Name, data.Prompt)
	}, publicPersonaIdKey)
	return err
}

func (m *customPersonaModel) FindAllByUser(ctx context.Context, userId int64) ([]*Persona, error) {
	var resp []*Persona
	query := fmt.Sprintf("select %s from %s where user_id = $1", personaRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customPersonaModel) formatPublicPersonaUserIdAndNameKey(userId int64, name string) string {
	return fmt.Sprintf("%s:user:%v:name:%v", cachePublicPersonaPrefix, userId, name)
}
