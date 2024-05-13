package image_api

import (
	"fmt"
	"gin_blog/global"
	response "gin_blog/models/common"
	"gin_blog/plugins/qiniu"
	"github.com/gin-gonic/gin"
	"io"
	"path"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否成功
	Msg       string `json:"msg"`        // 消息
}

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

	var resList []FileUploadResponse
	for _, file := range files {
		filePath := path.Join("uploads", file.Filename)
		//判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size > 2 {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "文件大小超过2M",
			})
			continue
		}
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.LOG.Error(err)
			continue
		}
	}
	//lm.LoggerApp.Send("上传文件成功", c)
	response.OkWithMessage("文件上传成功", c)

}

// ImageUpload 图片上传到七牛云
func (i ImageApi) ImageUpload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		response.FailWithMessage("请上传图片", c)
		return
	}

	// 打开文件
	readFile, _ := file.Open()
	// 读取文件内容
	fileData, _ := io.ReadAll(readFile)
	filePath, err := qiniu.UploadImage(fileData, "gin_blog")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(filePath, "文件上传成功", c)

}
