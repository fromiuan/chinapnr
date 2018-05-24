//扫码支付接口
package trade

type ScanPayReq struct {
	Version         string `json:"version"`
	CmdID           string `json:"cmd_id"`
	MerCustID       string `json:"mer_cust_id"`
	UserCustID      string `json:"user_cust_id"`
	OrderDate       string `json:"order_date"`
	OrderID         string `json:"order_id"`
	GoodsDesc       string `json:"goods_desc"`
	PayType         string `json:"pay_type"`
	RequestType     string `json:"request_type"`
	TransAmt        string `json:"trans_amt"`
	DivDetail       string `json:"div_detail"`
	DivCustID       string `json:"divCustId"`
	DivAcctID       string `json:"divAcctId"`
	DivAmt          string `json:"divAmt"`
	DivFreezeFg     string `json:"divFreezeFg"`
	OperUserID      string `json:"oper_user_Id"`
	GoodsType       string `json:"goods_type"`
	OrderExpireTime string `json:"order_expire_time"`
	DeviceInfo      string `json:"device_info"`
	RetURL          string `json:"ret_url"`
	BgRetURL        string `json:"bg_ret_url"`
	MerPriv         string `json:"mer_priv"`
	Extension       string `json:"extension"`
	SecondaryMerID  string `json:"secondary_mer_id"`
}

type ScanPayRsp struct {
	CmdID           string `json:"cmd_id"`
	RespCode        string `json:"resp_code"`
	RespDesc        string `json:"resp_desc"`
	MerCustID       string `json:"mer_cust_id"`
	UserCustID      string `json:"user_cust_id"`
	OrderDate       string `json:"order_date"`
	OrderID         string `json:"order_id"`
	PlatformSeqID   string `json:"platform_seq_id"`
	GoodsDesc       string `json:"goods_desc"`
	PayType         string `json:"pay_type"`
	RequestType     string `json:"request_type"`
	TransAmt        string `json:"trans_amt"`
	DivDetail       string `json:"div_detail"`
	DivCustID       string `json:"divCustId"`
	DivAcctID       string `json:"divAcctId"`
	DivAmt          string `json:"divAmt"`
	DivFreezeFg     string `json:"divFreezeFg"`
	QrcodeURL       string `json:"qrcode_url"`
	OperUserID      string `json:"oper_user_id"`
	GoodsType       string `json:"goods_type"`
	OrderExpireTime string `json:"order_expire_time"`
	DeviceInfo      string `json:"device_info"`
	RetURL          string `json:"ret_url"`
	BgRetURL        string `json:"bg_ret_url"`
	MerPriv         string `json:"mer_priv"`
	Extension       string `json:"extension"`
}
