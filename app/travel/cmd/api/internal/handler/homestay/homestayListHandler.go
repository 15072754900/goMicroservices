package homestay

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"look-cp/app/travel/cmd/api/internal/logic/homestay"
	"look-cp/app/travel/cmd/api/internal/svc"
	"look-cp/app/travel/cmd/api/internal/types"
)

func HomestayListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestay.NewHomestayListLogic(r.Context(), svcCtx)
		resp, err := l.HomestayList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
