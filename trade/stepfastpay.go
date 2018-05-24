//快捷支付统合版二阶段
package trade

type StepFastPayReq struct {
	Version        string `json:"version"`
	CmdID          string `json:"cmd_id"`
	MerCustID      string `json:"mer_cust_id"`
	UserCustID     string `json:"user_cust_id"`
	OrderDate      string `json:"order_date"`
	OrderID        string `json:"order_id"`
	SmsCode        string `json:"sms_code"`
	RetURL         string `json:"ret_url"`
	BgRetURL       string `json:"bg_ret_url"`
	MerPriv        string `json:"mer_priv"`
	Extension      string `json:"extension"`
	RequestType    string `json:"request_type"`
	SecondaryMerID string `json:"secondary_mer_id"`
}

type StepFastPayRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	UserCustID    string `json:"user_cust_id"`
	OrderDate     string `json:"order_date"`
	OrderID       string `json:"order_id"`
	PlatformSeqID string `json:"platform_seq_id"`
	SmsCode       string `json:"sms_code"`
	InCustID      string `json:"in_cust_id"`
	InAcctID      string `json:"in_acct_id"`
	BankID        string `json:"bank_id"`
	CardNo        string `json:"card_no"`
	BindCardID    string `json:"bind_card_id"`
	TransAmt      string `json:"trans_amt"`
	DivDetail     string `json:"div_detail"`
	FeeAmt        string `json:"fee_amt"`
	FeeCustID     string `json:"fee_cust_id"`
	FeeAcctID     string `json:"fee_acct_id"`
	RetURL        string `json:"ret_url"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}
