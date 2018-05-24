//H5收银台
package trade

type H5Req struct {
	Version         string `json:"version"`
	CmdID           string `json:"cmd_id"`
	MerCustID       string `json:"mer_cust_id"`
	UserCustID      string `json:"user_cust_id"`
	OrderDate       string `json:"order_date"`
	OrderID         string `json:"order_id"`
	TransAmt        string `json:"trans_amt"`
	InCustID        string `json:"in_cust_id"`
	InAcctID        string `json:"in_acct_id"`
	DivDetail       string `json:"div_detail"`
	GoodsDesc       string `json:"goods_desc"`
	GoodsType       string `json:"goods_type"`
	OrderExpireTime string `json:"order_expire_time"`
	RetURL          string `json:"ret_url"`
	BgRetURL        string `json:"bg_ret_url"`
	MerPriv         string `json:"mer_priv"`
	Extension       string `json:"Extension"`
}

type H5Rsp struct {
	CmdID           string `json:"cmd_id"`
	RespCode        string `json:"resp_code"`
	RespDesc        string `json:"resp_desc"`
	PlatformSeqID   string `json:"platform_seq_id"`
	MerCustID       string `json:"mer_cust_id"`
	OrderDate       string `json:"order_date"`
	OrderID         string `json:"order_id"`
	UserCustID      string `json:"user_cust_id"`
	TokenID         string `json:"token_id"`
	PayInfo         string `json:"pay_info"`
	BankID          string `json:"bank_id"`
	CardNo          string `json:"card_no"`
	BindCardID      string `json:"bind_card_id"`
	DcFlag          string `json:"dc_flag"`
	TransAmt        string `json:"trans_amt"`
	IsRaw           string `json:"is_raw"`
	BuyerID         string `json:"buyer_id"`
	GoodsDesc       string `json:"goods_desc"`
	GoodsType       string `json:"goods_type"`
	AttachInfo      string `json:"attach_info"`
	OrderExpireTime string `json:"order_expire_time"`
	InCustID        string `json:"in_cust_id"`
	InAcctID        string `json:"in_acct_id"`
	DivDetail       string `json:"div_detail"`
	RetURL          string `json:"ret_url"`
	BgRetURL        string `json:"bg_ret_url"`
	MerPriv         string `json:"mer_priv"`
	Extension       string `json:"Extension"`
}
