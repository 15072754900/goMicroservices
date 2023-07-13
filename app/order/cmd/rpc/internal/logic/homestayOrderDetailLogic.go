package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"look-cp/app/order/model"
	"look-cp/common/xerr"

	"look-cp/app/order/cmd/rpc/internal/svc"
	"look-cp/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayOrderDetailLogic {
	return &HomestayOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomestayOrderDetail 民宿订单详情
func (l *HomestayOrderDetailLogic) HomestayOrderDetail(in *pb.HomestayOrderDetailReq) (*pb.HomestayOrderDetailResp, error) {
	// todo: add your logic here and delete this line

	homestayOrder, err := l.svcCtx.HomestayOrderModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "HomestayOrderModel  FindOneBySn db err : %v , sn : %s", err, in.Sn)
	}

	var resp pb.HomestayOrder // 这里需要的就是一个homestayOrder
	if homestayOrder != nil {
		_ = copier.Copy(&resp, homestayOrder)

		resp.CreateTime = homestayOrder.CreateTime.Unix()
		resp.LiveStartDate = homestayOrder.LiveStartDate.Unix()
		resp.LiveEndDate = homestayOrder.LiveEndDate.Unix()
	}

	return &pb.HomestayOrderDetailResp{
		HomestayOrder: &resp,
	}, nil
}
