package image_api

import (
	"fmt"
	"gin_blog/global"
	response "gin_blog/models/common"
	"github.com/gin-gonic/gin"
	"path"
)

// ImageUploadView 上传图片
func (i ImageApi) ImageUploadView(c *gin.Context) {

	forms, err := c.MultipartForm()
	if err != nil {
		global.LOG.Error(err)
		response.FailWithMessage("请设置请求头为multipart/form-data", c)
		return
	}
	files, ok := forms.File["images"]
	fmt.Println(forms, files, ok, c.GetHeader("Content-Type"))
	if !ok {
		response.FailWithMessage("不存在文件，请上传图片--", c)
		return
	}

	for _, file := range files {
		filePath := path.Join("uploads", file.Filename)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.LOG.Error(err)
			continue
		}
	}
	//lm.LoggerApp.Send("上传文件成功", c)
	response.OkWithMessage("文件上传成功", c)

}

func (i ImageApi) ImageUploadView1(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取文件名
	fmt.Println(fileHeader.Header)
	fmt.Println(fileHeader.Size)
	fmt.Println(fileHeader.Filename)
}
