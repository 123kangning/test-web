package main

import (
	"test/book/dal"
	"test/book/routes"

	"github.com/gin-gonic/gin"
)

const MySQLDSN = "root:root@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True"

func main() {
	// 初始化数据库
	dal.ConnectDB(MySQLDSN)

	// 创建Gin引擎
	r := gin.Default()

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务
	r.Run(":8080")
}
