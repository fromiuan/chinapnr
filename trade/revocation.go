//反扫撤销接口
package trade

type RevocationReq struct {
	Version              string `json:"version"`
	CmdID                string `json:"cmd_id"`
	MerCustID            string `json:"mer_cust_id"`
	OrderDate            string `json:"order_date"`
	OrderID              string `json:"order_id"`
	OrginalPlatformSeqID string `json:"orginal_platform_seq_id"`
	Remark               string `json:"remark"`
	BgRetURL             string `json:"bg_ret_url"`
	MerPriv              string `json:"mer_priv"`
	Extension            string `json:"extension"`
}

type RevocationRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	OrderDate     string `json:"order_date"`
	OrderID       string `json:"order_id"`
	PlatformSeqID string `json:"platform_seq_id"`
	Remark        string `json:"remark"`
	BgRetURL      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}
