// 4.2.3.	快捷卡绑卡接口
package user

import (
	"reflect"
)

type BindCardReq struct {
	UserCustId string `json:"user_cust_id"`
	OrderId    string `json:"order_id"`
	OrderDate  string `json:"order_date"`
	BankId     string `json:"bank_id"`
	DcFlag     string `json:"dc_flag"`
	CardNo     string `json:"card_no"`
	CardMobike string `json:"card_mobile"`
	CardProv   string `json:"card_prov"`
	CardArea   string `json:"card_area"`
	SmsCode    string `json:"sms_code"`
	BgRetUrl   string `json:"bg_ret_url"`
	RetUrl     string `json:"ret_url"`
	MerPriv    string `json:"mer_priv"`
	Extension  string `json:"extension"`
}

type BindCardRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	UserCustID    string `json:"user_cust_id"`
	OrderID       string `json:"order_id"`
	OrderDate     string `json:"order_date"`
	PlatformSeqID string `json:"platform_seq_id"`
	BindCardID    string `json:"bind_card_id"`
	RetURL        string `json:"ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}

func (this BindCardReq) APICmId() string {
	return "103"
}

func (this BindCardReq) Params() map[string]string {
	return nil
}

func (this BindCardReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
