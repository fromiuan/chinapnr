//	交易状态查询接口
package query

import (
	"reflect"
)

type TradeStatusReq struct {
	OrderID   string `json:"order_id"`   // 必须 订单号
	OrderDate string `json:"order_date"` // 必须 订单日期
	TransType string `json:"trans_type"` // 必须 交易查询类型

	MerPriv   string `json:"mer_priv"`  // 可选 商户私有域
	Extension string `json:"extension"` // 可选 扩展域
}

type TradeStatusRsp struct {
	CmdID         string `json:"cmd_id"`
	MerCustID     string `json:"mer_cust_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	OrderID       string `json:"order_id"`
	OrderDate     string `json:"order_date"`
	TransType     string `json:"trans_type"`
	TransStat     string `json:"trans_stat"`
	TransAmt      string `json:"trans_amt"`
	PlatformSeqID string `json:"platform_seq_id"`
	DivDetail     string `json:"div_detail"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}

func (this TradeStatusReq) APICmId() string {
	return "301"
}

func (this TradeStatusReq) Params() map[string]string {
	return nil
}

func (this TradeStatusReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
