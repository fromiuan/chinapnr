//公众号支付原生态js支付接口
package trade

type EmsCnplPay struct {
	AppID     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"Package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}
