package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"look-cp/app/usercenter/cmd/rpc/usercenter"
	"look-cp/app/usercenter/model"
	"look-cp/common/xerr"

	"look-cp/app/usercenter/cmd/rpc/internal/svc"
	"look-cp/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdLogic {
	return &GetUserAuthByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdLogic) GetUserAuthByUserId(in *pb.GetUserAuthByUserIdReq) (*pb.GetUserAuthByUserIdResp, error) {
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByUserIdAuthType(l.ctx, in.UserId, in.AuthType)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth  fail"), "err : %v , in : %+v", err, in)
	}

	var respUserAuth usercenter.UserAuth
	_ = copier.Copy(&respUserAuth, userAuth)

	return &pb.GetUserAuthByUserIdResp{
		UserAuth: &respUserAuth,
	}, nil
}
