package admin

import (
	"fmt"
	"gin02/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) DoLogin(c *gin.Context) {
	verifyValue := c.PostForm("verifyValue")
	captchaId := c.PostForm("captchaId")
	flag := models.VerifyCaptcha(captchaId, verifyValue)
	if flag {
		c.String(http.StatusOK, "验证码验证成功")
	} else {
		c.String(http.StatusOK, "验证码验证失败")
	}

}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}
