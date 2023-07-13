package homestayBusiness

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"look-cp/app/travel/model"
	"look-cp/app/usercenter/cmd/rpc/usercenter"
	"look-cp/common/xerr"

	"look-cp/app/travel/cmd/api/internal/svc"
	"look-cp/app/travel/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessDetailLogic {
	return &HomestayBusinessDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessDetailLogic) HomestayBusinessDetail(req types.HomestayBusinessDetailReq) (*types.HomestayBusinessDetailResp, error) {

	homestayBusiness, err := l.svcCtx.HomestayBusinessModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), " HomestayBussinessDetail  FindOne db fail ,id  : %d , err : %v", req.Id, err)
	}

	var typeHomestayBusinessBoss types.HomestayBusinessBoss
	if homestayBusiness != nil {

		userResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
			Id: homestayBusiness.UserId,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("get boss info fail"), "get boss info fail ,  userId : %d ,err:%v", homestayBusiness.UserId, err)
		}
		if userResp.User != nil && userResp.User.Id > 0 {
			_ = copier.Copy(&typeHomestayBusinessBoss, userResp.User)
		}
	}

	return &types.HomestayBusinessDetailResp{
		Boss: typeHomestayBusinessBoss,
	}, nil
}
