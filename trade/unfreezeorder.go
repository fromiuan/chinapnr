//订单解冻接口
package trade

import (
	"github.com/fromiuan/chinapnr/encoding"
	"reflect"
)

type UnfreezeOrderReq struct {
	Version              string      `json:"version"`
	CmdID                string      `json:"cmd_id"`
	MerCustID            string      `json:"mer_cust_id"`
	UserCustID           string      `json:"user_cust_id"`
	OrderDate            string      `json:"order_date"`
	OrderID              string      `json:"order_id"`
	OrginalPlatformSeqID string      `json:"orginal_platform_seq_id"`
	DivDetail            []DivDetail `json:"-"`
	BgRetURL             string      `json:"bg_ret_url"`
	MerPriv              string      `json:"mer_priv"`
	Extension            string      `json:"extension"`
}

type UnfreezeOrderRsp struct {
	CmdID         string      `json:"cmd_id"`
	RespCode      string      `json:"resp_code"`
	RespDesc      string      `json:"resp_desc"`
	MerCustID     string      `json:"mer_cust_id"`
	UserCustID    string      `json:"user_cust_id"`
	OrderDate     string      `json:"order_date"`
	OrderID       string      `json:"order_id"`
	PlatformSeqID string      `json:"platform_seq_id"`
	SuccessCnt    string      `json:"success_cnt"`
	SuccessAmt    string      `json:"success_amt"`
	FailCnt       string      `json:"fail_cnt"`
	FailAmt       string      `json:"fail_amt"`
	FailDivDetail []DivDetail `json:"fail_div_detail"`
	BgRetURL      string      `json:"bg_ret_url"`
	MerPriv       string      `json:"mer_priv"`
	Extension     string      `json:"extension"`
}

func (this UnfreezeOrderReq) APICmId() string {
	return "212"
}

func (this UnfreezeOrderReq) Params() map[string]string {
	data := make(map[string]string)
	divdetail, err := encoding.Marshal(this.DivDetail)
	if err == nil {
		return data
	}
	data["div_detail"] = string(divdetail)
	return nil
}

func (this UnfreezeOrderReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
