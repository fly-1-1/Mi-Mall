package admin

import (
	"gin06/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

type MainController struct {
}

func (con MainController) Index(c *gin.Context) {
	//获取userinfo session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言 是否为一个string
	userinfoStr, ok := userinfo.(string)
	if ok {
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		//获取所有权限
		var accessList []models.Access
		models.DB.Where("module_id=?", 0).Preload("AccessList").Find(&accessList)

		var roleAccess []models.RoleAccess
		models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}
		//循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
		for i := 0; i < len(accessList); i++ {
			if _, ok := roleAccessMap[accessList[i].Id]; ok {
				accessList[i].Checked = true
			}
			for j := 0; j < len(accessList[i].AccessList); j++ {
				if _, ok := roleAccessMap[accessList[i].AccessList[j].Id]; ok {
					accessList[i].AccessList[j].Checked = true
				}
			}
		}

		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{

			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userinfoStruct[0].IsSuper,
		})

	} else {
		c.Redirect(302, "/admin/login")
	}

}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
