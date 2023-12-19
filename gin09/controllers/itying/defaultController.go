package itying

import (
	"github.com/gin-gonic/gin"
	. "github.com/hunterhug/go_image"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	c.String(200, "默认首页")
}
func (con DefaultController) Thumbnail1(c *gin.Context) {

	//按宽度进行比例缩放，输入输出都是文件
	//filename string, savepath string, width int
	filename := "static/upload/20231219/1702935534.png"
	savepath := "static/upload/20231219/1702935534--.png"
	err := ScaleF2F(filename, savepath, 600)
	if err != nil {
		c.String(200, "生成图片失败")
		return
	}
	c.String(200, "Thumbnail1 成功")
}

func (con DefaultController) Thumbnail2(c *gin.Context) {

	//按宽度进行比例缩放，输入输出都是文件
	//filename string, savepath string, width int
	filename := "static/upload/20231219/1702935534.png"
	savepath := "static/upload/20231219/1702935534--.png"
	err := ThumbnailF2F(filename, savepath, 400, 400)
	if err != nil {
		c.String(200, "生成图片失败")
		return
	}
	c.String(200, "Thumbnail2 成功")
}

func (con DefaultController) Qrcode1(c *gin.Context) {
	var png []byte
	png, err := qrcode.Encode("https://www.baidu.com", qrcode.Medium, 516)
	if err != nil {
		c.String(200, "生成二维码失败")
		return
	}
	c.String(200, string(png))
}

func (con DefaultController) Qrcode2(c *gin.Context) {
	savepath := "static/upload/qrcode.png"
	err := qrcode.WriteFile("", qrcode.Medium, 2048, savepath)
	if err != nil {
		c.String(200, "生成二维码失败")
		return
	}
	file, _ := ioutil.ReadFile(savepath)
	c.String(200, string(file))
}
