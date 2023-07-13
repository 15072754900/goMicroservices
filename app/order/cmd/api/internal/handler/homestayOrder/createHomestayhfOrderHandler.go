package homestayOrder

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"look-cp/app/order/cmd/api/internal/logic/homestayOrder"
	"look-cp/app/order/cmd/api/internal/svc"
	"look-cp/app/order/cmd/api/internal/types"
)

func CreateHomestayhfOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateHomestayOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			// 这里作者使用result中间件，但是我还不知道中间件在哪注册，所以暂时不用。
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := homestayOrder.NewCreateHomestayhfOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateHomestayhfOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
