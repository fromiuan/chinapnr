// 4.3.6.	代扣支付接口
package trade

import (
	"reflect"
)

type WithholdPayReq struct {
	UserCustID string `json:"user_cust_id"` // 必须 用户客户号
	OrderDate  string `json:"order_date"`   // 必须 订单日期
	OrderID    string `json:"order_id"`     // 必须 订单号
	TransAmt   string `json:"trans_amt"`    // 必须 交易金额
	BindCardID string `json:"bind_card_id"` // 必须 绑定银行卡ID
	SignSeqID  string `json:"sign_seq_id"`  // 必须 代扣签约流水
	InCustID   string `json:"in_cust_id"`   // 必须 入账客户号
	InAcctID   string `json:"in_acct_id"`   // 必须 入账账户号
	BgRetURL   string `json:"bg_ret_url"`   // 必须 商户后台应答地址

	MerPriv        string `json:"mer_priv"`         // 可选 商户私有域
	Extension      string `json:"extension"`        // 可选 扩展域
	SecondaryMerID string `json:"secondary_mer_id"` // 可选 二级商户号
}

type WithholdPayRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	UserCustID    string `json:"user_cust_id"`
	OrderDate     string `json:"order_date"`
	OrderID       string `json:"order_id"`
	PlatformSeqID string `json:"platform_seq_id"`
	TransAmt      string `json:"trans_amt"`
	BindCardID    string `json:"bind_card_id"`
	SignSeqID     string `json:"sign_seq_id"`
	InCustID      string `json:"in_cust_id"`
	InAcctID      string `json:"inAcctId"`
	FeeAmt        string `json:"fee_amt"`
	FeeCustID     string `json:"fee_cust_id"`
	FeeAcctID     string `json:"fee_acct_id"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}

func (this WithholdPayReq) APICmId() string {
	return "207"
}

func (this WithholdPayReq) Params() map[string]string {
	return nil
}

func (this WithholdPayReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
