package ecode

var (
	// 特殊错误
	OK                = New(0, "成功")
	NoLogin           = New(-101, "未登录")
	ServerErr         = New(-500, "服务器繁忙")
	AccessDenied      = New(-403, "访问权限不足")
	ErrRecordNotFound = New(-601, "查询结果为空")

	// 临时订单
	Temp_trade_20220108_Trade1NotFind   = New(501001, "定金订单不存在")
	Temp_trade_20220108_Trade2NotFind   = New(501002, "补款订单不存在")
	Temp_trade_20220108_Trade1Repeat    = New(501003, "定金订单已经被绑定")
	Temp_trade_20220108_Trade2Repeat    = New(501004, "补款订单已经被绑定")
	Temp_trade_20220108_Trade1AmountErr = New(501005, "定金金额不正确")
	Temp_trade_20220108_Trade2AmountErr = New(501006, "补款金额不正确")

	// 微信
	WxPassportIsNull         = New(502001, "微信通行证不存在")
	WxArticleUpdateProhibit  = New(502002, "文章已经发布，不允许修改。")
	WxOffiaccountIsNotLinked = New(502003, "该公众号未关联小程序")
	WxArticleURLError        = New(502004, "文章链接错误")
	WxArticleNotTempURLError = New(502005, "创建文章请使用临时链接")

	// 验证码错误
	CaptchaError     = New(503001, "验证码错误")
	CaptchaSendError = New(503002, "验证码发送失败")

	// 文件相关错误
	FileTypeUnknown = New(504001, "未知文件类型")
)
