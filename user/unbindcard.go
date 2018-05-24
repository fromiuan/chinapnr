// 4.2.5.	银行卡解绑接口
package user

import (
	"reflect"
)

type UnBindCardReq struct {
	UserCustID   string `json:"user_cust_id"`   // 必须 用户客户号	由汇付生成，用户的唯一性标识
	OrderDate    string `json:"order_date"`     // 必须 订单日期	格式为YYYYMMDD，例如：20160307
	OrderID      string `json:"order_id"`       // 必须 订单号	由商户生成，必须保证唯一，20位内数字或字母的组合
	BankID       string `json:"bank_id"`        // 必须 银行代码	具体见附件二：快捷支付支持银行列表
	CardID       string `json:"card_id"`        // 必须 银行卡绑定ID	绑定银行卡ID
	CardMobile   string `json:"card_mobile"`    // 必须 银行卡预留手机号	银行卡预留手机号
	CardBussType string `json:"card_buss_type"` // 必须 卡业务类型	0:取现，1：代扣，2：快捷
	SmsOrderDate string `json:"sms_order_date"` // 必须 短信验证码日期	与sms_code对应的验证码发送订单日期
	SmsOrderID   string `json:"sms_order_id"`   // 必须 短信验证码订单号	与sms_code对应的验证码发送订单号
	SmsCode      string `json:"sms_code"`       // 必须 短信验证码	短信验证码
	BgRetURL     string `json:"bg_ret_url"`     // 必须 商户后台应答地址	通过后台异步通知商户解绑结果	注意： 1) 使用时不要包含中文	2) 必须是外网地址

	MerPriv   string `json:"mer_priv"`  // 可选 商户私有域	为商户的自定义字段，该字段在交易完成后由本平台原样返回
	Extension string `json:"extension"` // 可选 扩展域	用于扩展请求参数
}

type UnBindCardRsp struct {
	CmdID        string `json:"cmd_id"`
	RespCode     string `json:"resp_code"`
	RespDesc     string `json:"resp_desc"`
	MerCustID    string `json:"mer_cust_id"`
	UserCustID   string `json:"user_cust_id"`
	OrderID      string `json:"order_id"`
	OrderDate    string `json:"order_date"`
	BankID       string `json:"bank_id"`
	CardID       string `json:"card_id"`
	CardBussType string `json:"cardBussType"`
	CardMobile   string `json:"card_mobile"`
	SmsSeqID     string `json:"sms_seq_id"`
	SmsCode      string `json:"sms_code"`
	ProcessStat  string `json:"process_stat"`
	BgRetURL     string `json:"bg_ret_url"`
	MerPriv      string `json:"mer_priv"`
	Extension    string `json:"extension"`
}

func (this UnBindCardReq) APICmId() string {
	return "105"
}

func (this UnBindCardReq) Params() map[string]string {
	return nil
}

func (this UnBindCardReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
