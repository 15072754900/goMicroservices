package uniqueid

import (
	"fmt"
	"look-cp/common/tool"
	"time"
)

// SnPrefix 生成sn单号
type SnPrefix string

const (
	SN_PREFIX_HOMESTAY_ORDER SnPrefix = "HSO" // 民宿订单前缀
	SN_PREFIX_THIRD_PAYMENT  SnPrefix = "PMT" // 第三方支付流水记录前缀
)

// GenSn 生成单号
func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20220619120000"), tool.Krand(8, tool.KC_RAND_KIND_NUM))
}
