//退款接口
package trade

import (
	"reflect"

	"github.com/fromiuan/chinapnr/encoding"
)

type RefundReq struct {
	OrderDate            string            `json:"order_date"`              // 必须 订单日期
	OrderID              string            `json:"order_id"`                // 必须 订单号
	TransAmt             string            `json:"trans_amt"`               // 必须 退款金额
	OrginalPlatformSeqID string            `json:"orginal_platform_seq_id"` // 必须 原交易平台流水id
	DivDetail            []RefundDivDetail `json:"-"`                       // 必须 分账账户串
	BgRetURL             string            `json:"bg_ret_url"`              // 必须 商户后台应答地址

	QuickpayPageFlag string `json:"quickpay_page_flag"` // 可选 快捷支付页面版标识
	UserCustID       string `json:"user_cust_id"`       // 可选 用户客户号
	MerPriv          string `json:"mer_priv"`           // 可选 商户私有域
	Extension        string `json:"extension"`          // 可选 扩展域
}

type RefundRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	UserCustID    string `json:"user_cust_id"`
	OrderDate     string `json:"order_date"`
	OrderID       string `json:"order_id"`
	PlatformSeqID string `json:"platform_seq_id"`
	TransAmt      string `json:"trans_amt"`
	TransStat     string `json:"trans_stat"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}

type RefundDivDetail struct {
	DivCustId string `json:"divCustId"`
	DivAcctId string `json:"divAcctId"`
	DivAmt    string `json:"divAmt"`
}

func (this RefundReq) APICmId() string {
	return "205"
}

func (this RefundReq) Params() map[string]string {
	data := make(map[string]string)
	divDetail, _ := encoding.Marshal(this.DivDetail)
	data["div_detail"] = string(divDetail)
	return data
}

func (this RefundReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
