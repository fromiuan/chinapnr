//快捷支付APP版接口
package trade

import (
	"reflect"

	"github.com/fromiuan/chinapnr/encoding"
)

type FasePayAppReq struct {
	UserCustID string `json:"user_cust_id"` // 可选  用户客户号
	OrderDate  string `json:"order_date"`   // 必须  订单日期
	OrderID    string `json:"order_id"`     // 必须  订单号
	TransAmt   string `json:"trans_amt"`    // 必须  交易金额
	BgRetURL   string `json:"bg_ret_url"`   // 必须  商户后台应答地址

	DivDetail      []FaseAppPayDivDetail `json:"-"`                // 可选  分账账户串
	InCustID       string                `json:"in_cust_id"`       // 可选  入账客户号
	InAcctID       string                `json:"in_acct_id"`       // 可选  入账账户号
	DeviceInfo     string                `json:"device_info"`      // 可选  设备号
	IPAddr         string                `json:"ip_addr"`          // 可选  IP地址
	LocationVal    string                `json:"location_val"`     // 可选  经纬度
	RetURL         string                `json:"ret_url"`          // 可选  页面返回URL
	MerPriv        string                `json:"mer_priv"`         // 可选  商户私有域
	Extension      string                `json:"extension"`        // 可选  扩展域
	SecondaryMerID string                `json:"secondary_mer_id"` // 可选  二级商户号
}

type FasePayAppRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	UserCustID    string `json:"user_cust_id"`
	AcctID        string `json:"acct_id"`
	OrderDate     string `json:"order_date"`
	OrderID       string `json:"order_id"`
	PlatformSeqID string `json:"platform_seq_id"`
	BankID        string `json:"bank_id"`
	CardNo        string `json:"card_no"`
	BindCardID    string `json:"bind_card_id"`
	DcFlag        string `json:"dc_flag"`
	TransAmt      string `json:"trans_amt"`
	DivDetail     string `json:"div_detail"`
	DivCustID     string `json:"divCustId"`
	DivAcctID     string `json:"divAcctId"`
	DivAmt        string `json:"divAmt"`
	DivFreezeFg   string `json:"divFreezeFg"`
	FeeAmt        string `json:"fee_amt"`
	FeeCustID     string `json:"fee_cust_id"`
	FeeAcctID     string `json:"fee_acct_id"`
	RetURL        string `json:"ret_url"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}

//div_detail
type FaseAppPayDivDetail struct {
	DivCustId   string `json:"divCustId"`
	DivAcctId   string `json:"divAcctId"`
	DivAmt      string `json:"divAmt"`
	DivFreezeFg string `json:"divFreezeFg"`
}

func (this FasePayAppReq) APICmId() string {
	return "208"
}

func (this FasePayAppReq) Params() map[string]string {
	if this.DivDetail == nil {
		return nil
	}
	data := make(map[string]string)
	divDetail, _ := encoding.Marshal(this.DivDetail)
	data["div_detail"] = string(divDetail)
	return data
}

func (this FasePayAppReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
