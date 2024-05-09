package status_type

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type AccountStatusType int64

const (
	TypeTel   AccountStatusType = 1 // 手机号注册
	TypeQQ    AccountStatusType = 2 // QQ注册
	TypeGitee AccountStatusType = 3 // gitee注册
)

func (t *AccountStatusType) Scan(value interface{}) error {
	var l int
	switch val := value.(type) {
	case []uint8:
		l, _ = strconv.Atoi(string(val))
	case int64:
		l = int(val)
	}
	*t = AccountStatusType(l)
	return nil
}

func (t *AccountStatusType) Value() (driver.Value, error) {
	// 将数字转换为值
	return int64(*t), nil
}

func (t *AccountStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *AccountStatusType) String() string {
	var tp string
	switch *t {
	case TypeTel:
		tp = "手机号注册"
	case TypeQQ:
		tp = "QQ注册"
	case TypeGitee:
		tp = "gitee注册"
	}
	return tp
}

func Parse(statusType string) AccountStatusType {
	switch statusType {
	case "qq":
		return TypeQQ
	case "gitee":
		return TypeGitee
	}
	return TypeQQ
}
