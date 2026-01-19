// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package state

import (
	"net/http"

	"ginchat2/internal/logic/state"
	"ginchat2/internal/svc"
	"ginchat2/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func AppendStateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppendStateReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := state.NewAppendStateLogic(r.Context(), svcCtx)
		err := l.AppendState(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
