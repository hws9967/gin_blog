package api

import (
	"gin_blog/api/advert_api"
	"gin_blog/api/image_api"
	"gin_blog/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImageApi    image_api.ImageApi
	AdvertApi   advert_api.AdvertApi
}

var ApiGroupApp = new(ApiGroup)
