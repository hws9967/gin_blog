package image_api

import (
	"fmt"
	"gin_blog/global"
	response "gin_blog/models/common"
	"gin_blog/services/image_ser"
	"gin_blog/utils"
	"gin_blog/utils/jwts"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strings"
	"time"
)

func (ImageApi) ImageUploadDataView(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("参数校验失败", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	if claims.Role == 3 {
		response.FailWithMessage("游客用户不可上传图片", c)
		return
	}

	fileName := file.Filename
	basePath := global.Config.Upload.Path
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, image_ser.WhiteImageList) {
		response.FailWithMessage("非法文件", c)
		return
	}
	// 判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		msg := fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2fMB， 设定大小为：%dMB ", size, global.Config.Upload.Size)
		response.FailWithMessage(msg, c)
		return
	}

	// 创建这个文件夹 /uploads/file/nick_name
	dirList, err := os.ReadDir(basePath)
	if err != nil {
		response.FailWithMessage("文件目录不存在", c)
		return
	}
	if !isInDirEntry(dirList, claims.NickName) {
		// 创建这个目录
		err := os.Mkdir(path.Join(basePath, claims.NickName), 077)
		if err != nil {
			global.Log.Error(err)
		}
	}
	// 1.如果存在重名，就加随机字符串 时间戳
	// 2.直接+时间戳
	now := time.Now().Format("20060102150405")
	fileName = nameList[0] + "_" + now + "." + suffix
	filePath := path.Join(basePath, claims.NickName, fileName)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData("/"+filePath, c)
	return

}

func isInDirEntry(dirList []os.DirEntry, name string) bool {
	for _, entry := range dirList {
		if entry.Name() == name && entry.IsDir() {
			return true
		}
	}
	return false
}
