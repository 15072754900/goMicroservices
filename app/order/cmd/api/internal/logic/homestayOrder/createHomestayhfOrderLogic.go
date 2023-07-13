package homestayOrder

import (
	"context"
	"github.com/pkg/errors"
	"look-cp/app/order/cmd/rpc/order"
	"look-cp/app/travel/cmd/rpc/pb"
	"look-cp/common/ctxdata"
	"look-cp/common/xerr"

	"look-cp/app/order/cmd/api/internal/svc"
	"look-cp/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHomestayhfOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateHomestayhfOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHomestayhfOrderLogic {
	return &CreateHomestayhfOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateHomestayhfOrder create a order
func (l *CreateHomestayhfOrderLogic) CreateHomestayhfOrder(req *types.CreateHomestayOrderReq) (*types.CreateHomestayOrderResp, error) {
	// todo: myMessage already done.

	// 通过之前注册的rpc服务，在这里调用并判断返回值，以及使用
	homestayResp, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &pb.HomestayDetailReq{
		Id: req.HomestayId,
	})
	if err != nil {
		return nil, err
	}
	if homestayResp.Homestay != nil || homestayResp.Homestay.Id == 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("homestay no exist"), "CreateHomestayOrder homestay no exists id : %d", req.HomestayId)
	}

	userId := ctxdata.GetUidFromCtx(l.ctx)

	resp, err := l.svcCtx.OrderRpc.CreateHomestayOrder(l.ctx, &order.CreateHomestayOrderReq{
		HomestayId:    req.HomestayId,
		IsFood:        req.IsFood,
		LiveStartTime: req.LiveStartTime,
		LiveEndTime:   req.LiveEndTime,
		UserId:        userId,
		LivePeopleNum: req.LivePeopleNum,
		Remark:        req.Remark,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("create homestay order fail"), "create homestay order rpc CreateHomestayOrder fail req: %+v, err: %v", req, err)
	}

	return &types.CreateHomestayOrderResp{
		OrderSn: resp.Sn,
	}, nil
}
