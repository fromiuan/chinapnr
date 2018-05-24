//代发接口
package trade

type IssuingReq struct {
	Version       string `json:"version"`
	CmdID         string `json:"cmd_id"`
	MerCustID     string `json:"mer_cust_id"`
	OrderDate     string `json:"order_date"`
	PayCustID     string `json:"pay_cust_id"`
	PayAcctID     string `json:"pay_acct_id"`
	OrderID       string `json:"order_id"`
	TransAmt      string `json:"trans_amt"`
	AcctType      string `json:"acct_type"`
	BankCardNo    string `json:"bank_card_no"`
	BankID        string `json:"bank_id"`
	BankName      string `json:"bank_name"`
	BankProv      string `json:"bank_prov"`
	BankArea      string `json:"bank_area"`
	BankSubbranch string `json:"bank_subbranch"`
	UniteBankCode string `json:"unite_bank_code"`
	Purpose       string `json:"purpose"`
	CustName      string `json:"cust_name"`
	Mobile        string `json:"mobile"`
	CertType      string `json:"cert_type"`
	CertID        string `json:"cert_id"`
	TransMode     string `json:"trans_mode"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}

type IssuingRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	OrderDate     string `json:"order_date"`
	OrderID       string `json:"order_id"`
	PlatformSeqID string `json:"platform_seq_id"`
	FeeAmt        string `json:"fee_amt"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}
