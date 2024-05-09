package status_type

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type ArticleStatusType int8

const (
	TypeAll  ArticleStatusType = 1 // 全部可见
	TypeMe   ArticleStatusType = 2 // 仅自己可见
	TypeFans ArticleStatusType = 3 // 仅粉丝可见
)

func (t *ArticleStatusType) Scan(value interface{}) error {
	v, _ := value.([]uint8)
	val, _ := strconv.Atoi(string(v))
	*t = ArticleStatusType(val)
	return nil
}

func (t *ArticleStatusType) Value() (driver.Value, error) {
	// 将数字转换为值
	return int64(*t), nil
}

func (t *ArticleStatusType) MarshalJSON() ([]byte, error) {
	var tp string
	switch *t {
	case TypeAll:
		tp = "全部可见"
	case TypeMe:
		tp = "仅自己可见"
	case TypeFans:
		tp = "仅粉丝可见"
	}
	return json.Marshal(tp)
}
