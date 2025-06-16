package model

import "gorm.io/gen"

type Searcher interface {
	// Search 根据指定条件查询书籍
	//
	// SELECT * FROM book
	// WHERE publish_date is not null
	// {{if book != nil}}
	//   {{if book.ID > 0}}
	//     AND id = @book.ID
	//   {{else if book.Author != ""}}
	//     AND author=@book.Author
	//   {{end}}
	// {{end}}
	Search(book *gen.T) ([]*gen.T, error)
}
