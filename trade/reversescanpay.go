//反扫消费接口
package trade

type ReverseScanPayReq struct {
	Version        string `json:"version"`
	CmdID          string `json:"cmd_id"`
	MerCustID      string `json:"mer_cust_id"`
	UserCustID     string `json:"user_cust_id"`
	OrderDate      string `json:"order_date"`
	OrderID        string `json:"order_id"`
	PayType        string `json:"pay_type"`
	TransAmt       string `json:"trans_amt"`
	DivDetail      string `json:"div_detail"`
	GoodsDesc      string `json:"goods_desc"`
	GoodsType      string `json:"goods_type"`
	AuthCode       string `json:"auth_code"`
	OperUserID     string `json:"oper_user_Id"`
	DeviceInfo     string `json:"device_info"`
	BgRetURL       string `json:"bg_ret_url"`
	MerPriv        string `json:"mer_priv"`
	Extension      string `json:"extension"`
	SecondaryMerID string `json:"secondary_mer_id"`
}

type ReverseScanPayRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	UserCustID    string `json:"user_cust_id"`
	OrderDate     string `json:"order_date"`
	OrderID       string `json:"order_id"`
	PayType       string `json:"pay_type"`
	TransAmt      string `json:"trans_amt"`
	DivDetail     string `json:"div_detail"`
	GoodsDesc     string `json:"goods_desc"`
	GoodsType     string `json:"goods_type"`
	AuthCode      string `json:"auth_code"`
	OperUserID    string `json:"oper_user_id"`
	DeviceInfo    string `json:"device_info"`
	PlatformSeqID string `json:"platform_seq_id"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}
