package api

import (
	"gin_blog/api/advert_api"
	"gin_blog/api/image_api"
	"gin_blog/api/menu_api"
	"gin_blog/api/settings_api"
	"gin_blog/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImageApi    image_api.ImageApi
	AdvertApi   advert_api.AdvertApi
	MenuAPi     menu_api.MenuApi
	UserApi     user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
