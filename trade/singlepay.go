//单笔代扣(裸扣)支付接口
package trade

type SinglePayReq struct {
	Version    string `json:"version"`
	CmdID      string `json:"cmd_id"`
	MerCustID  string `json:"mer_cust_id"`
	OrderID    string `json:"order_id"`
	OrderDate  string `json:"order_date"`
	UserName   string `json:"user_name"`
	CertID     string `json:"cert_id"`
	UserMobile string `json:"user_mobile"`
	CardNo     string `json:"card_no"`
	AcctID     string `json:"acct_id"`
	TransAmt   string `json:"trans_amt"`
	TransType  string `json:"trans_type"`
	BgRetURL   string `json:"bg_ret_url"`
	MerPriv    string `json:"mer_priv"`
	Extension  string `json:"extension"`
}

type SinglePayRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	OrderID       string `json:"order_id"`
	OrderDate     string `json:"order_date"`
	PlatformSeqID string `json:"platform_seq_id"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
	AcctDate      string `json:"acct_date"`
	UserName      string `json:"user_name"`
	CertID        string `json:"cert_id"`
	UserMobile    string `json:"user_mobile"`
	CardNo        string `json:"card_no"`
	AcctID        string `json:"acct_id"`
	TransAmt      string `json:"trans_amt"`
	TransType     string `json:"trans_type"`
}
