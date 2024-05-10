package file_upload

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"gin_blog/global"
	"gin_blog/models"
	"gin_blog/models/image_type"
	"gorm.io/gorm"
	"os"
	"path"
)

func FileUp(imageType int, filename string, fileDate []byte, notUse bool) (imageModel models.ImageModel, err error) {
	//根据传入类型放入不同目录
	//修改文件名称为上传时间，以防重名覆盖的情况
	imagePath := checkDirectory(imageType, filename)
	//计算文件的hash
	hash := Md5(fileDate)

	//查询文件是否已经存在
	err = global.DB.Where("image_key = ?", hash).Take(&imageModel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		//说明文件不为空且纯在
		global.LOG.Warnf("文件%s已存在", filename)
		return imageModel, errors.New("文件已存在")
	}
	//此时上传文件至指定的目录
	out, err := os.Create(imagePath)
	if err != nil {
		defer out.Close()
		out.Write(fileDate)
	}
	imageModel = models.ImageModel{
		Path:      "/" + imagePath,
		ImageType: image_type.ImageType(imageType),
		Key:       hash,
		Name:      filename,
		NotUse:    notUse,
	}
	return imageModel, nil
}

func checkDirectory(imageType int, imageName string) (imagePath string) {
	it := image_type.ImageType(int64(imageType))
	basePath := global.CONFIG.Local.Path
	switch it {
	case image_type.TypeAvatar:
		return path.Join(basePath, "avatar", imageName)
	case image_type.TypeCover:
		return path.Join(basePath, "cover", imagePath)
	default:
		return ""
	}
}

// Md5加密
func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
