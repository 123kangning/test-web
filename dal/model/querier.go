// dal/model/querier.go

package model

import "gorm.io/gen"

// 通过添加注释生成自定义方法

type Querier interface {
	// GetByID
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error) // 返回结构体和error

	// GetByIDReturnMap
	// SELECT * FROM @@table WHERE id in @ids
	GetByIDReturnMap(ids ...int) ([]gen.M, error) // 返回 map 和 error

	// GetBooksByAuthor
	// SELECT * FROM @@table WHERE author=@author
	GetBooksByAuthor(author string) ([]*gen.T, error) // 返回数据切片和 error
}
