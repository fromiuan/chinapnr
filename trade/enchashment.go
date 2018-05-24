// 取现接口
package trade

import (
	"reflect"
)

type EnchashmentReq struct {
	UserCustID     string `json:"user_cust_id"`      // 必须 用户客户号
	OrderDate      string `json:"order_date"`        // 必须 订单日期
	OrderID        string `json:"order_id"`          // 必须 订单号
	TransAmt       string `json:"trans_amt"`         // 必须 交易金额
	CashBindCardID string `json:"cash_bind_card_id"` // 必须 取现绑定银行卡ID
	CashType       string `json:"cash_type"`         // 必须 取现方式
	BgRetURL       string `json:"bg_ret_url"`        // 必须 商户后台应答地址
	FeeObj         string `json:"fee_obj"`           // 可选 手续费收取对象
	FeeAcctID      string `json:"fee_acct_id"`       // 可选 手续费收取子账户
	MerPriv        string `json:"mer_priv"`          // 可选 商户私有域
	Extension      string `json:"extension"`         // 可选 扩展域
}

type EnchashmentRsp struct {
	CmdID          string `json:"cmd_id"`
	RespCode       string `json:"resp_code"`
	RespDesc       string `json:"resp_desc"`
	MerCustID      string `json:"mer_cust_id"`
	UserCustID     string `json:"user_cust_id"`
	OrderDate      string `json:"order_date"`
	OrderID        string `json:"order_id"`
	PlatformSeqID  string `json:"platform_seq_id"`
	TransAmt       string `json:"trans_amt"`
	RealTransAmt   string `json:"real_trans_amt"`
	CashBindCardID string `json:"cash_bind_card_id"`
	BankID         string `json:"bank_id"`
	FeeAmt         string `json:"fee_amt"`
	FeeCustID      string `json:"fee_cust_id"`
	FeeAcctID      string `json:"fee_acct_id"`
	TransStat      string `json:"trans_stat"`
	BgRetURL       string `json:"bg_ret_url"`
	MerPriv        string `json:"mer_priv"`
	Extension      string `json:"extension"`
}

func (this EnchashmentReq) APICmId() string {
	return "202"
}

func (this EnchashmentReq) Params() map[string]string {
	return nil
}

func (this EnchashmentReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
