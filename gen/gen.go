package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "core/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	gormdb, _ := gorm.Open(mysql.Open("root:12345678@(127.0.0.1:3306)/feature_distributor?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	g.Execute()
}
