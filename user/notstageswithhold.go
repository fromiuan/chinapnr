//	非消费分期代扣签约绑卡接口
package user

type NotStageWithholdReq struct {
	UserCustID string `json:"user_cust_id"`
	OrderDate  string `json:"order_date"`
	OrderID    string `json:"order_id"`
	BankID     string `json:"bank_id"`
	DcFlag     string `json:"dc_flag"`
	CardNo     string `json:"card_no"`
	CardMobile string `json:"card_mobile"`
	CardProv   string `json:"card_prov"`
	CardArea   string `json:"card_area"`
	SmsCode    string `json:"sms_code"`
	BgRetURL   string `json:"bg_ret_url"`
	RetURL     string `json:"ret_url"`
	MerPriv    string `json:"mer_priv"`
	Extension  string `json:"extension"`
}

type NotStageWithholdRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	UserCustID    string `json:"user_cust_id"`
	OrderDate     string `json:"order_date"`
	OrderID       string `json:"order_id"`
	PlatformSeqID string `json:"platform_seq_id"`
	BankID        string `json:"bank_id"`
	CardNo        string `json:"card_no"`
	BindCardID    string `json:"bind_card_id"`
	BgRetURL      string `json:"bg_ret_url"`
	RetURL        string `json:"ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}
