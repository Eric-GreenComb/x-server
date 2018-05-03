package regexp

import (
	"testing"
)

func TestIsEmail(t *testing.T) {

	_is := IsEmail("ministor@126.com")

	if !_is {
		t.Errorf("Isn't Mail")
	}

	_is = IsEmail("ministor11@126.cn.com")

	if !_is {
		t.Errorf("Isn't Mail")
	}

	_is = IsEmail("minis@tor11@126.cn.com")

	if _is {
		t.Errorf("Isn't Mail")
	}
}

func TestIsMobile(t *testing.T) {

	_is := IsMobile("13810167616")

	if !_is {
		t.Errorf("Isn't Mobile")
	}

}
