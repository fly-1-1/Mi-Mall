package admin

import (
	"gin05/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	var managerList []models.Manager
	models.DB.Preload("Role").Find(&managerList)

	//fmt.Printf("%#v", managerList)

	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}

func (con ManagerController) Add(c *gin.Context) {
	//获取所有的角色
	var roleList []models.Role
	models.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {

	roleId, err1 := models.Int(c.PostForm("role_id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/manger/add")
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	email := strings.Trim(c.PostForm("email"), " ")

	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户名或密码长度不合法", "/admin/manager/add")
		return
	}
	//判断管理员是否存在
	var managerList []models.Manager
	models.DB.Where("username = ?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(c, "此管理员已经存在", "/admin/manager/add")
		return
	}
	//增加管理员
	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Email:    email,
		Mobile:   mobile,
		RoleId:   roleId,
		Status:   1,
		AddTime:  int(models.GetUnix()),
	}
	err2 := models.DB.Create(&manager).Error
	if err2 != nil {
		con.Error(c, "增加管理员失败", "/admin/manager/add")
	}
	con.Success(c, "增加管理员成功", "/admin/manager/add")
}

func (con ManagerController) Edit(c *gin.Context) {

	//获取管理员
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误 请重试", "/admin/manager")
		return
	}
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)

	var roleList []models.Role
	models.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {

	//获取
	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误 请重试", "/admin/manager")
		return
	}

	roleId, err2 := models.Int(c.PostForm("role_id"))
	if err2 != nil {
		con.Error(c, "传入数据错误 请重试", "/admin/manager")
		return
	}

	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	email := strings.Trim(c.PostForm("email"), " ")

	manage := models.Manager{Id: id}
	models.DB.Find(&manage)
	manage.Username = username
	manage.Email = email
	manage.Mobile = mobile
	manage.RoleId = roleId

	//判断密码是否为空 位空修改密码否则不修改
	if password != "" {
		//判断密码长度l
		if len(password) < 6 {
			con.Error(c, "密码长度不合法 不可小于6位", "/admin/manager/edit?id="+models.String(id))
			return
		}
		manage.Password = models.Md5(password)
	}

	if len(mobile) > 11 {
		con.Error(c, "手机号码长度不可超过11位", "/admin/manager/edit?id="+models.String(id))
		return
	}

	err3 := models.DB.Save(&manage).Error
	if err3 != nil {
		con.Error(c, "修改失败", "/admin/manager/edit?id="+models.String(id))
		return
	}
	con.Success(c, "修改成功", "/admin/manager")
}

func (con ManagerController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误 请重试", "/admin/manager")
	} else {
		manager := models.Manager{Id: id}
		models.DB.Delete(&manager)
		con.Success(c, "删除角色成功", "/admin/manager")
	}
}
