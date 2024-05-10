package api

import (
	"gin_blog/api/image_api"
	"gin_blog/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImageApi    image_api.ImageApi
}

var ApiGroupApp = new(ApiGroup)
