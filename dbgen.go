package main

import (
	mysqld "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func asd() {
	config := mysqld.Config{
		User:   "dev",
		Passwd: "dev",
		Net:    "tcp",
		Addr:   "127.0.0.1:33060",
		DBName: "dev",
	}
	dsn := config.FormatDSN()
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/domain/models/gen",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	gormconn, _ := gorm.Open(mysql.Open(dsn))
	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormconn) // reuse your gorm db
	//all := g.GenerateAllTable()
	users := g.GenerateModel("users")
	pusers := g.GenerateModel("projects", gen.FieldRelate(field.Many2Many, "Users", users,
		&field.RelateConfig{
			//RelateSlice: true,
			GORMTag: field.GormTag{"many2many": []string{"project_user"}},
		}),
	)
	g.ApplyBasic(
		users, pusers,
	)
	//g.GenerateModel("projects", gen.FieldRelate("Users","Users", "users")),
	g.Execute()
}
