// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"database/sql"

	"ginchat2/internal/svc"
	"ginchat2/internal/types"
	"ginchat2/models"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/x/errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) error {
	if _, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username); err != sqlx.ErrNotFound {
		return errors.New(5001, "该用户名已注册")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	email := sql.NullString{
		Valid:  req.Email != "",
		String: req.Email,
	}
	phone := sql.NullString{
		Valid:  req.Phone != "",
		String: req.Phone,
	}
	l.svcCtx.UserModel.Insert(l.ctx, &models.User{
		Id:           0,
		Username:     req.Username,
		PasswordHash: string(hash),
		Email:        email,
		Phone:        phone,
	})
	return nil
}
