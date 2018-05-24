//	裸扣交易状态查询接口
package query

type BuckleStatusReq struct {
	OrderID    string `json:"order_id"`
	MerBatchID string `json:"mer_batch_id"`
	TransDate  string `json:"trans_date"`
	MerPriv    string `json:"mer_priv"`
	Extension  string `json:"extension"`
}

type BuckleStatusRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	MerBatchID    string `json:"mer_batch_id"`
	OrderID       string `json:"order_id"`
	OrderDate     string `json:"order_date"`
	UserName      string `json:"user_name"`
	CertID        string `json:"cert_id"`
	TransAmt      string `json:"trans_amt"`
	CardNo        string `json:"card_no"`
	UserMobile    string `json:"user_mobile"`
	AcctDate      string `json:"acct_date"`
	PlatformSeqID string `json:"platform_seq_id"`
	TransType     string `json:"trans_type"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}
