// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"strconv"
	"time"

	"ginchat2/internal/svc"
	"ginchat2/internal/types"
	"ginchat2/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err == models.ErrNotFound {
		return nil, errors.New(5002, "不存在该用户")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New(5003, "密码错误")
	}
	expiresAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire))
	claims := jwt.MapClaims{
		"sub":    strconv.FormatInt(user.Id, 10), // 标准字段
		"exp":    expiresAt.Unix(),
		"iat":    time.Now().Unix(),
		"userId": strconv.FormatInt(user.Id, 10), // 这个会被存入上下文
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// HS256 expects a byte slice key, otherwise jwt will raise "key is of invalid type".
	signed, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return nil, err
	}
	resp = &types.LoginResp{
		Token: signed,
	}
	return
}
