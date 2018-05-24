Chinapnr SDK for Golang


## 已实现接口

#### 用户管理接口
	// 个人用户开户
	// 企业用户开户
	//	快捷卡绑卡接口
	// 	绑定取现卡接口
	// 	银行卡解绑接口
	// 	独立验卡接口
	// 	发送短信
	//	 一阶段绑卡签约短信发送接口
	//	一阶段快捷支付短信发送接口
#### 交易类接口
	// 	卡快捷支付接口
	//	取现接口202
	//	转账接口203
	//	网银支付接口204
	//	退款接口205
	//	代扣支付接口207
	//	订单解冻接口212
	// 	APP支付（支付宝、微信、银联各种支付）218
	// 	快捷支付APP版接口
#### 查询类接口
	//	交易状态查询接口301
	// 	卡bin查询接口302


## 集成流程

从[汇付天下](http://www.chinapnr.com/)申请商户账号即可。

#### 沙箱环境

申请商户支付会从汇付得到测试商户，将测试商户切换至沙箱环境即可`dev=true`。

#### 应用信息配置

参考汇付给出的开发接口文档进行应用的配置。

#### 文档所需

汇付天下新支付+接口规范
返回错误说明
证书下载操作指引
支付+FAQ
java sadk

#### demo

``` Golang

package main

import (
	"github.com/fromiuan/chinapnr"
	"github.com/fromiuan/chinapnr/user"

	"fmt"
	"io/ioutil"
	"time"
)

var (
	version    = "10" // 版本号
	mercustid  = "66660000000xxxxx" // 商户号
	pfxPass    = "1111" // 证书密码
	pfxContent = []byte{} // 证书内容

	client *chinapnr.ChinapnrPay
)

func init() {
	content, err := ioutil.ReadFile("cert.pfx")
	if err != nil {
		panic("read file error")
	}
	pfxContent = content

	client = chinapnr.New(version, mercustid, pfxPass, pfxContent, true)
}

func main() {
	// 个人开户
	req := user.PersonReq{
		OrderId:    "4259587412696",
		OrderDate:  time.Now().Format("20060102"),
		UserName:   "张三丰",
		CertId:     "11022919950222593X",
		UserMobile: "130xxxxxxxx",
		CustProv:   "0011",
		CustArea:   "1100",
		BgRetUrl:   "http://www.chinapnr.com/",
	}
	rsp, err := client.User_PorsonRequest(req)
	if err != nil {
		fmt.Println("main error", err)
	}
}

```