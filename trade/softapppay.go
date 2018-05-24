//APP支付（支付宝、微信、银联各种支付）
package trade

import (
	"reflect"
)

type SoftAppPayReq struct {
	OrderDate string `json:"order_date"` // 必须 订单日期
	OrderID   string `json:"order_id"`   // 必须 订单号
	PayType   string `json:"pay_type"`   // 必须 支付类型
	TransAmt  string `json:"trans_amt"`  // 必须 交易金额
	GoodsDesc string `json:"goods_desc"` // 必须 商品描述
	BgRetURL  string `json:"bg_ret_url"` // 必须 商户后台应答地址

	UserCustID      string `json:"user_cust_id"`      // 可选 用户客户号
	InCustID        string `json:"in_cust_id"`        // 可选 入账客户号
	InAcctID        string `json:"in_acct_id"`        // 可选 入账账户号
	DivDetail       string `json:"div_detail"`        // 可选 分账串
	IsRaw           string `json:"is_raw"`            // 可选 是否原生态
	AppID           string `json:"app_id"`            // 可选 APPID
	BuyerID         string `json:"buyer_id"`          // 可选 买家用户ID
	GoodsType       string `json:"goods_type"`        // 可选 商品类型
	AttachInfo      string `json:"attach_info"`       // 可选 附加信可选息
	GoodTag         string `json:"good_tag"`          // 可选 商品标记
	OrderExpireTime string `json:"order_expire_time"` // 可选 订单超时时间
	DeviceInfo      string `json:"device_info"`       // 可选 设备号
	IPAddr          string `json:"ip_addr"`           // 可选 IP地址
	LocationVal     string `json:"location_val"`      // 可选 经纬度
	RetURL          string `json:"ret_url"`           // 可选 支付返回地址
	MerPriv         string `json:"mer_priv"`          // 可选 商户私有域
	Extension       string `json:"extension"`         // 可选 扩展域
	SecondaryMerID  string `json:"secondary_mer_id"`  // 可选 二级商户号
}

type SoftAppPayRsp struct {
	RespCode        string `json:"resp_code"`
	RespDesc        string `json:"resp_desc"`
	PlatformSeqID   string `json:"platform_seq_id"`
	PayURL          string `json:"pay_url"`
	TokenID         string `json:"token_id"`
	PayInfo         string `json:"pay_info"`
	OrderDate       string `json:"order_date"`
	OrderID         string `json:"order_id"`
	CmdID           string `json:"cmd_id"`
	MerCustID       string `json:"mer_cust_id"`
	UserCustID      string `json:"user_cust_id"`
	PayType         string `json:"pay_type"`
	TransAmt        string `json:"trans_amt"`
	InCustID        string `json:"in_cust_id"`
	InAcctID        string `json:"in_acct_id"`
	DivDetail       string `json:"div_detail"`
	IsRaw           string `json:"is_raw"`
	AppID           string `json:"app_id"`
	BuyerID         string `json:"buyer_id"`
	GoodsDesc       string `json:"goods_desc"`
	GoodsType       string `json:"goods_type"`
	AttachInfo      string `json:"attach_info"`
	GoodTag         string `json:"good_tag"`
	OrderExpireTime string `json:"order_expire_time"`
	DeviceInfo      string `json:"device_info"`
	IPAddr          string `json:"ip_addr"`
	LocationVal     string `json:"location_val"`
	RetURL          string `json:"ret_url"`
	BgRetURL        string `json:"bg_ret_url"`
	MerPriv         string `json:"mer_priv"`
	Extension       string `json:"extension"`
}

type SoftAppPayDetail struct {
	DivCustId   string `json:"divCustId"`
	DivAcctId   string `json:"divAcctId"`
	DivAmt      string `json:"divAmt"`
	DivFreezeFg string `json:"divFreezeFg"`
}

func (this SoftAppPayReq) APICmId() string {
	return "218"
}

func (this SoftAppPayReq) Params() map[string]string {
	return nil
}

func (this SoftAppPayReq) ExtParams() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
