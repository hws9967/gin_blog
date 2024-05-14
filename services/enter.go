package services

import (
	"gin_blog/services/image_ser"
	"gin_blog/services/user_ser"
)

type ServiceGroup struct {
	ImageService image_ser.ImageService
	UserService  user_ser.UserService
}

var ServiceApp = new(ServiceGroup)
