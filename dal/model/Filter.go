package model

import "gorm.io/gen"

// Filter 自定义Filter接口
type Filter interface {
	// FilterWithColumn
	// SELECT * FROM @@table WHERE @@column=@value
	FilterWithColumn(column string, value string) ([]gen.T, error)
}
