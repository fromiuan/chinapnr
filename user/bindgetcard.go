// 4.2.4.	绑定取现卡接口
package user

import (
	"reflect"
)

type BindGetCardReq struct {
	UserCustID string `json:"user_cust_id"` // 必须  用户唯一标识号 由汇付生成，用户的唯一性标识
	OrderDate  string `json:"order_date"`   // 必须  订单日期 格式为“yyyyMMdd”
	OrderID    string `json:"order_id"`     // 必须  订单号
	BankID     string `json:"bank_id"`      // 必须  银行代号
	CardNo     string `json:"card_no"`      // 必须  银行卡号
	DcFlag     string `json:"dc_flag"`      // 必须  借贷标记 0：借记卡（固定）
	CardProv   string `json:"card_prov"`    // 必须  银行卡开户省份
	CardArea   string `json:"card_area"`    // 必须  银行卡开户地区
	BgRetURL   string `json:"bg_ret_url"`   // 必须  后台返回地址 通过后台异步通知商户绑卡结果	注意： 1) 使用时不要包含中文	2) 必须是外网地址

	MerPriv   string `json:"mer_priv"`  // 可选  商户私有域 为商户的自定义字段，该字段在交易完成后由本平台原样返回
	Extension string `json:"extension"` // 可选  扩展域 用于扩展请求参数
}

type BindGetCardRsp struct {
	CmdID          string `json:"cmd_id"`
	RespCode       string `json:"resp_code"`
	RespDesc       string `json:"resp_desc"`
	MerCustID      string `json:"mer_cust_id"`
	UserCustID     string `json:"user_cust_id"`
	OrderDate      string `json:"order_date"`
	OrderID        string `json:"order_id"`
	PlatformSeqID  string `json:"platform_seq_id"`
	BankID         string `json:"bank_id"`
	CardNo         string `json:"card_no"`
	CashBindCardID string `json:"cash_bind_card_id"`
	DcFlag         string `json:"dc_flag"`
	CardProv       string `json:"card_prov"`
	CardArea       string `json:"card_area"`
	BgRetURL       string `json:"bg_ret_url"`
	MerPriv        string `json:"mer_priv"`
	Extension      string `json:"extension"`
}

func (this BindGetCardReq) APICmId() string {
	return "104"
}

func (this BindGetCardReq) Params() map[string]string {
	return nil
}

func (this BindGetCardReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
