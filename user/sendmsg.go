//	短信发送接口
package user

import (
	"reflect"
)

type SendMsgReq struct {
	UserCustID   string `json:"user_cust_id"`  // 用户客户号
	OrderID      string `json:"order_id"`      // 订单号
	OrderDate    string `json:"order_date"`    // 订单日期
	BusinessType string `json:"business_type"` // 业务类型
	UserMobile   string `json:"user_mobile"`   // 用户手机号
}

type SendMsgRsp struct {
	CmdID        string `json:"cmd_id"`
	RespCode     string `json:"resp_code"`
	RespDesc     string `json:"resp_desc"`
	MerCustID    string `json:"mer_cust_id"`
	UserCustID   string `json:"user_cust_id"`
	OrderID      string `json:"order_id"`
	OrderDate    string `json:"order_date"`
	BusinessType string `json:"business_type"`
	UserMobile   string `json:"user_mobile"`
}

func (this SendMsgReq) APICmId() string {
	return "110"
}

func (this SendMsgReq) ExtParams() map[string]string {
	return nil
}

func (this SendMsgReq) Params() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
