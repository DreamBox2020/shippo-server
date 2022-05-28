package ecode

var Messages = map[int]string{
	0:    "成功",
	-500: "服务器繁忙",
	-101: "未登录",
	-403: "访问权限不足",

	// temp
	501001: "定金订单不存在",
	501002: "补款订单不存在",
	501003: "定金订单已经被绑定",
	501004: "补款订单已经被绑定",
	501005: "定金金额不正确",
	501006: "补款金额不正确",

	// wx
	502001: "微信通行证不存在",
}
