package homestayBusiness

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"look-cp/app/travel/cmd/api/internal/svc"
	"look-cp/app/travel/cmd/api/internal/types"
	"look-cp/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessListLogic {
	return &HomestayBusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessListLogic) HomestayBusinessList(req *types.HomestayBusinessListReq) (*types.HomestayBusinessListResp, error) {
	// todo: add your logic here and delete this line

	whereBuilder := l.svcCtx.HomestayBusinessModel.SelectBuilder()
	list, err := l.svcCtx.HomestayBusinessModel.FindPageListByIdDESC(l.ctx, whereBuilder, req.LastId, req.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HomestayBussinessList FindPageListByIdDESC db fail ,  req : %+v , err:%v", req, err)
	}

	var resp []types.HomestayBusinessListInfo
	if len(list) > 0 {
		for _, item := range list {
			var typeHomestayBusinessListInfo types.HomestayBusinessListInfo
			_ = copier.Copy(&typeHomestayBusinessListInfo, item)

			resp = append(resp, typeHomestayBusinessListInfo)
		}
	}

	return &types.HomestayBusinessListResp{
		List: resp,
	}, nil
}
