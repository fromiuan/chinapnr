// 4.2.2.	企业用户开户接口
package user

import (
	"reflect"

	"github.com/fromiuan/chinapnr/encoding"
)

type CorpReq struct {
	AppId               string             `json:"apply_id"`              // 必须 企业开户申请号 由商户提供,申请企业开户的唯一标识
	OrderId             string             `json:"order_id"`              // 必须 订单号 由商户生成，必须保证唯一，20位内的字母或数字组合
	OrderDate           string             `json:"order_date"`            // 必须 订单日期 格式为YYYYMMDD，例如：20160307
	OrderType           string             `json:"operate_type"`          // 必须 操作类型 		00090000 新增 00090000 新增
	CorpLicenseType     string             `json:"corp_license_type"`     // 必须 企业证照类型 	01030100.普通营业执照企业 01030101.三证合一企业
	CorpType            string             `json:"corp_type"`             // 必须 企业类型 		01030000.普通企业 01030000.普通企业
	CorpName            string             `json:"corp_name"`             // 必须 企业名称		企业的公司全称
	LicenseStartDate    string             `json:"license_start_date"`    // 必须 证照起始日期 企业的营业证照起始日期，精确到年月日
	LicenseEndDate      string             `json:"license_end_date"`      // 必须 证照结束日期 企业的营业证照结束日期，精确到年月日，支持“永久”
	CorpBusinessAddress string             `json:"corp_business_address"` // 必须 企业经营地址 企业的经营地址
	CorpRegAddress      string             `json:"corp_reg_address"`      // 必须 企业注册地址 企业注册地址
	CorpFiexdTelePhone  string             `json:"corp_fixed_telephone"`  // 必须 企业固定电话 企业固定电话
	BusinessScope       string             `json:"business_scope"`        // 必须 企业的经营范围 企业的经营范围
	ControllingShare    []ControllingShare `json:"-"`                     // 必须 控股股东
	LegalName           string             `json:"legal_name"`            // 必须 法人姓名
	LegalCertType       string             `json:"legal_cert_type"`       // 必须 法人证件类型 法人证件类型,01020101:护照,01020102:军官证,01020103士兵证,01020104:回乡证,01020105警官证,01020106:台胞证,01020107:其他
	LegalCertId         string             `json:"legal_cert_id"`         // 必须 法人证件号码 法人的证件号码，与证件类型对应
	LegalCertStartDate  string             `json:"legal_cert_start_date"` // 必须 法人证件起始日期 法人的证件起始日期，精确到年月日
	LegalCertEndDate    string             `json:"legal_cert_end_date"`   // 必须 法人证件起始日期 法人的证件结束日期，精确到年月日，支持“永久”
	LegalMobile         string             `json:"legal_mobile"`          // 必须 法人手机号码 	法人的手机号码
	ContactName         string             `json:"contact_name"`          // 必须 企业联系人姓名 	企业联系人的姓名
	ContactMobile       string             `json:"contact_mobile"`        // 必须 联系人手机号 	联系人的手机号码
	ContactEmail        string             `json:"contact_email"`         // 必须 联系人邮箱 	联系人的邮箱地址
	BankAcctName        string             `json:"bank_acct_name"`        // 必须 开户银行账户名 	企业开户银行账户名，需与企业名称保持一致
	BankId              string             `json:"bank_id"`               // 必须 开户银行 	具体见附件：开户银行代号
	BankAcctNo          string             `json:"bank_acct_no"`          // 必须 开户银行账号 	企业开户银行账号
	BankBranch          string             `json:"bank_branch"`           // 必须 开户银行支行名称 	企业开户银行的支行名称
	BankProv            string             `json:"bank_prov"`             // 必须 开户银行省份 	银行卡开户省份
	BankArea            string             `json:"bank_area"`             // 必须 开户银行地区 	银行卡开户地区
	BgRetUrl            string             `json:"bg_ret_url"`            // 必须 商户后台应答地址 	通过后台异步通知商户开户结果 注意： 1) 使用时不要包含中文 2) 必须是外网地址
	BusinessCode        string             `json:"business_code"`         // 可选 营业执照注册号 企业的营业执照注册号
	InstitutionCode     string             `json:"institution_code"`      // 可选 组织机构代码 企业的组织机构代码证
	TaxCode             string             `json:"tax_code"`              // 可选 企业的税务登记号  税务登记证号
	SocialCreditCode    string             `json:"social_credit_code"`    // 可选 统一社会信用代码  企业的统一社会信用代码
	Industry            string             `json:"industry"`              // 可选 行业 	企业行业类别，具体见附件：行业
	RetUrl              string             `json:"ret_url"`               // 可选 页面返回URL 	交易完成后，本平台系统把交易结果通过页面方式，发送到该地址上
	MerPriv             string             `json:"mer_priv"`              // 可选 商户私有域 	为商户的自定义字段，该字段在交易完成后由本平台原样返回
	Extension           string             `json:"extension"`             // 可选 扩展域

}

type ControllingShare struct {
	Name     string `json:"name"`
	CertType string `json:"certType"`
	CertId   string `json:"certId"`
}

type CorpRsp struct {
	CmdId          string `json:"cmd_id"`
	RespCode       string `json:"resp_code"`
	RespDesc       string `json:"resp_desc"`
	ApplyId        string `json:"apply_id"`
	UserCustId     string `json:"user_cust_id"`
	OrderId        string `json:"order_id"`
	OrderDate      string `json:"order_date"`
	MerCustId      string `json:"mer_cust_id"`
	CorpName       string `json:"corp_name"`
	AcctID         string `json:"acct_id"`
	CashBindCardId string `json:"cash_bind_card_id"`
	AuditStatus    string `json:"audit_status"`
	AuditDesc      string `json:"audit_desc"`
	BgRetUrl       string `json:"bg_ret_url"`
	RetUlr         string `json:"ret_url"`
	MerPriv        string `json:"mer_priv"`
	Extension      string `json:"extension"`
}

func (this CorpReq) APICmId() string {
	return "102"
}

func (this CorpReq) ExtParams() map[string]string {
	data := make(map[string]string)
	controlling, _ := encoding.Marshal(this.ControllingShare)
	data["controlling_shareholder"] = string(controlling)
	return nil
}

func (this CorpReq) Params() map[string]string {
	elem := reflect.ValueOf(&this).Elem()
	type_ := elem.Type()
	map_ := make(map[string]string)
	for i := 0; i < type_.NumField(); i++ {
		map_[type_.Field(i).Tag.Get("json")] = elem.Field(i).String()
	}
	return map_
}
