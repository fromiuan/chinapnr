//	快捷支付统合版一阶段
package user

type FastUnderunifyReq struct {
	Version    string      `json:"version"`
	CmdID      string      `json:"cmd_id"`
	MerCustID  string      `json:"mer_cust_id"`
	UserCustID string      `json:"user_cust_id"`
	OrderDate  string      `json:"order_date"`
	OrderID    string      `json:"order_id"`
	TransAmt   string      `json:"trans_amt"`
	DivDetail  []DivDetail `json:"-"`
	InCustID   string      `json:"in_cust_id"`
	InAcctID   string      `json:"in_acct_id"`
	CardNo     string      `json:"card_no"`
	CardMobile string      `json:"card_ mobile"`
	CardProv   string      `json:"card_prov"`
	CardArea   string      `json:"card_area"`
	BgRetURL   string      `json:"bg_ret_url"`
	MerPriv    string      `json:"mer_priv"`
	Extension  string      `json:"extension"`
}

// div_detail
type DivDetail struct {
	DivCustID   string `json:"divCustId"`
	DivAcctID   string `json:"divAcctId"`
	DivAmt      string `json:"divAmt"`
	DivFreezeFg string `json:"divFreezeFg"`
}

type FastUnderunifyRsp struct {
	CmdID      string `json:"cmd_id"`
	RespCode   string `json:"resp_code"`
	RespDesc   string `json:"resp_desc"`
	MerCustID  string `json:"mer_cust_id"`
	UserCustID string `json:"user_cust_id"`
	OrderDate  string `json:"order_date"`
	OrderID    string `json:"order_id"`
	BgRetURL   string `json:"bg_ret_url"`
	MerPriv    string `json:"mer_priv"`
	Extension  string `json:"extension"`
}
