//	快捷绑卡查询接口
package query

type FaseBindCardReq struct {
	UserCustId string `json:"user_cust_id"`
}

type FaseBindCardRsp struct {
	CmdID     string     `json:"cmd_id"`
	RespCode  string     `json:"resp_code"`
	RespDesc  string     `json:"resp_desc"`
	MerCustID string     `json:"mer_cust_id"`
	Counts    string     `json:"counts"`
	CardList  []CardList `json:"-"`
}

type CardList struct {
	BindCardID string `json:"bindCardId"`
	CardNo     string `json:"cardNo"`
	BankID     string `json:"bankId"`
	DcFlag     string `json:"dcFlag"`
	Status     string `json:"status"`
}
