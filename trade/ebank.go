//网银支付接口
package trade

import (
	"reflect"
)

type EbankReq struct {
	Version        string      `json:"version"`
	CmdID          string      `json:"cmd_id"`
	MerCustID      string      `json:"mer_cust_id"`
	UserCustID     string      `json:"user_cust_id"`
	OrderID        string      `json:"order_id"`
	OrderDate      string      `json:"order_date"`
	GateID         string      `json:"gate_id"`
	TransAmt       string      `json:"trans_amt"`
	RetURL         string      `json:"ret_url"`
	BgRetURL       string      `json:"bg_ret_url"`
	MerPriv        string      `json:"mer_priv"`
	Extension      string      `json:"extension"`
	DivDetail      []DivDetail `json:"-"`
	SecondaryMerID string      `json:"secondary_mer_id"`
}

type EbankRsp struct {
	CmdID         string      `json:"cmd_id"`
	RespCode      string      `json:"resp_code"`
	RespDesc      string      `json:"resp_desc"`
	MerCustID     string      `json:"mer_cust_id"`
	UserCustID    string      `json:"user_cust_id"`
	OrderID       string      `json:"order_id"`
	OrderDate     string      `json:"order_date"`
	TransAmt      string      `json:"trans_amt"`
	PlatformSeqID string      `json:"platform_seq_id"`
	BgRetURL      string      `json:"bg_ret_url"`
	RetURL        string      `json:"ret_url"`
	MerPriv       string      `json:"mer_priv"`
	Extension     string      `json:"extension"`
	DivDetail     []DivDetail `json:"div_detail"`
}

type DivDetail struct {
	DivCustId string `json:"divCustId"`
	DivAcctId string `json:"divAcctId"`
	DivAmt    string `json:"divAmt"`
}

func (this EbankReq) APICmId() string {
	return "204"
}

func (this EbankReq) Params() map[string]string {
	return nil
}

func (this EbankReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
