//快捷支付接口
package trade

import (
	"reflect"

	"github.com/fromiuan/chinapnr/encoding"
)

type FasePayReq struct {
	UserCustID  string          `json:"user_cust_id"` // 必须  用户客户号
	OrderDate   string          `json:"order_date"`   // 必须  订单日期 请和一阶段快捷支付短信发送接口的订单日期一致
	OrderID     string          `json:"order_id"`     // 必须  订单号 请和一阶段快捷支付短信发送接口的订单号一致
	TransAmt    string          `json:"trans_amt"`    // 必须  交易金额 金额格式是###.00
	BindCardID  string          `json:"bind_card_id"` // 必须  绑定银行卡ID
	DivDetail   []FastDivDetail `json:"-"`            // 必须  分账账户串
	DivCustID   string          `json:"divCustId"`    // 必须  分账客户号
	DivAcctID   string          `json:"divAcctId"`    // 必须  分账账户号
	DivAmt      string          `json:"divAmt"`       // 必须  分账金额
	DivFreezeFg string          `json:"divFreezeFg"`  // 必须  是否冻结标志
	BgRetURL    string          `json:"bg_ret_url"`   // 必须  商户后台应答地址

	SmsCode        string `json:"sms_code"`         // 可选  短信验证码
	RetURL         string `json:"ret_url"`          // 可选  页面返回URL
	MerPriv        string `json:"mer_priv"`         // 可选  商户私有域
	Extension      string `json:"extension"`        // 可选  扩展域
	SecondaryMerID string `json:"secondary_mer_id"` // 可选  二级商户号
}

type FasePayRsp struct {
	CmdID         string          `json:"cmd_id"`
	RespCode      string          `json:"resp_code"`
	RespDesc      string          `json:"resp_desc"`
	MerCustID     string          `json:"mer_cust_id"`
	UserCustID    string          `json:"user_cust_id"`
	OrderDate     string          `json:"order_date"`
	OrderID       string          `json:"order_id"`
	PlatformSeqID string          `json:"platform_seq_id"`
	TransAmt      string          `json:"trans_amt"`
	BindCardID    string          `json:"bind_card_id"`
	SmsCode       string          `json:"sms_code"`
	DivDetail     []FastDivDetail `json:"div_detail"`
	DivCustID     string          `json:"divCustId"`
	DivAcctID     string          `json:"divAcctId"`
	DivAmt        string          `json:"divAmt"`
	DivFreezeFg   string          `json:"divFreezeFg"`
	FeeAmt        string          `json:"fee_amt"`
	FeeCustID     string          `json:"fee_cust_id"`
	FeeAcctID     string          `json:"fee_acct_id"`
	RetURL        string          `json:"ret_url"`
	BgRetURL      string          `json:"bg_ret_url"`
	MerPriv       string          `json:"mer_priv"`
	Extension     string          `json:"extension"`
}

type FastDivDetail struct {
	DivCustID   string `json:"divCustId"`
	DivAcctID   string `json:"divAcctId"`
	DivAmt      string `json:"divAmt"`
	DivFreezeFg string `json:"divFreezeFg"`
}

func (this FasePayReq) APICmId() string {
	return "201"
}

func (this FasePayReq) Params() map[string]string {
	data := make(map[string]string)
	divDetail, _ := encoding.Marshal(this.DivDetail)
	data["div_detail"] = string(divDetail)
	return data
}

func (this FasePayReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
