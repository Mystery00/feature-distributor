package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "core/db/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	gormdb, _ := gorm.Open(mysql.Open("root:12345678@(127.0.0.1:3306)/feature_distributor?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb)

	g.ApplyBasic(
		g.GenerateModel("project",
			gen.FieldType("id", "int64"),
		),
		g.GenerateModel("toggle",
			gen.FieldType("id", "int64"),
			gen.FieldType("project_id", "int64"),
			gen.FieldType("enable", "bool"),
			gen.FieldType("value_type", "int8"),
			gen.FieldType("default_value", "int64"),
			gen.FieldType("return_value_when_disable", "int64"),
		),
		g.GenerateModel("toggle_value",
			gen.FieldType("id", "int64"),
			gen.FieldType("toggle_id", "int64"),
		),
	)
	g.Execute()
}
