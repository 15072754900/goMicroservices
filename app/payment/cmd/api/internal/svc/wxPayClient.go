package svc

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"look-cp/app/payment/cmd/api/internal/config"
	"look-cp/common/xerr"
)

func NewWxPayClientV3(c config.Config) (*core.Client, error) {
	mchPrivateKey, err := utils.LoadPrivateKey(c.WxPayConf.PrivateKey)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("wechat pay fail"), " wechat pay init fail ，mchPrivateKey err : %v \n", err)
	}

	ctx := context.Background()
	// Initialize the client with the merchant's private key, ect.. and make it able to check cert and some regular check ...
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(c.WxPayConf.MchId, c.WxPayConf.SerialNo, mchPrivateKey, c.WxPayConf.APIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("wechat pay fail"), "new wechat pay client err:%s", err)
	}

	return client, nil
}
