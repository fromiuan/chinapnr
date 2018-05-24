//	个体户信息补录接口
package user

import (
	"reflect"
)

type UpExtPersonCardReq struct {
	UserCustID  string `json:"user_cust_id"`
	CustAddress string `json:"cust_address"`
	Occupation  string `json:"occupation"`
	MerPriv     string `json:"mer_priv"`
	Extension   string `json:"extension"`
}

type UpExtPersonCardRsp struct {
	CmdID      string `json:"cmd_id"`
	RespCode   string `json:"resp_code"`
	RespDesc   string `json:"resp_desc"`
	MerCustID  string `json:"mer_cust_id"`
	UserCustID string `json:"user_cust_id"`
	MerPriv    string `json:"mer_priv"`
	Extension  string `json:"extension"`
}

func (this UpExtPersonCardReq) APICmId() string {
	return "109"
}

func (this UpExtPersonCardReq) Params() map[string]string {
	return nil
}

func (this UpExtPersonCardReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
