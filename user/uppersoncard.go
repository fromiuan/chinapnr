// 个体户证照上传接口
package user

import (
	"reflect"
)

type UpPersionCardReq struct {
	UserCustID   string `json:"user_cust_id"`
	OrderID      string `json:"order_id"`
	OrderDate    string `json:"order_date"`
	OperateType  string `json:"operate_type"`
	BusinessCode string `json:"business_code"`
	BgRetURL     string `json:"bg_ret_url"`
	RetURL       string `json:"ret_url"`
	MerPriv      string `json:"mer_priv"`
	Extension    string `json:"extension"`
}

type UpPersionCardRsp struct {
	CmdID        string `json:"cmd_id"`
	RespCode     string `json:"resp_code"`
	RespDesc     string `json:"resp_desc"`
	MerCustID    string `json:"mer_cust_id"`
	OrderID      string `json:"order_id"`
	OrderDate    string `json:"order_date"`
	UserCustID   string `json:"user_cust_id"`
	BusinessCode string `json:"business_code"`
	BgRetURL     string `json:"bg_ret_url"`
	OperateType  string `json:"operate_type"`
	RetURL       string `json:"ret_url"`
	MerPriv      string `json:"mer_priv"`
	Extension    string `json:"extension"`
}

func (this UpPersionCardReq) APICmId() string {
	return "107"
}

func (this UpPersionCardReq) Params() map[string]string {
	return nil
}

func (this UpPersionCardReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
