package commont_ser

import (
	"fmt"
	"gin_blog/global"
	"gin_blog/models"
	"gorm.io/gorm"
)

type ListOption struct {
	Page         models.PageInfo // 页码等关键信息
	Like         []string        // 模糊查询条件
	Preload      []string        //	预加载字段
	Where        []string        // 精确查询条件
	Wheres       []*gorm.DB
	Debug        bool // 是否开启debug模式
	PreloadWhere any  // 预加载字段精确查询条件
}

type CommonResponse struct {
	List  any `json:"list"`
	Count int `json:"count"`
}

func CommonList[T any](model T, option ListOption) (list []T, count int, err error) {
	var DB *gorm.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	} else {
		DB = global.DB
	}

	offset := (option.Page.Page - 1) * option.Page.Limit

	//where匹配
	query := DB.Where(model)

	if len(option.Where) > 0 {
		for _, s := range option.Where {
			query = query.Where(s)
		}
	}

	if len(option.Wheres) > 0 {
		for _, s := range option.Where {
			query = query.Where(s)
		}
	}
	//模糊查询
	if option.Page.Key != "" {
		//遍历其中的like
		for index, column := range option.Like {
			if index == 0 {
				query = query.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Page.Key))
			} else {
				query = query.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Page.Key))
			}
		}
	}

	//查询总条数
	count = int(query.Find(&list).RowsAffected)

	query = query.Limit(option.Page.Limit).Offset(offset).Order(option.Page.Sort)

	//预加载
	if len(option.Preload) > 0 {
		for _, preload := range option.Preload {
			query = query.Preload(preload)
		}
	}

	//预加载精确查询
	preloadWhere, ok := option.PreloadWhere.([][]string)
	if ok {
		for _, db := range preloadWhere {
			l := db[1:]
			interfaceValuse := make([]interface{}, len(l))
			for i, v := range l {
				interfaceValuse[i] = v
			}
			query = query.Preload(db[0], interfaceValuse...)
		}
	}

	//执行查询
	err = query.Find(&list).Error

	return list, count, err
}
