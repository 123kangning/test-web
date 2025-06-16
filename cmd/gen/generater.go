package main

// gorm gen configure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen/field"
	"test/book/dal/model"

	"gorm.io/gorm"

	"gorm.io/gen"
)

const MySQLDSN = "root:root@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True"

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db
}

func main() {
	// 指定生成代码的具体相对目录(相对当前文件)，默认为：./query
	// 默认生成需要使用WithContext之后才可以查询的代码，但可以通过设置gen.WithoutContext禁用该模式
	g := gen.NewGenerator(gen.Config{
		// 默认会在 OutPath 目录生成CRUD代码，并且同目录下生成 model 包
		// 所以OutPath最终package不能设置为model，在有数据库表同步的情况下会产生冲突
		// 若一定要使用可以通过ModelPkgPath单独指定model package的名称
		OutPath:           "../../dal/query/",
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,

		// gen.WithoutContext：禁用WithContext模式
		// gen.WithDefaultQuery：生成一个全局Query对象Q
		// gen.WithQueryInterface：生成Query接口
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// 通常复用项目中已有的SQL连接配置db(*gorm.DB)
	// 非必需，但如果需要复用连接时的gorm.Config或需要连接数据库同步表信息则必须设置
	gormDB := connectDB(MySQLDSN)
	//gormDB, _ := gorm.Open(rawsql.New(rawsql.Config{
	//	FilePath: []string{"../../init.sql"},
	//}))
	g.UseDB(gormDB)

	// 从连接的数据库为所有表生成Model结构体和CRUD代码
	// 也可以手动指定需要生成代码的数据表

	userBook := g.GenerateModelAs("user_books", "UserBooksDO")
	// 配置多对多关系
	user := g.GenerateModelAs("users", "UsersDO",
		gen.FieldRelate(field.Many2Many,
			"Books", g.GenerateModelAs("books", "BooksDO"),
			&field.RelateConfig{
				RelateSlicePointer: true,
				GORMTag: field.GormTag{
					"many2many":      []string{"user_books"}, // 中间表名
					"foreignKey":     []string{"ID"},         // 当前表在中间表的外键
					"joinForeignKey": []string{"UserID"},     // 中间表中指向当前表的外键
					"references":     []string{"ID"},         // 关联表的主键
					"joinReferences": []string{"BookID"},     // 中间表中指向关联表的外键
				},
			},
		))
	//book := g.GenerateModelAs("books", "BooksDO")
	book := g.GenerateModelAs("books", "BooksDO",
		gen.FieldRelate(field.Many2Many,
			"Users", user,
			&field.RelateConfig{
				RelateSlicePointer: true,
				GORMTag: field.GormTag{
					"many2many":      []string{"user_books"},
					"foreignKey":     []string{"ID"},
					"joinForeignKey": []string{"BookID"},
					"references":     []string{"ID"},
					"joinReferences": []string{"UserID"},
				},
			},
		),
	)

	alarm_rule_channels := g.GenerateModelAs("alarm_rule_channels", "AlarmRuleChannelsDO")
	alarm_rules := g.GenerateModelAs("alarm_rules", "AlarmRulesDO",
		gen.FieldRelate(field.Many2Many, "Channels",
			g.GenerateModelAs("alarm_channels", "AlarmChannelsDO"),
			&field.RelateConfig{
				RelateSlicePointer: true,
				GORMTag: field.GormTag{
					"many2many":      []string{"alarm_rule_channels"},
					"foreignKey":     []string{"ID"},
					"joinForeignKey": []string{"RuleID"},
					"references":     []string{"ID"},
					"joinReferences": []string{"ChannelID"},
				},
			}))
	alarm_channels := g.GenerateModelAs("alarm_channels", "AlarmChannelsDO",
		gen.FieldRelate(field.Many2Many, "Rules",
			alarm_rules,
			&field.RelateConfig{
				RelateSlicePointer: true,
				GORMTag: field.GormTag{
					"many2many":      []string{"alarm_rule_channels"},
					"foreignKey":     []string{"ID"},
					"joinForeignKey": []string{"ChannelID"},
					"references":     []string{"ID"},
					"joinReferences": []string{"RuleID"},
				},
			}),
	)

	// 应用所有模型
	g.ApplyBasic(
		book,
		user,
		userBook,
		alarm_rule_channels,
		alarm_rules,
		alarm_channels,
	)
	// 通过ApplyInterface添加为book表添加自定义方法
	g.ApplyInterface(func(model.Querier) {}, book)
	// 为`Book`添加 `Filter`接口
	g.ApplyInterface(func(model.Filter) {}, book)
	// 通过ApplyInterface添加为book表添加Searcher接口
	g.ApplyInterface(func(model.Searcher) {}, book)
	// 执行并生成代码
	g.Execute()
}
