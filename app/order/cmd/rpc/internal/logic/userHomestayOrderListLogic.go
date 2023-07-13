package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"look-cp/app/order/model"
	"look-cp/common/xerr"

	"look-cp/app/order/cmd/rpc/internal/svc"
	"look-cp/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserHomestayOrderList 获取更新后的民宿订单 (后面要整体把握一下，需求和代码之间的关系，实现的逻辑)
func (l *UserHomestayOrderListLogic) UserHomestayOrderList(in *pb.UserHomestayOrderListReq) (*pb.UserHomestayOrderListResp, error) {
	// todo: add your logic here and delete this line

	whereBuilder := l.svcCtx.HomestayOrderModel.SelectBuilder()
	// 如果筛选的部分支持，否则全部输出
	if in.TradeState >= model.HomestayOrderTradeStateCancel && in.TradeState <= model.HomestayOrderTradeStateExpire {
		whereBuilder = whereBuilder.Where(squirrel.Eq{"trade_state": in.TradeState})
	}

	list, err := l.svcCtx.HomestayOrderModel.FindPageListByIdDESC(l.ctx, whereBuilder, in.LastId, in.PageSize)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get users's homestay order err : %v , in : %+v", err, in)
	}

	var resp []*pb.HomestayOrder
	if len(list) > 0 {
		for _, homestayOrder := range list {
			var pbHomestayOrder pb.HomestayOrder
			_ = copier.Copy(&pbHomestayOrder, homestayOrder)

			pbHomestayOrder.CreateTime = homestayOrder.CreateTime.Unix()
			pbHomestayOrder.LiveStartDate = homestayOrder.LiveStartDate.Unix()
			pbHomestayOrder.LiveEndDate = homestayOrder.LiveEndDate.Unix()

			resp = append(resp, &pbHomestayOrder)
		}
	}
	return &pb.UserHomestayOrderListResp{
		List: resp,
	}, nil
}
