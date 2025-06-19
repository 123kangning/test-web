package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"test/book/dal"
	"test/book/dal/model"
	"test/book/dal/query"
	"testing"
)

func TestMain(m *testing.M) {
	// 初始化测试数据库
	dal.ConnectDB("root:root@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True")

	// 运行测试
	m.Run()
}

func TestBookCRUD(t *testing.T) {
	// 初始化测试数据
	book := model.BooksDO{
		Title:  "测试书籍",
		Author: "测试作者",
		Price:  100,
	}

	t.Run("CreateBook", func(t *testing.T) {
		err := query.BooksDO.WithContext(context.Background()).Create(&book)
		assert.NoError(t, err)
	})

	t.Run("UpdatePrice", func(t *testing.T) {
		_, err := query.BooksDO.WithContext(context.Background()).
			Where(query.BooksDO.Author.Eq("测试作者")).
			Update(query.BooksDO.Price, 200)
		assert.NoError(t, err)
	})

	t.Run("QueryBooks", func(t *testing.T) {
		books, err := query.BooksDO.WithContext(context.Background()).
			Where(query.BooksDO.Author.Eq("测试作者")).
			Find()
		assert.NoError(t, err)
		assert.Greater(t, len(books), 0)
	})

	t.Run("DeleteBook", func(t *testing.T) {
		_, err := query.BooksDO.WithContext(context.Background()).
			Where(query.BooksDO.ID.Eq(book.ID)).
			Delete()
		assert.NoError(t, err)
	})
}
