package chinapnr

import (
	"github.com/fromiuan/chinapnr/query"
	"github.com/fromiuan/chinapnr/trade"
	"github.com/fromiuan/chinapnr/user"
)

// 个人用户开户
func (this *ChinapnrPay) User_PorsonRequest(param user.PersonReq) (results *user.PersonRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// 企业用户开户
func (this *ChinapnrPay) User_CropRequest(param user.CorpReq) (results *user.CorpRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

//	快捷卡绑卡接口
func (this *ChinapnrPay) User_BindCard(param user.BindCardReq) (apiurl string, err error) {
	apiurl, err = this.getURL(param)
	return apiurl, err
}

// 	绑定取现卡接口
func (this *ChinapnrPay) User_BindGetCard(param user.BindGetCardReq) (results *user.BindCardRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// 	银行卡解绑接口
func (this *ChinapnrPay) User_UnBindCard(param user.UnBindCardReq) (results *user.BindCardRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// 	独立验卡接口
func (this *ChinapnrPay) User_VeriCard(param user.VeriCardReq) (results *user.BindCardRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// 	发送短信
func (this *ChinapnrPay) User_SendMsg(param user.SendMsgReq) (results *user.SendMsgRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

//	 一阶段绑卡签约短信发送接口
func (this *ChinapnrPay) User_SetpSendMsg(param user.SetpSendMsgReq) (results *user.SetpSendMsgRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

//	一阶段快捷支付短信发送接口
func (this *ChinapnrPay) User_SetpPay(param user.SetpPayReq) (results *user.SetpPayRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// ==========================================
// 	卡快捷支付接口
func (this *ChinapnrPay) Trade_FasePay(param trade.FasePayReq) (apiurl string, err error) {
	apiurl, err = this.getURL(param)
	return apiurl, err
}

//	取现接口202
func (this *ChinapnrPay) Trade_Enchashment(param trade.EnchashmentReq) (results *trade.EnchashmentRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

//	转账接口203
func (this *ChinapnrPay) Trade_Transfer(param trade.TransferReq) (results *trade.TransferRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

//	网银支付接口204
func (this *ChinapnrPay) Trade_Ebank(param trade.EbankReq) (results *trade.EbankRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

//	退款接口205
func (this *ChinapnrPay) Trade_Refund(param trade.RefundReq) (results *trade.RefundRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

//	代扣支付接口207
func (this *ChinapnrPay) Trade_WithholdPay(param trade.WithholdPayReq) (results *trade.WithholdPayRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

//	订单解冻接口212
func (this *ChinapnrPay) Trade_UnfreezeOrder(param trade.UnfreezeOrderReq) (results *trade.UnfreezeOrderRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// 	APP支付（支付宝、微信、银联各种支付）218
func (this *ChinapnrPay) Trade_SoftAppPay(param trade.SoftAppPayReq) (results *trade.SoftAppPayRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// 	快捷支付APP版接口
func (this *ChinapnrPay) Trade_FasePayApp(param trade.FasePayAppReq) (apiurl string, err error) {
	apiurl, err = this.getURL(param)
	return apiurl, err
}

// ==========================================
//	交易状态查询接口301
func (this *ChinapnrPay) Query_TradeStatus(param query.TradeStatusReq) (results *query.TradeStatusRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// 	卡bin查询接口302
func (this *ChinapnrPay) Query_CardBin(param query.CardBinReq) (results *query.TradeStatusRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}
