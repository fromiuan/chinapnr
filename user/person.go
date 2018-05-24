// 4.2.1.	个人用户开户接口
package user

import (
	"reflect"
)

// 个人用户开户接口
type PersonReq struct {
	OrderId    string `json:"order_id"`    // 必须 订单号 由商户生成，必须保证唯一
	OrderDate  string `json:"order_date"`  // 必须 订单日期  格式为YYYYMMDD，例如：20160307
	UserName   string `json:"user_name"`   // 必须 用户姓名 用户的真实姓名
	CertId     string `json:"cert_id"`     // 必须 身份证号 用户的身份证号
	UserMobile string `json:"user_mobile"` // 必须 手机号 用户的手机号
	CustProv   string `json:"cust_prov"`   // 必须 用户省份 用户省份，如上海0031
	CustArea   string `json:"cust_area"`   // 必须 用户地区 用户地区，如上海3100
	BgRetUrl   string `json:"bg_ret_url"`  // 必须 商户后台应答地址 通过后台异步通知商户开户结果 1) 使用时不要包含中文	2) 必须是外网地址

	SoloFlg     string `json:"solo_flg"`     // 可选 个体户标识 不传默认个体户标志否：00000100个人开户；是：00000101个体户开户
	ValiData    string `json:"vali_date"`    // 可选 证件有效期 填写证件上有效期的截至日期，格式为：YYYYMMDD，例如：20290420
	CustAddress string `json:"cust_address"` // 可选 住址 填写证件上的住址
	OccUpation  string `json:"occupation"`   // 可选 职业 见个人职业分类表
	UserEmail   string `json:"user_email"`   // 可选 邮箱 用户的电子邮箱
	PayPassword string `json:"pay_password"` // 可选 支付密码 个体户标识为是时，必须输入
	MerProv     string `json:"mer_priv"`     // 可选 商户私有域 为商户的自定义字段，该字段在交易完成后由本平台原样返回
	Extension   string `json:"extension"`    // 可选 扩展域 用于扩展请求参数
}

// 个人用户开户接口返回参数
type PersonRsp struct {
	CmdId         string `json:"cmd_id"`
	RespCode      string `json:"resp_code"`
	RrspDesc      string `json:"resp_desc"`
	MerCustId     string `json:"mer_cust_id"`
	OrderId       string `json:"order_id"`
	OrderDate     string `json:"order_date"`
	UserCustId    string `json:"user_cust_id"`
	AcctId        string `json:"acct_id"`
	UserMobile    string `json:"user_mobile"`
	PlatFormSeqId string `json:"platform_seq_id"`
	FeeAmt        string `json:"fee_amt"`
	FeeAcctId     string `json:"fee_acct_id"`
	BgRetUrl      string `json:"bg_ret_url"`
	MerPriv       string `json:"mer_priv"`
	Extension     string `json:"extension"`
}

func (this PersonReq) APICmId() string {
	return "101"
}

func (this PersonReq) ExtParams() map[string]string {
	return nil
}

func (this PersonReq) Params() map[string]string {
	// j, _ := encoding.Marshal(persion)
	// m := make(map[string]string)
	// encoding.UnMarshal(j, &m)
	// return m
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
