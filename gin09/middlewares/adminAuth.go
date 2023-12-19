package middlewares

import (
	"encoding/json"
	"fmt"
	"gin09/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	fmt.Println(pathname)
	fmt.Println("pathname", pathname)
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
		} else {
			urlPath := strings.Replace(pathname, "/admin/", "", 1)
			if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {

				// 1、根据角色获取当前角色的权限列表,然后把权限id放在一个map类型的对象里面
				var roleAccess []models.RoleAccess
				models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
				roleAccessMap := make(map[int]int)
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = v.AccessId
				}
				// 2、获取当前访问的url对应的权限id 判断权限id是否在角色对应的权限
				access := models.Access{}
				models.DB.Where("url = ?", urlPath).Find(&access)
				fmt.Println(access)

				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(200, "没有权限")
					c.Abort()
				}

			}

		}
	} else {
		//用户未登录
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}

}

//排除权限判断的方法

func excludeAuthPath(urlPath string) bool {
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()

	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	// return true
	fmt.Println(excludeAuthPathSlice)
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
