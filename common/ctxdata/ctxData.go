package ctxdata

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from uid
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if uid, err := jsonUid.Int64(); err == nil {
			uid = uid
		} else {
			logx.WithContext(ctx).Error("GetUidFromCtx err %+v", err)
		}
	}
	return uid
}
