package models

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 创建store
// var store = base64Captcha.DefaultMemStore
// RedisStore 实现接口
var store base64Captcha.Store = RedisStore{}

// MakeCaptcha 获取验证码
func MakeCaptcha() (id, b64s string, err error) {
	var driver base64Captcha.Driver

	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driver = driverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err = c.Generate()

	return id, b64s, err
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
