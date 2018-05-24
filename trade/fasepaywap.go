//快捷支付WEB版接口
package trade

type FasePayWayReq struct {
	Version        string             `json:"version"`
	CmdID          string             `json:"cmd_id"`
	MerCustID      string             `json:"mer_cust_id"`
	UserCustID     string             `json:"user_cust_id"`
	OrderDate      string             `json:"order_date"`
	OrderID        string             `json:"order_id"`
	TransAmt       string             `json:"trans_amt"`
	DivDetail      []FasePayDivDetail `json:"-"`
	InCustID       string             `json:"in_cust_id"`
	InAcctID       string             `json:"in_acct_id"`
	DeviceInfo     string             `json:"device_info"`
	IPAddr         string             `json:"ip_addr"`
	LocationVal    string             `json:"location_val"`
	RetURL         string             `json:"ret_url"`
	BgRetURL       string             `json:"bg_ret_url"`
	MerPriv        string             `json:"mer_priv"`
	Extension      string             `json:"extension"`
	SecondaryMerID string             `json:"secondary_mer_id"`
}

type FasePayWayRsp struct {
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
	TransAmt      string `json:"trans_amt"`
	DcFlag        string `json:"dc_flag"`
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
type FasePayDivDetail struct {
	DivCustId   string `json:"divCustId"`
	DivAcctId   string `json:"divAcctId"`
	DivAmt      string `json:"divAmt"`
	DivFreezeFg string `json:"divFreezeFg"`
}
