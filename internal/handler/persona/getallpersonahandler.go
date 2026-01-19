// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package persona

import (
	"net/http"

	"ginchat2/internal/logic/persona"
	"ginchat2/internal/svc"
	xhttp "github.com/zeromicro/x/http"
)

func GetAllPersonaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := persona.NewGetAllPersonaLogic(r.Context(), svcCtx)
		resp, err := l.GetAllPersona()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
