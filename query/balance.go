//	余额查询接口
package query

type BalanceReq struct {
	UserCustId string `json:"user_cust_id"`
	AcctId     string `json:"acct_id"`
}

type BalanceRsp struct {
	CmdID         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RespDesc      string `json:"resp_desc"`
	MerCustID     string `json:"mer_cust_id"`
	UserCustID    string `json:"user_cust_id"`
	AcctID        string `json:"acct_id"`
	Balance       string `json:"balance"`
	AcctBalance   string `json:"acct_balance"`
	FreezeBalance string `json:"freeze_balance"`
	AcctStat      string `json:"acct_stat"`
}
