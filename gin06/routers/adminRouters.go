package routers

import (
	"gin06/controllers/admin"
	"gin06/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.InitAdminAuthMiddleware) //在这里可以加载中间件
	{

		adminRouters.GET("/", admin.MainController{}.Index)
		adminRouters.GET("/welcome", admin.MainController{}.Welcome)

		adminRouters.GET("/login", admin.LoginController{}.Index)
		adminRouters.GET("/loginOut", admin.LoginController{}.LoginOut)
		adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
		adminRouters.POST("/doLogin", admin.LoginController{}.DoLogin)

		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.POST("/manager/doAdd", admin.ManagerController{}.DoAdd)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.POST("/manager/doEdit", admin.ManagerController{}.DoEdit)
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)

		adminRouters.GET("/focus", admin.FocusController{}.Index)
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)

		adminRouters.GET("/role", admin.RoleController{}.Index)
		adminRouters.GET("/role/add", admin.RoleController{}.Add)
		adminRouters.POST("/role/doAdd", admin.RoleController{}.DoAdd)
		adminRouters.POST("/role/doEdit", admin.RoleController{}.DoEdit)
		adminRouters.GET("/role/edit", admin.RoleController{}.Edit)
		adminRouters.GET("/role/delete", admin.RoleController{}.Delete)
		adminRouters.GET("/role/auth", admin.RoleController{}.Auth)
		adminRouters.POST("/role/doAuth", admin.RoleController{}.DoAuth)

		adminRouters.GET("/access", admin.AccessController{}.Index)
		adminRouters.GET("/access/add", admin.AccessController{}.Add)
		adminRouters.POST("/access/doAdd", admin.AccessController{}.DoAdd)

		adminRouters.GET("/access/edit", admin.AccessController{}.Edit)
		adminRouters.POST("/access/doEdit", admin.AccessController{}.DoEdit)
		adminRouters.GET("/access/delete", admin.AccessController{}.Delete)

	}
}
