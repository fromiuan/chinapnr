//转账接口
package trade

import (
	"reflect"
)

type TransferReq struct {
	OrderID     string `json:"order_id"`     // 必须 订单号
	OrderDate   string `json:"order_date"`   // 必须 订单日期
	OutCustID   string `json:"out_cust_id"`  // 必须 出账客户号
	OutAcctID   string `json:"out_acct_id"`  // 必须 出账子账户号
	InCustID    string `json:"in_cust_id"`   // 必须 入账客户号
	InAcctID    string `json:"in_acct_id"`   // 必须 入账子账户号
	TransferAmt string `json:"transfer_amt"` // 必须 转账金额
	BgRetURL    string `json:"bg_ret_url"`   // 必须 商户后台应答地址
	MerPriv     string `json:"mer_priv"`     // 必须 商户私有域

	TransferType string `json:"transfer_type"` // 可选 转账类型
	Extension    string `json:"extension"`     // 可选 入参扩展域
}

type TransferRsp struct {
	CmdID        string `json:"cmd_id"`
	RespCode     string `json:"resp_code"`
	RespDesc     string `json:"resp_desc"`
	MerCustID    string `json:"mer_cust_id"`
	OrderID      string `json:"order_id"`
	OrderDate    string `json:"order_date"`
	TransferType string `json:"transfer_type"`
	OutCustID    string `json:"out_cust_id"`
	OutAcctID    string `json:"out_ acct _id"`
	InCustID     string `json:"in_cust_id"`
	InAcctID     string `json:"in_ acct _id"`
	TransferAmt  string `json:"transfer_amt"`
	BgRetURL     string `json:"bg_ret_url"`
	MerPriv      string `json:"mer_priv"`
	Extension    string `json:"extension"`
}

func (this TransferReq) APICmId() string {
	return "203"
}

func (this TransferReq) Params() map[string]string {
	return nil
}

func (this TransferReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
