//	一阶段绑卡签约短信发送接口
package user

import (
	"reflect"
)

type SetpSendMsgReq struct {
	UserCustID     string `json:"user_cust_id"`     // 必须 用户客户号	由汇付生成，用户的唯一性标识
	OrderID        string `json:"order_id"`         // 必须 订单号	由商户生成，必须保证唯一，20位内的字母或数字组合
	OrderDate      string `json:"order_date"`       // 必须 订单日期	格式为YYYYMMDD，例如：20160307
	CardVerifyType string `json:"card_verify_type"` // 必须 验卡类型	快捷绑卡： 02	消费分期代扣绑卡: 03	非消费分期代扣绑卡: 04
	DcFlag         string `json:"dc_flag"`          // 必须 借贷标记	具体见附件：开户银行代号	0--借记，储蓄卡	1--贷记，信用卡
	CardNo         string `json:"card_no"`          // 必须 银行卡号	本次快捷绑卡待绑定的银行卡号
	CardMobile     string `json:"card_mobile"`      // 必须 银行预留手机号	本次快捷绑卡待绑定的银行卡，在银行开户时的预留手机号
	BgRetURL       string `json:"bg_ret_url"`       // 必须 商户后台应答地址	通过后台异步通知商户短信发送结果	1) 使用时不要包含中文	2) 必须是外网地址

	BankID    string `json:"bank_id"`   // 可输 开户银行代号	快捷绑卡： 02	消费分期代扣绑卡: 03	非消费分期代扣绑卡: 04
	CardProv  string `json:"card_prov"` // 可输 银行卡开户省份	本次快捷绑卡待绑定的银行卡开户省份
	CardArea  string `json:"card_area"` // 可输 银行卡开户地区	本次快捷绑卡待绑定的银行卡开户地区
	MerPriv   string `json:"mer_priv"`  // 可选 商户私有域	为商户的自定义字段，该字段在交易完成后由本平台原样返回
	Extension string `json:"extension"` // 可选 扩展域	用于扩展请求参数
}

type SetpSendMsgRsp struct {
	CmdID          string `json:"cmd_id"`
	RespCode       string `json:"resp_code"`
	RespDesc       string `json:"resp_desc"`
	MerCustID      string `json:"mer_cust_id"`
	UserCustID     string `json:"user_cust_id"`
	OrderID        string `json:"order_id"`
	OrderDate      string `json:"order_date"`
	CardVerifyType string `json:"card_verify_type"`
	BgRetURL       string `json:"bg_ret_url"`
	MerPriv        string `json:"mer_priv"`
	Extension      string `json:"extension"`
}

func (this SetpSendMsgReq) APICmId() string {
	return "111"
}

func (this SetpSendMsgReq) Params() map[string]string {
	return nil
}

func (this SetpSendMsgReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
