package routers

import (
	"gin10/controllers/itying"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)
		defaultRouters.GET("/thumbnail1", itying.DefaultController{}.Thumbnail1)
		defaultRouters.GET("/thumbnail2", itying.DefaultController{}.Thumbnail2)
		defaultRouters.GET("/qrcode1", itying.DefaultController{}.Qrcode1)
		defaultRouters.GET("/qrcode2", itying.DefaultController{}.Qrcode2)

	}
}
