package middlewares

import (
	"encoding/json"
	"fmt"
	"gin03/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	fmt.Println(pathname)
	//获取session保存信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言 是否为一个string
	userinfoStr, ok := userinfo.(string)

	if ok {
		//判断userinfo信息是否存在
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
			}
		}
	} else {
		//用户未登录
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}

}
