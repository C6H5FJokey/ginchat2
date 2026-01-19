// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package persona

import (
	"net/http"

	"ginchat2/internal/logic/persona"
	"ginchat2/internal/svc"
	"ginchat2/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
)

func AddPersonaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddPersonaReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := persona.NewAddPersonaLogic(r.Context(), svcCtx)
		err := l.AddPersona(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
