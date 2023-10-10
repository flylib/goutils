package gormgen

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateTableStruct(db *gorm.DB, options ...Option) {
	opt := option{}
	opt.OutPath = "./model"
	for _, f := range options {
		f(&opt)
	}
	//根据配置实例化 gen
	g := gen.NewGenerator(opt.Config)
	//使用数据库的实例
	g.UseDB(db)
	//模型结构体的命名规则
	g.WithModelNameStrategy(opt.ModelNameHandler)
	//模型文件的命名规则
	g.WithFileNameStrategy(opt.FileNameHandler)

	//数据类型映射
	// 自定义字段的数据类型
	dataMap := map[string]func(detailType string) (dataType string){
		"bool":      func(detailType string) (dataType string) { return "bool" },
		"tinyint":   func(detailType string) (dataType string) { return "int8" },
		"smallint":  func(detailType string) (dataType string) { return "int16" },
		"mediumint": func(detailType string) (dataType string) { return "int32" },
		"bigint":    func(detailType string) (dataType string) { return "int64" },
		"int":       func(detailType string) (dataType string) { return "int64" },
	}

	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)
	if len(opt.genTables) == 0 {
		//生成数据库内所有表的结构体
		g.GenerateAllTable()
	} else {
		for _, table := range opt.genTables {
			//生成某张表的结构体
			g.GenerateModel(table)
		}
	}
	//执行
	g.Execute()
}
