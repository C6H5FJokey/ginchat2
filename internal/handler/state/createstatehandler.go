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

func CreateStateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateStateReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := state.NewCreateStateLogic(r.Context(), svcCtx)
		err := l.CreateState(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
