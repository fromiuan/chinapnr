//批量代扣(裸扣)支付接口
package trade

type BatchPayReq struct {
	Version     string `json:"version"`
	CmdID       string `json:"cmd_id"`
	MerCustID   string `json:"mer_cust_id"`
	MerBatchID  string `json:"mer_batch_id"`
	BatchDate   string `json:"batch_date"`
	TransType   string `json:"trans_type"`
	AcctID      string `json:"acct_id"`
	TransDetail string `json:"trans_detail"`
	BgRetURL    string `json:"bg_ret_url"`
	MerPriv     string `json:"mer_priv"`
	Extension   string `json:"extension"`
	UserName    string `json:"user_name"`
	CertID      string `json:"cert_id"`
	UserMobile  string `json:"user_mobile"`
	CardNo      string `json:"card_no"`
	TransAmt    string `json:"trans_amt"`
	OrderID     string `json:"order_id"`
}

type BatchPayRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	MerBatchID    string `json:"mer_batch_id"`
	BatchDate     string `json:"batch_date"`
	TransType     string `json:"trans_type"`
	TransDetail   string `json:"trans_detail"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
	OrderID       string `json:"order_id"`
	TransStatus   string `json:"trans_status"`
	TransDesc     string `json:"trans_desc"`
	AcctDate      string `json:"acct_date"`
	UserName      string `json:"user_name"`
	CertID        string `json:"cert_id"`
	UserMobile    string `json:"user_mobile"`
	CardNo        string `json:"card_no"`
	AcctID        string `json:"acct_id"`
	TransAmt      string `json:"trans_amt"`
	PlatformSeqID string `json:"platform_seq_id"`
}
