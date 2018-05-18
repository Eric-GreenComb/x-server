package regexp

import (
	"regexp"
)

// IsGoodPwd 是否是合格的密码
func IsGoodPwd(pwd string) (bool, string) {
	if pwd == "" {
		return false, "密码不能为空"
	}
	if len(pwd) < 6 {
		return false, "密码至少6位"
	}
	return true, ""
}

// IsEmail 是否是email
func IsEmail(email string) bool {
	if email == "" {
		return false
	}
	ok, _ := regexp.MatchString(`^([a-zA-Z0-9]+[_|\_|\.|\-]?)*[_a-z\-A-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.|\-]?)*[a-zA-Z0-9\-]+\.[0-9a-zA-Z]{2,6}$`, email)
	return ok
}

// IsUsername 是否只包含数字, 字母 -, _
func IsUsername(username string) bool {
	if username == "" {
		return false
	}
	ok, _ := regexp.MatchString(`[^0-9a-zA-Z_\-]`, username)
	return !ok
}

const (
	regularMobile = "^1([38][0-9]|14[57]|5[^4])\\d{8}$"
)

// IsMobile check if mobile
func IsMobile(mobileNum string) bool {
	reg := regexp.MustCompile(regularMobile)
	return reg.MatchString(mobileNum)
}
