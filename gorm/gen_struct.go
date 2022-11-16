package gorm

import (
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

type Config struct {
	//文件生成路径
	ModelPkgPath string
	//类型映射 column type <=> go type
	MapTypes map[string]string
	//指定需要生成的表，空则生成所有
	GenTables []string
	// generate model global configuration
	FieldNullable     bool // generate pointer when field is nullable
	FieldCoverable    bool // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values
	FieldSignable     bool // detect integer field's unsigned type, adjust generated data type
	FieldWithIndexTag bool // generate with gorm index tag
	FieldWithTypeTag  bool // generate with gorm column type tag
}

func GenerateTableStruct(db *gorm.DB, config Config) {
	//根据配置实例化 gen
	g := gen.NewGenerator(gen.Config{
		ModelPkgPath:      config.ModelPkgPath,
		FieldNullable:     config.FieldNullable,
		FieldCoverable:    config.FieldCoverable,
		FieldSignable:     config.FieldSignable,
		FieldWithIndexTag: config.FieldWithIndexTag,
		FieldWithTypeTag:  config.FieldWithTypeTag,
	})
	//使用数据库的实例
	g.UseDB(db)

	//模型结构体的命名规则
	g.WithModelNameStrategy(func(tableName string) (modelName string) {
		tableName = strings.TrimLeft(tableName, "_")
		if strings.HasPrefix(tableName, "tbl") {
			return firstUpper(tableName[3:])
		}
		if strings.HasPrefix(tableName, "table") {
			return firstUpper(tableName[5:])
		}

		return firstUpper(tableName)
	})
	//模型文件的命名规则
	g.WithFileNameStrategy(func(tableName string) (fileName string) {
		if strings.HasPrefix(tableName, "tbl") {
			return firstLower(tableName[3:])
		}
		if strings.HasPrefix(tableName, "table") {
			return firstLower(tableName[5:])
		}
		return tableName
	})
	//数据类型映射
	var dataMap map[string]func(string) (dataType string)
	for sqlType, goType := range config.MapTypes {
		dataMap[sqlType] = func(string) (dataType string) {
			return goType
		}
	}
	//使用上面的类型映射
	g.WithDataTypeMap(dataMap)
	if len(config.GenTables) == 0 {
		//生成数据库内所有表的结构体
		g.GenerateAllTable()
	} else {
		for _, table := range config.GenTables {
			//生成某张表的结构体
			g.GenerateModel(table)
		}
	}

	//执行
	g.Execute()
}

//字符串第一位改成小写
func firstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

//字符串第一位改成大写
func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
