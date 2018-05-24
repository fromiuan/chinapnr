// 4.2.6.	独立验卡接口
package user

import (
	"reflect"
)

type VeriCardReq struct {
	OrderID   string `json:"order_id"`
	OrderDate string `json:"order_date"`
	CertID    string `json:"cert_id"`
	CardMp    string `json:"card_mp"`
	DcFlag    string `json:"dc_flag"`
	CardNo    string `json:"card_no"`
	CardName  string `json:"card_name"`
	BankID    string `json:"bank_id"`
	CheckType string `json:"check_type"`
	BgRetURL  string `json:"bg_ret_url"`
	MerPriv   string `json:"mer_priv"`
	Extension string `json:"extension"`
}

type VeriCardRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	OrderID       string `json:"order_id"`
	OrderDate     string `json:"order_date"`
	PlatformSeqID string `json:"platform_seq_id"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}

func (this VeriCardReq) APICmId() string {
	return "106"
}

func (this VeriCardReq) Params() map[string]string {
	return nil
}

func (this VeriCardReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
