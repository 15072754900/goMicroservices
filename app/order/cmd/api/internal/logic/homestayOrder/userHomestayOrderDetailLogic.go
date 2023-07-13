package homestayOrder

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"look-cp/app/order/cmd/rpc/order"
	"look-cp/app/order/model"
	"look-cp/app/payment/cmd/rpc/payment"
	"look-cp/common/ctxdata"
	"look-cp/common/tool"
	"look-cp/common/xerr"

	"look-cp/app/order/cmd/api/internal/svc"
	"look-cp/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderDetailLogic {
	return &UserHomestayOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderDetailLogic) UserHomestayOrderDetail(req types.UserHomestayOrderDetailReq) (*types.UserHomestayOrderDetailResp, error) {
	// todo: add your logic here and delete this line
	// Finish done work

	userId := ctxdata.GetUidFromCtx(l.ctx)

	resp, err := l.svcCtx.OrderRpc.HomestayOrderDetail(l.ctx, &order.HomestayOrderDetailReq{
		Sn: req.Sn,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get homestay order detali fail"), " rpc get HomestayOrderDetail err %+v err, sn : %v", err, req.Sn)
	}

	var typesOrderDetail types.UserHomestayOrderDetailResp
	if resp.HomestayOrder != nil && resp.HomestayOrder.UserId == userId {
		copier.Copy(&typesOrderDetail, resp.HomestayOrder)

		// 重置价格, 获取起始数据并将其转换为元， 保留指定位数并转化为小数
		typesOrderDetail.OrderTotalPrice = tool.Fen2Yuan(resp.HomestayOrder.OrderTotalPrice)
		typesOrderDetail.FoodTotalPrice = tool.Fen2Yuan(resp.HomestayOrder.FoodTotalPrice)
		typesOrderDetail.HomestayTotalPrice = tool.Fen2Yuan(resp.HomestayOrder.HomestayTotalPrice)
		typesOrderDetail.HomestayPrice = tool.Fen2Yuan(resp.HomestayOrder.HomestayPrice)
		typesOrderDetail.FoodPrice = tool.Fen2Yuan(resp.HomestayOrder.FoodPrice)
		typesOrderDetail.MarketHomestayPrice = tool.Fen2Yuan(resp.HomestayOrder.MarketHomestayPrice)

		// 支付信息.
		if typesOrderDetail.TradeState != model.HomestayOrderTradeStateCancel {
			paymentResp, err := l.svcCtx.PaymentRpc.GetPaymentSuccessRefundByOrderSn(l.ctx, &payment.GetPaymentSuccessRefundByOrderSnReq{
				OrderSn: resp.HomestayOrder.Sn,
			})
			if err != nil {
				logx.WithContext(l.ctx).Errorf("Failed to get order payment information err : %v , orderSn:%s", err, resp.HomestayOrder.Sn)
			}

			if paymentResp.PaymentDetail != nil {
				typesOrderDetail.PayTime = paymentResp.PaymentDetail.PayTime
				typesOrderDetail.PayType = paymentResp.PaymentDetail.PayMode
			}
		}

		return &typesOrderDetail, nil
	}
	return nil, nil
}
