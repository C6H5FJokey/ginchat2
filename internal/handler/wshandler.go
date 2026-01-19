// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"
	"strconv"

	"ginchat2/internal/svc"
	"ginchat2/internal/types"
	"ginchat2/internal/ws"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
)

func wsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	hub := ws.GetHub(svcCtx)
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errors.New(4001, err.Error()))
			logx.Debug(err)
			return
		}

		claims := jwt.MapClaims{}
		tok, err := jwt.ParseWithClaims(req.Token, claims, func(token *jwt.Token) (any, error) {
			return []byte(svcCtx.Config.Auth.AccessSecret), nil
		})
		if err != nil || tok == nil || !tok.Valid {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New(4002, err.Error()))
			logx.Debug(err)
			return
		}

		uidStr, _ := claims["userId"].(string)
		if uidStr == "" {
			uidStr, _ = claims["sub"].(string)
		}
		userId, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil || userId <= 0 {
			httpx.Error(w, errors.New(4003, err.Error()))
			return
		}
		logx.Debug(userId)
		if _, err := ws.NewClient(r.Context(), w, r, hub, logx.WithContext(r.Context()), userId); err != nil {
			// 只有在 Upgrade 失败(未 hijack)时才允许写 HTTP 响应
			httpx.Error(w, err)
			return
		}
		// Upgrade 成功后不能再写任何 HTTP 响应（否则会 hijacked 报错）
	}
}
