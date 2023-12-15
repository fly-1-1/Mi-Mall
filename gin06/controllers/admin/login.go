package admin

import (
	"encoding/json"
	"fmt"
	"gin06/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {

	fmt.Println(models.Md5("12345"))

	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) DoLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//1 验证码是否正确
	verifyValue := c.PostForm("verifyValue")
	captchaId := c.PostForm("captchaId")
	flag := models.VerifyCaptcha(captchaId, verifyValue)
	//con.Success(c, "验证验证码成功", "/admin")
	if flag {
		//查询数据库判断用户密码是否存在
		var userinfoList []models.Manager
		password = models.Md5(password)
		models.DB.Where("username = ? and password = ?", username, password).Find(&userinfoList)

		if len(userinfoList) > 0 {
			//执行登录 保存用户信息,跳转
			session := sessions.Default(c)
			//把结构体转为json字符串
			userinfoSlice, _ := json.Marshal(userinfoList)
			session.Set("userinfo", string(userinfoSlice))
			session.Save()
			con.Success(c, "登陆成功", "/admin")

		} else {
			con.Error(c, "用户或密码错误", "/admin/login")
		}
	} else {
		con.Error(c, "验证验证码失败", "/admin/login")
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

func (con LoginController) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	con.Success(c, "退出登陆成功", "/admin")
}
