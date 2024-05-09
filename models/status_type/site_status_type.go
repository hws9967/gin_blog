package status_type

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type SiteStatusType int8

const (
	TypeAudit  SiteStatusType = 0 // 待审核
	TypePass   SiteStatusType = 1 // 已通过
	TypeReject SiteStatusType = 2 // 被驳回
)

func (t *SiteStatusType) Scan(value interface{}) error {
	v, _ := value.([]uint8)
	val, _ := strconv.Atoi(string(v))
	*t = SiteStatusType(val)
	return nil
}

func (t *SiteStatusType) Value() (driver.Value, error) {
	// 将数字转换为值
	return int64(*t), nil
}

func (t *SiteStatusType) MarshalJSON() ([]byte, error) {
	var tp string
	switch *t {
	case TypeAudit:
		tp = "待审核"
	case TypePass:
		tp = "已通过"
	case TypeReject:
		tp = "被驳回"
	}
	return json.Marshal(tp)
}
