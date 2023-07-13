package homestay

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"look-cp/app/travel/cmd/api/internal/logic/homestay"
	"look-cp/app/travel/cmd/api/internal/svc"
	"look-cp/app/travel/cmd/api/internal/types"
)

func BusinessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BusinessListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 执行类似于洋葱模型的运行逻辑，先访问逻辑代码部分然后获取返回值，有时候会要求访问api的logic获取一个针对rpc的req然后去rpc的logic。总之就是获取一个返回值。
		l := homestay.NewBusinessListLogic(r.Context(), svcCtx)
		resp, err := l.BusinessList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
