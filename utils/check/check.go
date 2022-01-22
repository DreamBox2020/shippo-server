package check

import (
	"regexp"
	"strings"
)

func CheckPhone(s string) (matched bool) {

	list := [43]string{
		// 移动
		// ""134"", // 0～8
		"135",
		"136",
		"137",
		"138",
		"139",
		"147", // 数据卡
		// ""148"", // 数据卡
		"150",
		"151",
		"152",
		"157", // 固话卡
		"158",
		"159",
		// ""172"", // 数据卡
		"178",
		"182",
		"183",
		"184",
		"187",
		"188",
		"195",
		"197",
		"198",
		// 联通
		"130",
		"131",
		"132",
		// ""145"", // 数据卡
		// ""146"", // 数据卡
		"155",
		"156",
		"166",
		// ""171"", // 副号卡
		"175",
		"176",
		"185",
		"186",
		"196",
		//电信
		"133",
		// ""149"", // 数据卡
		"153",
		"173",
		"177",
		"180",
		"181",
		"189",
		"190",
		"191",
		"193",
		"199",
	}

	// 如果不是1～9开头，且不是11位数字
	matched, err := regexp.MatchString("^[1-9][0-9]{10}$", s)
	if err != nil {
		return false
	}
	if !matched {
		return false
	}

	// 如果是134开头，且第四位在0～8之间
	matched, err = regexp.MatchString("^134[0-8]", s)
	if err != nil {
		return false
	}
	if matched {
		return true
	}

	// 循环判断开头
	for i := 0; i < len(list); i++ {
		if strings.HasPrefix(s, list[i]) {
			return true
		}
	}

	return false
}

func CheckSmsCode(s string) (matched bool) {
	matched, err := regexp.MatchString("^[1-9][0-9]{5}$", s)
	if err != nil {
		return false
	}
	return
}

func CheckUUID(s string) (matched bool) {
	matched, err := regexp.MatchString("^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$", s)
	if err != nil {
		return false
	}
	return
}

func CheckPassport(s string) (matched bool) {
	matched, err := regexp.MatchString("^[0-9a-z]{32}$", s)
	if err != nil {
		return false
	}
	return
}

func CheckQQ(s string) (matched bool) {
	matched, err := regexp.MatchString("^[1-9][0-9]{4,9}$", s)
	if err != nil {
		return false
	}
	return
}

func CheckQQEmail(s string) (matched bool) {
	matched, err := regexp.MatchString("^[1-9][0-9]{4,9}@qq\\.com$", s)
	if err != nil {
		return false
	}
	return
}
