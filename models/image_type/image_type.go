package image_type

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type ImageType int64

const (
	TypeCover  ImageType = 1 // 封面或者背景图
	TypeAvatar ImageType = 2 // 头像
)

func (t *ImageType) Scan(value interface{}) error {
	var l int
	switch val := value.(type) {
	case []uint8:
		l, _ = strconv.Atoi(string(val))
	case int64:
		l = int(val)
	}
	*t = ImageType(l)
	return nil
}

func (t ImageType) Value() (driver.Value, error) {
	// 将数字转换为值
	return int64(t), nil
}

func (t ImageType) MarshalJSON() ([]byte, error) {
	var tp string
	switch t {
	case TypeAvatar:
		tp = "头像"
	case TypeCover:
		tp = "封面"
	}
	return json.Marshal(tp)
}
