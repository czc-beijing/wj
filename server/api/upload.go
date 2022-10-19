package api

import (
	"github.com/gin-gonic/gin"
	"wj/global"
	"wj/models/app"
	"wj/response"
)

type WebFileUpload struct {
}

func GetWebFileUpload() *WebFileUpload {
	return &WebFileUpload{}
}

// 图片文件上传
func (f *WebFileUpload) FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Failed("图片上传出错", c)
	}
	image := global.Config.Upload
	err = c.SaveUploadedFile(file, image.SavePath+file.Filename)
	if err != nil {
		return
	}
	imageURL := image.AccessUrl + file.Filename
	response.Success("上传图片成功", imageURL, c)
}

// 图片文件上传
func (f *WebFileUpload) FileUploadApp(c *gin.Context) {
	file, err := c.FormFile("file1")
	if err != nil {
		response.Failed("图片上传出错", c)
	}
	image := global.Config.Upload
	err = c.SaveUploadedFile(file, image.SavePath+file.Filename)
	if err != nil {
		return
	}
	imageURL := image.AccessUrl + file.Filename
	var res []app.FileInfo
	res[0].Path = imageURL
	response.Success("上传图片成功", res, c)
}
