// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package persona

import (
	"context"
	"strconv"

	"ginchat2/internal/svc"
	"ginchat2/internal/types"
	"ginchat2/models"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type DelPersonaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelPersonaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPersonaLogic {
	return &DelPersonaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelPersonaLogic) DelPersona(req *types.DelPersonaReq) error {
	userId, err := strconv.ParseInt(l.ctx.Value("userId").(string), 10, 64)
	if err == models.ErrNotFound {
		return errors.New(5002, "不存在该用户")
	}
	persona, err := l.svcCtx.PersonaModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
		break
	case models.ErrNotFound:
		return errors.New(2003, "不存在该persona")
	default:
		return err
	}
	if persona.UserId != userId {
		return errors.New(2003, "这个persona不是你的!")
	}
	if err := l.svcCtx.PersonaModel.Delete(l.ctx, req.Id); err != nil {
		return err
	}
	return nil
}
