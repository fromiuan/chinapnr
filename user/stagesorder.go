//	消费分期订单信息录入接口
package user

type StagesOrderReq struct {
	Version          string `json:"version"`
	CmdID            string `json:"cmd_id"`
	MerCustID        string `json:"mer_cust_id"`
	OrderID          string `json:"order_id"`
	OrderDate        string `json:"order_date"`
	ProductID        string `json:"product_id"`
	ProductAmt       string `json:"product_amt"`
	ProductPeriods   string `json:"product_periods"`
	ProductStartTime string `json:"product_start_time"`
	ProductEndTime   string `json:"product_end_time"`
	ProductLocation  string `json:"product_location"`
	BgRetURL         string `json:"bg_ret_url"`
	MerPriv          string `json:"mer_priv"`
	Extension        string `json:"extension"`
}

type StagesOrderRsp struct {
	CmdID            string `json:"cmd_id"`
	RespCode         string `json:"resp_code"`
	RespDesc         string `json:"resp_desc"`
	MerCustID        string `json:"mer_cust_id"`
	OrderID          string `json:"order_id"`
	OrderDate        string `json:"order_date"`
	ProductID        string `json:"product_id"`
	ProductAmt       string `json:"product_amt"`
	ProductPeriods   string `json:"product_periods"`
	ProductStartTime string `json:"product_start_time"`
	ProductEndTime   string `json:"product_end_time"`
	ProductLocation  string `json:"product_location"`
	BgRetURL         string `json:"bg_ret_url"`
	MerPriv          string `json:"mer_priv"`
	Extension        string `json:"extension"`
}
