package routers

import (
	"gin07/controllers/itying"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)
		defaultRouters.GET("/news", itying.DefaultController{}.News)

	}
}
