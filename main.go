package main

import (
	"context"
	"fmt"
	"test/book/dal"
	"test/book/dal/model"
	"test/book/dal/query"
)

// gen demo

// MySQLDSN MySQL data source name
const MySQLDSN = "root:root@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True"

func init() {
	dal.DB = dal.ConnectDB(MySQLDSN).Debug()
}

func main() {

	// 设置默认DB对象
	query.SetDefault(dal.DB)

	// 创建
	b1 := model.BooksDO{
		Title:  "《七米的Go语言之路》",
		Author: "七米",
		Price:  100,
	}
	err := query.BooksDO.WithContext(context.Background()).Create(&b1)
	if err != nil {
		fmt.Printf("create book fail, err:%v\n", err)
		return
	}

	// 更新
	ret, err := query.BooksDO.WithContext(context.Background()).
		Where(query.BooksDO.Author.Eq("七米")).
		Update(query.BooksDO.Price, 200)
	if err != nil {
		fmt.Printf("update book fail, err:%v\n", err)
		return
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)

	// 查询
	query1()
	filter()
	search()

	// 删除
	ret, err = query.BooksDO.WithContext(context.Background()).Where(query.BooksDO.ID.Eq(1)).Delete()
	if err != nil {
		fmt.Printf("delete book fail, err:%v\n", err)
		return
	}
	fmt.Println("RowsAffected:", ret.RowsAffected)
}
func query1() {
	fmt.Println("=====query1======")
	rets, err := query.BooksDO.WithContext(context.Background()).GetBooksByAuthor("七米")
	if err != nil {
		fmt.Printf("GetBooksByAuthor fail, err:%v\n", err)
		return
	}
	for i, b := range rets {
		fmt.Printf("%d:%v\n", i, b)
	}
	// 查询返回map
	bookMaps, err := query.BooksDO.WithContext(context.Background()).GetByIDReturnMap([]int{2, 3})
	if err != nil {
		fmt.Printf("GetByIDReturnMap fail, err:%v\n", err)
		return
	}
	for _, bookMap := range bookMaps {
		fmt.Println("-->")
		for k, v := range bookMap {
			fmt.Printf("%s:%v\n", k, v)
		}
		fmt.Println("<--")
	}
}
func filter() {
	fmt.Println("=====filter======")
	// 过滤
	rets, err := query.BooksDO.WithContext(context.Background()).FilterWithColumn("author", "wang")
	if err != nil {
		fmt.Printf("FilterWithColumn fail, err:%v\n", err)
		return
	}
	for i, b := range rets {
		fmt.Printf("%d:%+v\n", i, b)
	}
}
func search() {
	fmt.Println("=====search======")
	// 搜索
	b := &model.BooksDO{Author: "wang"}
	rets, err := query.BooksDO.WithContext(context.Background()).Search(b)
	if err != nil {
		fmt.Printf("Search fail, err:%v\n", err)
		return
	}
	for i, b := range rets {
		fmt.Printf("%d:%v\n", i, b)
	}
}
