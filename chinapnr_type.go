package chinapnr

type ChinapnrParam interface {
	// 用于提供访问的 method
	APICmId() string

	// 返回参数列表
	Params() map[string]string

	// 返回扩展
	ExtParams() map[string]string
}
