//	卡bin查询接口
package query

import (
	"reflect"
)

type CardBinReq struct {
	CardNo string `json:"card_no"`
}

type CardBinRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	PlatformSeqID string `json:"platform_seq_id"`
	CardNo        string `json:"card_no"`
	CardBin       string `json:"card_bin"`
	BankName      string `json:"bank_name"`
	CardType      string `json:"card_type"`
	BankNo        string `json:"bank_no"`
}

func (this CardBinReq) APICmId() string {
	return "302"
}

func (this CardBinReq) Params() map[string]string {
	return nil
}

func (this CardBinReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
