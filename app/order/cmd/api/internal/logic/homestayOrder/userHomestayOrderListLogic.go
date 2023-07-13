package homestayOrder

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"look-cp/app/order/cmd/rpc/order"
	"look-cp/common/ctxdata"
	"look-cp/common/tool"
	"look-cp/common/xerr"

	"look-cp/app/order/cmd/api/internal/svc"
	"look-cp/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderListLogic) UserHomestayOrderList(req types.UserHomestayOrderListReq) (*types.UserHomestayOrderListResp, error) {
	// todo: add your logic here and delete this line
	// work done

	userId := ctxdata.GetUidFromCtx(l.ctx) // get user login id

	resp, err := l.svcCtx.OrderRpc.UserHomestayOrderList(l.ctx, &order.UserHomestayOrderListReq{
		UserId:     userId,
		TradeState: req.TradeState,
		PageSize:   req.PageSize,
		LastId:     req.LastId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get user homestay order list"), "Failed to get user homestay order list err : %v ,req: %+v", err, req)
	}
	var typesUserHomestayOrderList []types.UserHomestayOrderListView

	if len(resp.List) > 0 {
		for _, homestayOrder := range resp.List {
			var typeHomestayOrder types.UserHomestayOrderListView
			_ = copier.Copy(&typeHomestayOrder, homestayOrder)

			typeHomestayOrder.OrderTotalPrice = tool.Fen2Yuan(homestayOrder.OrderTotalPrice)

			typesUserHomestayOrderList = append(typesUserHomestayOrderList, typeHomestayOrder)
		}
	}

	return &types.UserHomestayOrderListResp{
		List: typesUserHomestayOrderList,
	}, nil
}
